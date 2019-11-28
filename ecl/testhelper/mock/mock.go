package mock

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"testing"

	"gopkg.in/yaml.v2"
)

const (
	MaxMockNum  = 100
	MaxPollNum  = 100
	FakeTokenID = "0123456789abcdef01234567890abcdef"
)

type MockController struct {
	Mocks    map[string][]Mock //key should be API path
	Trackers map[string]*Tracker
	Mux      *http.ServeMux
	Server   *httptest.Server
	AuthURL  string
}

type Mock struct {
	Request        RequestDetail    `yaml:"request"`
	Response       ResponseDetail   `yaml:"response"`
	ExpectedStatus []string         `yaml:"expectedStatus"`
	NewStatus      string           `yaml:"newStatus"`
	Counter        PollingCondition `yaml:"counter"`
	Tracker        *Tracker
}

type Tracker struct {
	Status      string
	PollCounter int
}

type RequestDetail struct {
	Method string     `yaml:"method"`
	Query  url.Values `yaml:"query"`
	Body   string     `yaml:"body"`
}

type ResponseDetail struct {
	Code int    `yaml:"code"`
	Body string `yaml:"body"`
}

// Polling Condition is used for status transition, e.g. PENDING_CREATE to ACTIVE
type PollingCondition struct {
	MinSize int `yaml:"min"`
	MaxSize int `yaml:"max"`
}

func NewMockController() *MockController {
	mc := &MockController{}
	mc.Mocks = make(map[string][]Mock, MaxMockNum)
	mc.Trackers = make(map[string]*Tracker, MaxMockNum)
	mc.Mux = http.NewServeMux()
	mc.Server = httptest.NewServer(mc.Mux)

	// ECL(Enterprise Cloud) provider specific setting.
	mc.AuthURL = os.Getenv("OS_AUTH_URL")
	os.Setenv("OS_AUTH_URL", mc.Endpoint()+"v3/")

	return mc
}

func (mc *MockController) TerminateMockControllerSafety() {
	mc.Server.Close()
	os.Setenv("OS_AUTH_URL", mc.AuthURL)
}

func (mc MockController) Endpoint() string {
	return mc.Server.URL + "/"
}

func (mc *MockController) Register(t *testing.T, trackKey string, path string, mockdata string) {
	m := Mock{}
	m.Counter.MinSize = -1
	m.Counter.MaxSize = MaxPollNum
	err := yaml.Unmarshal([]byte(mockdata), &m)
	if err != nil {
		t.Errorf("Failed to unmarshal mockdata: %s\n", err)
	}

	_, ok := mc.Mocks[path]
	if !ok {
		mc.Mocks[path] = []Mock{}
	}

	_, ok = mc.Trackers[trackKey]
	if !ok {
		mc.Trackers[trackKey] = &Tracker{}
	}

	m.Tracker = mc.Trackers[trackKey]
	mc.Mocks[path] = append(mc.Mocks[path], m)

}

func (mc MockController) StartServer(t *testing.T) {
	for k, v := range mc.Mocks {
		mc.setupHandler(t, k, v)
	}
}

func (mc MockController) setupHandler(t *testing.T, path string, mocks []Mock) {
	mc.Mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		var found bool

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Errorf("Failed to read request body %v\n", r.Body)
		}

		for _, v := range mocks {
			if v.Request.Method != r.Method {
				continue
			}

			if len(v.Request.Query) != 0 {
				if !reflect.DeepEqual(v.Request.Query, r.URL.Query()) {
					continue
				}
			}

			if v.Request.Body != "" && v.Request.Body != string(body) {
				continue
			}

			if v.Request.Method != r.Method {
				continue
			}

			if v.Request.Body != "" && v.Request.Body != string(body) {
				continue
			}

			if len(v.ExpectedStatus) != 0 {
				if !v.Tracker.matchExpectedStatus(v.ExpectedStatus) {
					continue
				}
			}

			if v.Request.Method == "GET" {
				if !v.Tracker.matchPollingCondition(v.Counter.MinSize, v.Counter.MaxSize) {
					continue
				}
			}

			found = true
			if v.NewStatus != "" {
				v.Tracker.Status = v.NewStatus
			}

			if v.Request.Method == "GET" {
				v.Tracker.PollCounter += 1
			} else {
				v.Tracker.PollCounter = 0
			}

			if v.Response.Body != "" {
				w.Header().Add("Content-Type", "application/json")
			}
			w.Header().Add("X-Subject-Token", FakeTokenID)
			w.WriteHeader(v.Response.Code)
			if v.Response.Body != "" {
				fmt.Fprintf(w, v.Response.Body)
			}
			break
		}

		if !found {
			t.Errorf("No suitable mock found for Request API %v %v\n", r.Method, r.URL)
			w.Header().Add("X-Subject-Token", FakeTokenID)
			w.WriteHeader(404)
			fmt.Fprintf(w, "")
		}

	})
}

func (t Tracker) matchExpectedStatus(s []string) bool {
	for _, v := range s {
		if t.Status == v {
			return true
		}
	}
	return false
}

func (t Tracker) matchPollingCondition(min int, max int) bool {
	if min != -1 || max != -1 {
		if t.PollCounter < min || t.PollCounter > max {
			return false
		}
	}
	return true
}

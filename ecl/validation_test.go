package ecl

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
)

type testCase struct {
	val         interface{}
	f           schema.SchemaValidateFunc
	expectedErr *regexp.Regexp
}

func runTestCases(t *testing.T, cases []testCase) {
	matchErr := func(errs []error, r *regexp.Regexp) bool {
		// err must match one provided
		for _, err := range errs {
			if r.MatchString(err.Error()) {
				return true
			}
		}

		return false
	}

	for i, tc := range cases {
		_, errs := tc.f(tc.val, "test_property")

		if len(errs) == 0 && tc.expectedErr == nil {
			continue
		}

		if len(errs) != 0 && tc.expectedErr == nil {
			t.Fatalf("expected test case %d to produce no errors, got %v", i, errs)
		}

		if !matchErr(errs, tc.expectedErr) {
			t.Fatalf("expected test case %d to produce error matching \"%s\", got %v", i, tc.expectedErr, errs)
		}
	}
}

func TestValidationIntInSlice(t *testing.T) {
	runTestCases(t, []testCase{
		{
			val: 42,
			f:   IntInSlice([]int{1, 42}),
		},
		{
			val:         42,
			f:           IntInSlice([]int{10, 20}),
			expectedErr: regexp.MustCompile("expected [\\w]+ to be one of \\[10 20\\], got 42"),
		},
		{
			val:         "InvalidValue",
			f:           IntInSlice([]int{10, 20}),
			expectedErr: regexp.MustCompile("expected type of [\\w]+ to be an integer"),
		},
	})
}

func TestValidateInitialConfig(t *testing.T) {
	runTestCases(t, []testCase{
		{
			val: map[string]interface{}{"format": "set", "data": "c2V0"},
			f:   ValidateInitialConfig(),
		},
		{
			val: map[string]interface{}{"format": "text", "data": "c2V0"},
			f:   ValidateInitialConfig(),
		},
		{
			val: map[string]interface{}{"format": "a10-json", "data": "c2V0"},
			f:   ValidateInitialConfig(),
		},
		{
			val:         map[string]interface{}{"format": "json", "data": "c2V0"},
			f:           ValidateInitialConfig(),
			expectedErr: regexp.MustCompile(`expected test_property.format to be one of`),
		},
		{
			val:         map[string]interface{}{"data": "c2V0"},
			f:           ValidateInitialConfig(),
			expectedErr: regexp.MustCompile(`test_property.format is required`),
		},
		{
			val:         map[string]interface{}{"format": "set"},
			f:           ValidateInitialConfig(),
			expectedErr: regexp.MustCompile(`test_property.data is required`),
		},
		{
			val:         map[string]interface{}{"format": "set", "data": "c2V0", "foo": "bar"},
			f:           ValidateInitialConfig(),
			expectedErr: regexp.MustCompile(`unsupported key "foo"`),
		},
	})
}

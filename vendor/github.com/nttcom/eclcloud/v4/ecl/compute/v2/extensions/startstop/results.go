package startstop

import "github.com/nttcom/eclcloud/v4"

// StartResult is the response from a Start operation. Call its ExtractErr
// method to determine if the request succeeded or failed.
type StartResult struct {
	eclcloud.ErrResult
}

// StopResult is the response from Stop operation. Call its ExtractErr
// method to determine if the request succeeded or failed.
type StopResult struct {
	eclcloud.ErrResult
}

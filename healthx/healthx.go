package healthx

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/w6d-io/x/cmdx"

	"github.com/w6d-io/x/httpx"

	"github.com/gorilla/mux"

	"github.com/w6d-io/x/errorx"
)

// Alive handles the liveness probe
func Alive(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// AddAliveHandler add the liveness probe handler function into the http server
func AddAliveHandler(r *mux.Router) {
	r.Methods("GET").Path("/health/alive").HandlerFunc(Alive)
}

// AddReadyHandler add the readiness probe handler function into the http server
func AddReadyHandler(r *mux.Router, c Checker) {
	r.Methods("GET").Path("/health/ready").Handler(c)
}

// Checkable should return nil when the thing they are checking is healthy, and an error otherwise.
type Checkable interface {
	Healthy() error
}

// Checker provides a way to make an endpoint which can be probed for system health.
type Checker struct {
	// Checks are the Checkable to be checked when probing.
	Checks []Checkable

	// Unhealthyhandler is called when one or more of the checks are unhealthy.
	// If not provided DefaultUnhealthyHandler is called.
	UnhealthyHandler UnhealthyHandler

	// HealthyHandler is called when all checks are healthy.
	// If not provided, DefaultHealthyHandler is called.
	HealthyHandler http.HandlerFunc
}

func (c Checker) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	unhealthyHandler := c.UnhealthyHandler
	if unhealthyHandler == nil {
		unhealthyHandler = DefaultUnhealthyHandler
	}

	successHandler := c.HealthyHandler
	if successHandler == nil {
		successHandler = DefaultHealthyHandler
	}

	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if err := Check(c.Checks); err != nil {
		unhealthyHandler(w, r, err)
		return
	}

	successHandler(w, r)
}

// UnhealthyHandler type for unhealthy
type UnhealthyHandler func(w http.ResponseWriter, r *http.Request, err error)

// StatusResponse payloads to respect for check error
type StatusResponse struct {
	Status  string                 `json:"status"`
	Details *StatusResponseDetails `json:"details,omitempty"`
}

// StatusResponseDetails details for payload error response
type StatusResponseDetails struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// Check runs all healthy functions from the list
func Check(checks []Checkable) (err error) {
	var errs []error
	for _, c := range checks {
		if e := c.Healthy(); e != nil {
			errs = append(errs, e)
		}
	}

	switch len(errs) {
	case 0:
		err = nil
	case 1:
		err = errs[0]
	default:
		err = fmt.Errorf("multiple health check failure: %v", errs)
	}

	return
}

// DefaultHealthyHandler is the handler used if no healthy handler present
func DefaultHealthyHandler(w http.ResponseWriter, r *http.Request) {
	err := httpx.EncodeHTTPResponse(r.Context(), w, StatusResponse{
		Status: "ok",
	})
	cmdx.ShouldWithCtx(r.Context(), "failed to write JSON response in DefaultHealthyHandler", err)
}

// DefaultUnhealthyHandler is the handler used if no unhealthy handler present
func DefaultUnhealthyHandler(w http.ResponseWriter, r *http.Request, err error) {
	e := errorx.GetError(err)
	if e.StatusCode == 0 {
		e.StatusCode = http.StatusInternalServerError
	}
	w.WriteHeader(e.StatusCode)
	writeErr := httpx.EncodeHTTPResponse(r.Context(), w, StatusResponse{
		Status: "error",
		Details: &StatusResponseDetails{
			Code:    e.StatusCode,
			Message: e.Error(),
		},
	})
	cmdx.ShouldWithCtx(r.Context(), "failed to write JSON response in DefaultUnhealthyHandler", writeErr)
}

type checkGoRoutine struct {
	threshold int
}

// Healthy from checkGoRoutine implements the checkable interface
func (c checkGoRoutine) Healthy() error {
	count := runtime.NumGoroutine()
	if count > c.threshold {
		return errorx.New(nil, fmt.Sprintf("too many goroutines (%d > %d)", count, c.threshold))
	}
	return nil
}

// NewGoRoutineChecker returns a new threshold go routine Checker instance
func NewGoRoutineChecker(threshold int) Checkable {
	return &checkGoRoutine{threshold: threshold}
}

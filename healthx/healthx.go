/*
Copyright 2020 WILDCARD SA.

Licensed under the WILDCARD SA License, Version 1.0 (the "License");
WILDCARD SA is register in french corporation.
You may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.w6d.io/licenses/LICENSE-1.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is prohibited.
Created on 14/10/2021
*/

package healthx

import (
	"fmt"
	"github.com/w6d-io/x/httpx"
	"net/http"
	"runtime"

	"github.com/w6d-io/x/errorx"
	"github.com/w6d-io/x/logx"

	"github.com/gorilla/mux"
)

func Alive(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func AddAliveHandler(r *mux.Router) {
	r.Methods("GET").Path("/health/ready").HandlerFunc(Alive)
}

// Checkable should return nil when the thing they are checking is healthy, and an error otherwise.
type Checkable interface {
	Healthy() error
}

// Checker provides a way to make an endpoint which can be probed for system health.
type Checker struct {
	// Checks are the Checkables to be checked when probing.
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

type UnhealthyHandler func(w http.ResponseWriter, r *http.Request, err error)

type StatusResponse struct {
	Status  string                 `json:"status"`
	Details *StatusResponseDetails `json:"details,omitempty"`
}

type StatusResponseDetails struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

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

func DefaultHealthyHandler(w http.ResponseWriter, r *http.Request) {
	err := httpx.EncodeHTTPResponse(r.Context(), w, StatusResponse{
		Status: "ok",
	})
	if err != nil {
		logx.WithName(r.Context(), "DefaultHealthyHandler").Error(err, "Failed to write JSON response")
	}
}

func DefaultUnhealthyHandler(w http.ResponseWriter, r *http.Request, err error) {
	e := errorx.GetError(err)
	if e.Code == 0 {
		e.Code = http.StatusInternalServerError
	}
	w.WriteHeader(e.Code)
	writeErr := httpx.EncodeHTTPResponse(r.Context(), w, StatusResponse{
		Status: "error",
		Details: &StatusResponseDetails{
			Code:    e.Code,
			Message: e.Error(),
		},
	})
	if writeErr != nil {
		logx.WithName(r.Context(), "DefaultHealthyHandler").Error(err, "Failed to write JSON response")
	}
}

type checkGoRoutine struct {
	threshold int
}

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

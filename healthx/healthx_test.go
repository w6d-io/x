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

package healthx_test

import (
	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/w6d-io/x/healthx"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("To known the status of the application", func() {
	Context("In alive actions", func() {
		BeforeEach(func() {
		})
		AfterEach(func() {
		})
		It("returns a http.StatusOK", func() {
			w := httptest.NewRecorder()
			healthx.Alive(w, nil)
			Expect(w.Code).To(Equal(http.StatusOK))
		})
		It("succeed to add the handler", func() {
			healthx.AddAliveHandler(mux.NewRouter())
		})
	})
	Context("In ready check", func() {
		It("run successfully ServerHTTP without replacing default handler", func() {
			req := httptest.NewRequest("GET", "/health/ready", nil)
			w := httptest.NewRecorder()
			c := healthx.Checker{
				Checks: []healthx.Checkable{
					&nopOk{},
				},
			}
			c.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusOK))
		})
		It("raises an error due to bad method", func() {
			req := httptest.NewRequest("PUT", "/health/ready", nil)
			w := httptest.NewRecorder()
			c := healthx.Checker{}
			c.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusMethodNotAllowed))
		})
		It("raises an error due to checks failed", func() {
			req := httptest.NewRequest("GET", "/health/ready", nil)
			w := httptest.NewRecorder()
			c := healthx.Checker{
				Checks: []healthx.Checkable{
					&nopKo{},
				},
			}
			c.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusInternalServerError))

			By("add ko check")
			c.Checks = append(c.Checks, &nopKo{})
			c.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusInternalServerError))
		})
		It("run go routine check and failed", func() {
			req := httptest.NewRequest("GET", "/health/ready", nil)
			w := httptest.NewRecorder()
			c := healthx.Checker{
				Checks: []healthx.Checkable{
					healthx.NewGoRoutineChecker(0),
				},
			}
			c.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusInternalServerError))
		})
		It("run go routine check and succeeded", func() {
			req := httptest.NewRequest("GET", "/health/ready", nil)
			w := httptest.NewRecorder()
			c := healthx.Checker{
				Checks: []healthx.Checkable{
					healthx.NewGoRoutineChecker(1000),
				},
			}
			c.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusOK))
		})
	})
})

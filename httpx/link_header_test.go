/*
Copyright 2020 WILDCARD SA.

Licensed under the WILDCARD SA License, Version 1.0 (the "License");
WILDCARD SA is register in french corporation.
You may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.w6d.io/licenses/LICENSE-1.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is prohibited.
Created on 16/11/2021
*/

package httpx_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/w6d-io/x/httpx"
)

var _ = Describe("in api rest response with pages", func() {
	Context("goes through all behaviour", func() {
		BeforeEach(func() {
		})
		AfterEach(func() {
		})
		It("successfully gets all data from links", func() {
			linkRaw := `<https://api.github.com/user/repos?page=2&per_page=1>; rel="next"; one="1", <https://api.github.com/user/repos?page=87&per_page=1>; rel="last"`
			links := httpx.Parse(linkRaw)
			Expect(links[0].Rel).To(Equal("next"))
			Expect(links[0].URLRaw).To(Equal("https://api.github.com/user/repos?page=2&per_page=1"))
			Expect(links.String()).To(Equal(`<https://api.github.com/user/repos?page=2&per_page=1>; one="1"; rel="next", <https://api.github.com/user/repos?page=87&per_page=1>; rel="last"`))
			Expect(links[0].HasParam("none")).To(Equal(false))
			Expect(links[0].HasParam("one")).To(Equal(true))
			Expect(len(links.FilterByRel("last"))).To(Equal(1))
			Expect(links[0].Param("one")).To(Equal("1"))
			Expect(links[0].Param("two")).To(Equal(""))
			Expect(len(httpx.ParseMultiple([]string{
				linkRaw,
				linkRaw,
			}))).To(Equal(4))
		})
		It("stress the parser", func() {
			linkRaw := `<https://api.github.com/user/repos?page=2&per_page=1>; rel="next"; one="1", <https://api.github.com/user/repos?page=87&per_page=1>; rel="last"; zest1=; zest2=""="3"`
			_ = httpx.Parse(linkRaw)
		})
		It("stress the parser with no key", func() {
			linkRaw := `<https://api.github.com/user/repos?page=87&per_page=1>; rel="last"; "zest1"; =; zest2=""="3";`
			_ = httpx.Parse(linkRaw)
		})
		It("gets nil string", func() {
			var l httpx.Links = nil
			Expect(l.String()).To(Equal(`<nil>`))
		})
	})
})

/*
Copyright 2020 WILDCARD SA.

Licensed under the WILDCARD SA License, Version 1.0 (the "License");
WILDCARD SA is register in french corporation.
You may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.w6d.io/licenses/LICENSE-1.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is prohibited.
Created on 07/10/2021
*/

package cmdx_test

import (
	. "github.com/onsi/ginkgo"

	"github.com/w6d-io/x/cmdx"
)

var _ = Describe("helper functions testing", func() {
	Context("checking Must behaviour", func() {
		It("Must works without printing", func() {
			cmdx.Must(nil, "never write")
		})
		It("Must works without printing", func() {
			//cmdx.Must(errors.New("test exits"), "all is good")
		})
	})
})

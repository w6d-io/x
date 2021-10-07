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
	. "github.com/onsi/gomega"

	"github.com/w6d-io/x/cmdx"
)

var _ = Describe("Command functions testing", func() {
	Context("Variables is not empty", func() {
		It("Version responds correctly", func() {
			By("Set variables")
			version := "v0.1.0-test"
			commit := "2c246741bce7284a8e728d7ca492dc3a47cd3c85"
			buildTime := "Thu Oct 7 08:47:07 2021 +0200"

			By("instanciate Version")
			cmd := cmdx.Version(&version, &commit, &buildTime)
			By("execute version command")
			Expect(cmd.Execute()).To(Succeed())

		})
	})
	Context("Variables is empty", func() {
		It("Version responds correctly", func() {
			By("Set variables")
			version := ""
			commit := ""
			buildTime := ""

			By("instanciate Version")
			cmd := cmdx.Version(&version, &commit, &buildTime)
			By("execute version command")
			Expect(cmd.Execute()).To(Succeed())

		})
	})
})

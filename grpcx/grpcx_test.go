/*
Copyright 2020 WILDCARD SA.

Licensed under the WILDCARD SA License, Version 1.0 (the "License");
WILDCARD SA is register in french corporation.
You may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.w6d.io/licenses/LICENSE-1.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is prohibited.
Created on 30/10/2021
*/

package grpcx_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/w6d-io/x/grpcx"
	"google.golang.org/grpc/peer"
)

var _ = Describe("run all grpc package functions", func() {
	Context("in the main file", func() {
		BeforeEach(func() {
		})
		AfterEach(func() {
		})
		It("gets options", func() {
			o := grpcx.GrpcOptions()
			Expect(len(o)).To(Equal(2))
		})
		It("gets context field with all metadata", func() {
			var ctx context.Context
			By("Set peer into context", func() {
				addr := &mockAddr{Address: "127.0.0.1:8080"}
				ctx = context.Background()
				p := &peer.Peer{
					Addr: addr,
				}
				ctx = peer.NewContext(ctx, p)
			})
			nCtx := grpcx.BeforeGrpcFunc(ctx, nil)
			Expect(nCtx.Value("correlation_id")).ToNot(BeNil())
			Expect(nCtx.Value("kind")).ToNot(BeNil())
			Expect(nCtx.Value("kind").(string)).To(Equal("grpc"))
			Expect(nCtx.Value("ipaddress")).ToNot(BeNil())
			Expect(nCtx.Value("ipaddress").(string)).To(Equal("127.0.0.1"))
		})
		It("gets context field without ipaddress into metadata", func() {
			var ctx context.Context
			By("Set a bad host port into peer context", func() {
				addr := &mockAddr{Address: "127.0.0.1#8080"}
				ctx = context.Background()
				p := &peer.Peer{
					Addr: addr,
				}
				ctx = peer.NewContext(ctx, p)
			})
			nCtx := grpcx.BeforeGrpcFunc(ctx, nil)
			Expect(nCtx.Value("correlation_id")).ToNot(BeNil())
			Expect(nCtx.Value("kind")).ToNot(BeNil())
			Expect(nCtx.Value("kind").(string)).To(Equal("grpc"))
			Expect(nCtx.Value("ipaddress")).To(BeNil())

			By("set a bad ip address", func() {
				addr := &mockAddr{Address: "127.0.0..1:8080"}
				ctx = context.Background()
				p := &peer.Peer{
					Addr: addr,
				}
				ctx = peer.NewContext(ctx, p)

			})
			nCtx = grpcx.BeforeGrpcFunc(ctx, nil)
			Expect(nCtx.Value("correlation_id")).ToNot(BeNil())
			Expect(nCtx.Value("kind")).ToNot(BeNil())
			Expect(nCtx.Value("kind").(string)).To(Equal("grpc"))
			Expect(nCtx.Value("ipaddress")).ToNot(BeNil())
			Expect(nCtx.Value("ipaddress").(string)).To(Equal("-"))

			By("an empty context")
			ctx = context.Background()
			nCtx = grpcx.BeforeGrpcFunc(ctx, nil)
			Expect(nCtx.Value("correlation_id")).ToNot(BeNil())
			Expect(nCtx.Value("kind")).ToNot(BeNil())
			Expect(nCtx.Value("kind").(string)).To(Equal("grpc"))
			Expect(nCtx.Value("ipaddress")).To(BeNil())
		})
	})
})

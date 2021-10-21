/*
Copyright 2020 WILDCARD SA.

Licensed under the WILDCARD SA License, Version 1.0 (the "License");
WILDCARD SA is register in french corporation.
You may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.w6d.io/licenses/LICENSE-1.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is prohibited.
Created on 21/10/2021
*/

package grpcx

import (
	"context"
	"net"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/w6d-io/x/errorx"
)

// GrpcOptions is used in go-kit transport to handle error response and before http function
func GrpcOptions() []grpctransport.ServerOption {
	return []grpctransport.ServerOption{
		grpctransport.ServerErrorHandler(errorx.NewErrorHandler()),
		grpctransport.ServerBefore(BeforeGrpcFunc),
	}
}

// BeforeGrpcFunc adds metadata into context
func BeforeGrpcFunc(ctx context.Context, _ metadata.MD) context.Context {
	correlationID := uuid.New().String()
	ctx = context.WithValue(ctx, "correlation_id", correlationID)
	ctx = context.WithValue(ctx, "kind", "http")
	p, _ := peer.FromContext(ctx)
	ip := p.Addr.String()
	ip, _, err := net.SplitHostPort(ip)
	if err != nil {
		ctrl.Log.WithName("Transport.beforeGrpcFunc").WithValues("correlation_id", correlationID).Error(err, "get ipaddress failed")
	}
	if ip != "" {
		userIP := net.ParseIP(ip)
		if userIP == nil {
			ip = "-"
		}
	}
	ctx = context.WithValue(ctx, "ipaddress", ip)
	return ctx
}

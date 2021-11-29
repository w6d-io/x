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
	ctx = context.WithValue(ctx, "kind", "grpc")
	p, ok := peer.FromContext(ctx)
	if !ok {
		return ctx
	}
	ip := p.Addr.String()
	ip, _, err := net.SplitHostPort(ip)
	if err != nil {
		ctrl.Log.WithName("Transport.beforeGrpcFunc").WithValues("correlation_id", correlationID).Error(err, "get ipaddress failed")
		return ctx
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

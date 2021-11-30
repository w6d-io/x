package grpcx

import (
	"context"
	"github.com/w6d-io/x/logx"
	"net"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/google/uuid"
	"github.com/w6d-io/x/errorx"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
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
	ctx = context.WithValue(ctx, logx.CorrelationId, correlationID)
	ctx = context.WithValue(ctx, logx.Kind, "grpc")
	p, ok := peer.FromContext(ctx)
	if !ok {
		return ctx
	}
	ip := p.Addr.String()
	ip, _, err := net.SplitHostPort(ip)
	if err != nil {
		logx.WithName(ctx, "Transport.beforeHttpFunc").Error(err, "get ipaddress failed")
		return ctx
	}
	if ip != "" {
		userIP := net.ParseIP(ip)
		if userIP == nil {
			ip = "-"
		}
	}
	ctx = context.WithValue(ctx, logx.IpAddress, ip)
	return ctx
}

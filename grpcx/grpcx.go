package grpcx

import (
	"context"
	"net"
	"strings"

	"github.com/w6d-io/x/logx"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"

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
	ctx = context.WithValue(ctx, logx.CorrelationID, correlationID)
	ctx = context.WithValue(ctx, logx.Kind, "grpc")
	p, ok := peer.FromContext(ctx)
	if !ok {
		return ctx
	}
	ip := p.Addr.String()
	if strings.Contains(ip, ":") {
		var err error
		ip, _, err = net.SplitHostPort(ip)
		if err != nil {
			logx.WithName(ctx, "Transport.beforeHttpFunc").Error(err, "get ipaddress failed")
			ip = "-"
		}
	}
	if ip != "" {
		userIP := net.ParseIP(ip)
		if userIP == nil {
			ip = "-"
		}
	}
	ctx = context.WithValue(ctx, logx.IPAddress, ip)
	return ctx
}

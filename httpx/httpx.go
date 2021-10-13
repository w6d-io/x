package httpx

import (
	"context"
	"encoding/json"
	"net"
	"net/http"

	"github.com/w6d-io/x/errorx"

	"github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
	ctrl "sigs.k8s.io/controller-runtime"
)

// ReadRemoteIP tries to find the public address ip from the http header
func ReadRemoteIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

// EncodeHTTPResponse writes the error from response if the response is a type of endpoint.Failer
// or returns the json encoded error
func EncodeHTTPResponse(w http.ResponseWriter, response interface{}) error {

	if f, ok := response.(endpoint.Failer); ok && f.Failed() != nil {
		errorx.ErrorEncoder(f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

// BeforeHttpFunc adds metadata into context
func BeforeHttpFunc(ctx context.Context, req *http.Request) context.Context {
	correlationID := uuid.New().String()
	ctx = context.WithValue(ctx, "correlation_id", correlationID)
	ctx = context.WithValue(ctx, "kind", "http")
	ip := ReadRemoteIP(req)
	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		ctrl.Log.WithName("Transport.beforeHttpFunc").WithValues("correlation_id", correlationID).Error(err, "get ipaddress failed")
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

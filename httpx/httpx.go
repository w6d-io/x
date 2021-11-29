package httpx

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"strings"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"github.com/w6d-io/x/logx"

	"github.com/w6d-io/x/errorx"

	"github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
	ctrl "sigs.k8s.io/controller-runtime"
)

// ReadRemoteIP tries to find the public address ip from the http header
func ReadRemoteIP(r *http.Request) string {
	var ipAddress string
	if r.Header != nil {
		ipAddress = r.Header.Get("X-Real-Ip")
		if ipAddress == "" {
			ipAddress = r.Header.Get("X-Forwarded-For")
		}
	}
	if ipAddress == "" {
		ipAddress = r.RemoteAddr
	}
	return ipAddress
}

// EncodeHTTPResponse writes the error from response if the response is a type of endpoint.Failer
// or returns the json encoded error
func EncodeHTTPResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	log := logx.WithName(ctx, "EncodeHTTPResponse")

	if f, ok := response.(endpoint.Failer); ok && f.Failed() != nil {
		log.Error(f.Failed(), "")
		errorx.ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err, ok := response.(errorx.Error); ok {
		errorx.ErrorEncoder(ctx, &err, w)
		return nil
	}
	r, ok := response.(proto.Message)
	if ok {
		var b []byte
		var err error
		opt := &protojson.MarshalOptions{
			EmitUnpopulated: true,
		}
		b, _ = opt.Marshal(r)
		_, err = w.Write(b)
		return err
	}
	return json.NewEncoder(w).Encode(response)
}

// BeforeHttpFunc adds metadata into context
func BeforeHttpFunc(ctx context.Context, req *http.Request) context.Context {
	correlationID := uuid.New().String()
	ctx = context.WithValue(ctx, "correlation_id", correlationID)
	ctx = context.WithValue(ctx, "kind", "http")
	if req.URL != nil {
		ctx = context.WithValue(ctx, "uri", req.URL.RequestURI())
	}
	ctx = context.WithValue(ctx, "method", strings.ToUpper(req.Method))
	ip := ReadRemoteIP(req)
	ip, _, err := net.SplitHostPort(ip)
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

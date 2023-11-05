package httpx

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"strings"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"github.com/w6d-io/x/errorx"
	"github.com/w6d-io/x/logx"

	"github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
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
		log.Error(&err, "")
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

// BeforeHTTPFunc adds metadata into context
func BeforeHTTPFunc(ctx context.Context, req *http.Request) context.Context {
	correlationID := uuid.New().String()
	ctx = context.WithValue(ctx, logx.CorrelationID, correlationID)
	ctx = context.WithValue(ctx, logx.Kind, "http")
	if req.URL != nil {
		ctx = context.WithValue(ctx, logx.URI, req.URL.RequestURI())
	}
	ctx = context.WithValue(ctx, logx.Method, strings.ToUpper(req.Method))
	ip := ReadRemoteIP(req)
	if strings.Contains(ip, ":") {
		var err error
		ip, _, err = net.SplitHostPort(ip)
		if err != nil {
			logx.WithName(ctx, "Transport.beforeHttpFunc").Error(err, "get ipaddress failed")
			return ctx
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

/*
Copyright 2020 WILDCARD SA.

Licensed under the WILDCARD SA License, Version 1.0 (the "License");
WILDCARD SA is register in french corporation.
You may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.w6d.io/licenses/LICENSE-1.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is prohibited.
*/

// Package logx provides tools to handle log
//
// WithName provides a new logr.Logger:
//
// ctx := context.WithValue(c.Request.Context(), "correlation_id", correlationID)
// ctx = context.WithValue(ctx, "ipaddress", ip)
// ctx = context.WithValue(ctx, "kind", "http")
// log := logx.WithName(ctx, "Test")
//
// GetLogValues returns a keysAndValues model compatible with logrLogger.
//
// var log logr.Logger
// ctx := context.WithValue(c.Request.Context(), "correlation_id", correlationID)
// ctx = context.WithValue(ctx, "ipaddress", ip)
// ctx = context.WithValue(ctx, "kind", "http")
// kv := logx.GetLogValues(ctx)
// log = zapr.NewLogger(zapLog)
// log.Info("test", kv...)
//
package logx

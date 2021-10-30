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

package cmdx

import (
	"context"
	"fmt"
	"os"

	"github.com/w6d-io/x/logx"
)

var OsExit = os.Exit

// Must fatales with the optional message if err is not nil.
func Must(err error, message string, args ...interface{}) {
	if err == nil {
		return
	}
	_, _ = fmt.Fprintf(os.Stderr, message+"\n", args...)
	OsExit(1)
}

// Should checks the error and write the message in log in Error level
func Should(message string, err error) {
	ShouldWithCtx(nil, message, err)
}

// ShouldWithCtx checks the error and write the message in log in Error level
func ShouldWithCtx(ctx context.Context, message string, err error) {
	if err != nil {
		logx.WithName(ctx, "Should").Error(err, message)
	}
}

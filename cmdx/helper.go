package cmdx

import (
	"context"
	"fmt"
	"os"

	"github.com/w6d-io/x/logx"
)

// OsExit is for unit-test hack
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

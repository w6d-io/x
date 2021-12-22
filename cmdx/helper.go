package cmdx

import (
	"context"
	"fmt"
	"github.com/w6d-io/x/errorx"
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
	if e, ok := err.(*errorx.Error); ok {
		_, _ = fmt.Fprintf(os.Stderr, "%+v\n", e)
		e.ShowStack()
	} else {
		_, _ = fmt.Fprintf(os.Stderr, "%+v\n", err)
	}
	OsExit(1)
}

// Should checks the error and write the message in log in Error level
func Should(err error, message string) {
	ShouldWithCtx(context.TODO(), message, err)
}

// ShouldWithCtx checks the error and write the message in log in Error level
func ShouldWithCtx(ctx context.Context, message string, err error) {
	if err != nil {
		logx.WithName(ctx, "Should").Error(err, message)
	}
}

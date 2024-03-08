package errorx_test

import (
	"testing"

	"go.uber.org/zap/zapcore"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	zapraw "go.uber.org/zap"
	ctrl "sigs.k8s.io/controller-runtime"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestErrorX(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ErrorX Suite")
}

var _ = BeforeSuite(func() {
	encoder := zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	opts := zap.Options{
		Encoder:         zapcore.NewConsoleEncoder(encoder),
		Development:     true,
		StacktraceLevel: zapcore.PanicLevel,
	}
	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts), zap.RawZapOpts(zapraw.AddCaller())))
})

var _ = AfterSuite(func() {
})

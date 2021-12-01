package k8x_test

import (
	"context"
	"testing"

	"github.com/w6d-io/x/logx"

	"github.com/google/uuid"
	"go.uber.org/zap/zapcore"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	zapraw "go.uber.org/zap"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestK8X(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, " Suite")
}

var cfg *rest.Config
var k8sClient client.Client
var testEnv *envtest.Environment
var ctx context.Context
var scheme = runtime.NewScheme()

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
		Level:           zapcore.Level(int8(-2)),
	}
	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts), zap.RawZapOpts(zapraw.AddCaller(), zapraw.AddCallerSkip(-2))))
	By("bootstrapping test environment")
	testEnv = &envtest.Environment{
		ErrorIfCRDPathMissing: false,
	}

	var err error
	cfg, err = testEnv.Start()
	Expect(err).ToNot(HaveOccurred())
	Expect(cfg).ToNot(BeNil())
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	k8sClient, err = client.New(cfg, client.Options{Scheme: scheme})
	Expect(err).ToNot(HaveOccurred())
	Expect(k8sClient).ToNot(BeNil())

	correlationID := uuid.New().String()
	ctx = context.Background()
	ctx = context.WithValue(ctx, logx.CorrelationID, correlationID)

}, 60)

var _ = AfterSuite(func() {
	By("tearing down the test environment")
	err := testEnv.Stop()
	Expect(err).ToNot(HaveOccurred())
})

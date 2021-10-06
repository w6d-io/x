package flagx

import (
	"flag"
	"fmt"
	zapraw "go.uber.org/zap"
	"os"
	ctrl "sigs.k8s.io/controller-runtime"
	"strconv"
	"strings"
	"text/tabwriter"

	"go.uber.org/zap/zapcore"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

// JsonEncoderConfig returns an opinionated EncoderConfig
func JsonEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

// TextEncoderConfig returns an opinionated EncoderConfig
func TextEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
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
}

// OutputFormatFlag contains structure for managing zap encoding
type OutputFormatFlag struct {
	ZapOptions *zap.Options
	value      string
}

func (o *OutputFormatFlag) String() string {
	return o.value
}

func (o *OutputFormatFlag) Set(flagValue string) error {
	flagValue = LookupEnvOrString("LOG_FORMAT", flagValue)
	val := strings.ToLower(flagValue)
	switch val {
	case "json":
		o.ZapOptions.Encoder = zapcore.NewJSONEncoder(JsonEncoderConfig())
	case "text":
		o.ZapOptions.Encoder = zapcore.NewConsoleEncoder(TextEncoderConfig())
	default:
		return fmt.Errorf(`invalid "%s"`, flagValue)
	}
	o.value = flagValue
	return nil
}

// levelStrings contains level string supported
var levelStrings = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"error": zapcore.ErrorLevel,
}

// LevelFlag contains structure for managing zap level
type LevelFlag struct {
	ZapOptions *zap.Options
	value      string
}

func (l LevelFlag) String() string {
	return l.value
}

func (l LevelFlag) Set(flagValue string) error {
	flagValue = LookupEnvOrString("LOG_LEVEL", flagValue)
	level, validLevel := levelStrings[strings.ToLower(flagValue)]
	if !validLevel {
		logLevel, err := strconv.Atoi(flagValue)
		if err != nil {
			return fmt.Errorf("invalid log level \"%s\"", flagValue)
		}
		if logLevel > 0 {
			intLevel := -1 * logLevel
			l.ZapOptions.Level = zapcore.Level(int8(intLevel))
		} else {
			return fmt.Errorf("invalid log level \"%s\"", flagValue)
		}
	} else {
		l.ZapOptions.Level = level
	}
	l.value = flagValue
	return nil
}

// BindFlags custom flags
func BindFlags(o *zap.Options) {
	var outputFormat OutputFormatFlag
	outputFormat.ZapOptions = o
	flag.Var(&outputFormat, "log-format", "log encoding ( 'json' or 'text')")

	var level LevelFlag
	level.ZapOptions = o
	flag.Var(&level, "log-level", "log level verbosity. Can be 'debug', 'info', 'error', "+
		"or any integer value > 0 which corresponds to custom debug levels of increasing verbosity")
}

// LookupEnvOrString adds the capability to get env instead of param flag
func LookupEnvOrString(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}

// LookupEnvOrBool adds the capability to get env instead of param flag
func LookupEnvOrBool(key string, defaultVal bool) bool {
	if val, ok := os.LookupEnv(key); ok {
		b, _ := strconv.ParseBool(val)
		return b
	}
	return defaultVal
}

// UsageFor function for flag usage
func UsageFor(short string) func() {
	return func() {
		_, _ = fmt.Fprintf(os.Stderr, "USAGE\n")
		_, _ = fmt.Fprintf(os.Stderr, "  %s\n", short)
		_, _ = fmt.Fprintf(os.Stderr, "\n")
		_, _ = fmt.Fprintf(os.Stderr, "FLAGS\n")
		w := tabwriter.NewWriter(os.Stderr, 0, 2, 2, ' ', 0)
		flag.VisitAll(func(f *flag.Flag) {
			_, _ = fmt.Fprintf(w, "\t-%s %s\t%s\n", f.Name, f.DefValue, f.Usage)
		})
		_ = w.Flush()
		_, _ = fmt.Fprintf(os.Stderr, "\n")
	}
}

// Init the default flags
func Init() *string {

	var (
		configPath = flag.String("config", LookupEnvOrString("CONFIG", "/data/etc/config.yaml"), "The path for the config file")
	)

	opts := zap.Options{
		Development:     os.Getenv("RELEASE") != "prod",
		StacktraceLevel: zapcore.PanicLevel,
		Encoder:         zapcore.NewConsoleEncoder(TextEncoderConfig()),
	}
	BindFlags(&opts)
	flag.Usage = UsageFor(os.Args[0] + " [flags]")
	flag.Parse()
	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts), zap.RawZapOpts(zapraw.AddCaller(), zapraw.AddCallerSkip(-1))))

	return configPath

}

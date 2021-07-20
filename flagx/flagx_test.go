package flagx_test

import (
	"errors"
	"flag"
	"go.uber.org/zap/zapcore"
	"os"

	"github.com/w6d-io/x/flagx"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Flags", func() {
	Context("Check function", func() {
		var (
			opts zap.Options
		)
		BeforeEach(func() {
			opts = zap.Options{
				Encoder: zapcore.NewConsoleEncoder(flagx.TextEncoderConfig()),
			}
		})
		It("JsonEncoderConfig", func() {
			Expect(flagx.JsonEncoderConfig()).ToNot(BeNil())
		})
		It("TextEncoderConfig", func() {
			Expect(flagx.TextEncoderConfig()).ToNot(BeNil())
		})
		It("BindFlags", func() {
			flagx.BindFlags(&opts, flag.CommandLine)
		})
	})
	Context("Check flags methods check", func() {
		var (
			opts zap.Options
		)
		BeforeEach(func() {
			opts = zap.Options{
				Encoder: zapcore.NewConsoleEncoder(flagx.TextEncoderConfig()),
			}
		})
		levelFlag := flagx.LevelFlag{}
		levelFlag.ZapOptions = &opts
		When("level flag is used", func() {
			It("Flag is empty", func() {
				Expect(levelFlag.Set("")).Should(Equal(errors.New(`invalid log level ""`)))
			})
			It("invalid string level", func() {
				Expect(levelFlag.Set("no-level")).Should(Equal(errors.New(`invalid log level "no-level"`)))
			})
			It("invalid integer level", func() {
				Expect(levelFlag.Set("-1")).Should(Equal(errors.New(`invalid log level "-1"`)))
			})
			It("valid integer level", func() {
				Expect(levelFlag.Set("1")).To(BeNil())
			})
			It("valid string level", func() {
				Expect(levelFlag.Set("debug")).To(BeNil())
			})
		})
		outputFlag := flagx.OutputFormatFlag{}
		outputFlag.ZapOptions = &opts
		When("output format flag is used", func() {
			It("Flag is empty", func() {
				Expect(outputFlag.Set("")).Should(Equal(errors.New(`invalid ""`)))
			})
			It("invalid format", func() {
				Expect(outputFlag.Set("wrong-format")).Should(Equal(errors.New(`invalid "wrong-format"`)))
			})
			It("valid json format", func() {
				Expect(outputFlag.Set("json")).To(BeNil())
			})
			It("valid text format", func() {
				Expect(outputFlag.Set("text")).To(BeNil())
			})
		})
	})

	Context("lookup env string", func() {
		It("get variable value", func() {
			err := os.Setenv("TEST", "test")
			Expect(err).To(Succeed())
			Expect(flagx.LookupEnvOrString("TEST", "default")).To(Equal("test"))
		})
		It("get default value", func() {
			err := os.Unsetenv("TEST")
			Expect(err).To(Succeed())
			Expect(flagx.LookupEnvOrString("TEST", "default")).To(Equal("default"))
		})
		It("get variable value", func() {
			err := os.Setenv("TEST", "true")
			Expect(err).To(Succeed())
			Expect(flagx.LookupEnvOrBool("TEST", false)).To(Equal(true))
		})
		It("get default value", func() {
			err := os.Unsetenv("TEST")
			Expect(err).To(Succeed())
			Expect(flagx.LookupEnvOrBool("TEST", false)).To(Equal(false))
		})
	})
})

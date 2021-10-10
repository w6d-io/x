module github.com/w6d-io/x

go 1.16

replace github.com/miekg/dns => github.com/miekg/dns v1.1.25

require (
	github.com/go-kit/kit v0.12.0
	github.com/go-logr/logr v0.4.0
	github.com/google/uuid v1.3.0
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.14.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	go.uber.org/zap v1.19.1
	k8s.io/client-go v0.22.2 // indirect
	sigs.k8s.io/controller-runtime v0.9.2
)

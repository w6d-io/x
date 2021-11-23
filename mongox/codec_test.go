// +build !integration

package mongox_test

import (
	"errors"
	"reflect"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"google.golang.org/protobuf/types/known/timestamppb"

	. "github.com/w6d-io/x/mongox"
)

var _ = Describe("Codec", func() {
	Context("", func() {
		When("", func() {
			It("proto codec", func() {
				ts := timestamppb.New(time.Now())
				err := ProtoCodecFunc(bsoncodec.EncodeContext{}, &MockbsonWriter{}, reflect.ValueOf(ts))
				Expect(err).To(Succeed())
			})
			It("proto codec wront input type", func() {
				ts := 1
				err := ProtoCodecFunc(bsoncodec.EncodeContext{}, &MockbsonWriter{}, reflect.ValueOf(ts))
				Expect(err).NotTo(Succeed())
			})
			It("proto decodec", func() {
				ts := timestamppb.New(time.Now())
				err := ProtoDeCodecFunc(bsoncodec.DecodeContext{}, &MockbsonRead{}, reflect.ValueOf(ts))
				Expect(err).To(Succeed())
			})
			It("proto decodec error read time", func() {
				ts := timestamppb.New(time.Now())
				err := ProtoDeCodecFunc(bsoncodec.DecodeContext{}, &MockbsonRead{ErrReadDateTime: errors.New("error read time")}, reflect.ValueOf(ts))
				Expect(err).NotTo(Succeed())
			})
		})
	})
})

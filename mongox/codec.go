package mongox

import (
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var tsProto = timestamppb.New(time.Now())
var tsProtoType = reflect.TypeOf(tsProto)

var tsStr = ""
var tsStrType = reflect.TypeOf(tsStr)

// StrCodecFunc is the String Encoding Function
var StrCodecFunc = func(_ bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != tsStrType {
		return bsoncodec.ValueEncoderError{Name: "ObjectIDEncodeValue", Types: []reflect.Type{tsProtoType}, Received: val}
	}
	s := val.Interface().(string)
	t, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		return vw.WriteString(s)
	}
	return vw.WriteDateTime(t.Unix() * 1000)
}

// ProtoCodecFunc is the Proto Encoding Function
var ProtoCodecFunc = func(_ bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != tsProtoType {
		return bsoncodec.ValueEncoderError{Name: "ObjectIDEncodeValue", Types: []reflect.Type{tsProtoType}, Received: val}
	}
	s := val.Interface().(*timestamppb.Timestamp)
	// WriteDateTime takes milli seconds as argument
	return vw.WriteDateTime(s.Seconds * 1000)
}

// ProtoDeCodecFunc is the Proto Dencoding Function
var ProtoDeCodecFunc = func(_ bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	// this is the function when we read the datetime format
	read, err := vr.ReadDateTime()
	if err != nil {
		return err
	}
	// Read back value as milli seconds, convert it to seconds.
	val.Interface().(*timestamppb.Timestamp).Seconds = read / 1000
	return nil
}

// CodecRegistry create a new bsoncode
func CodecRegistry(options *Options) *bsoncodec.RegistryBuilder {

	var primitiveCodecs bson.PrimitiveCodecs
	rb := bsoncodec.NewRegistryBuilder()
	if options.ProtoCodec {
		bsoncodec.DefaultValueEncoders{}.RegisterDefaultEncoders(rb)
		bsoncodec.DefaultValueDecoders{}.RegisterDefaultDecoders(rb)
		rb.RegisterTypeEncoder(
			tsProtoType,
			bsoncodec.ValueEncoderFunc(ProtoCodecFunc),
		)

		rb.RegisterTypeDecoder(
			tsProtoType,
			bsoncodec.ValueDecoderFunc(ProtoDeCodecFunc),
		)
	}
	if options.StrCodec {
		bsoncodec.DefaultValueDecoders{}.RegisterDefaultDecoders(rb)
		rb.RegisterTypeEncoder(
			tsStrType,
			bsoncodec.ValueEncoderFunc(StrCodecFunc),
		)
	}
	primitiveCodecs.RegisterPrimitiveCodecs(rb)
	return rb
}

package mongox

import (
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var ts = timestamppb.New(time.Now())
var tsProtoType = reflect.TypeOf(ts)

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

// ProtoCodecRegistry create a new bsoncodec with proto support
func ProtoCodecRegistry() *bsoncodec.RegistryBuilder {

	var primitiveCodecs bson.PrimitiveCodecs
	rb := bsoncodec.NewRegistryBuilder()
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
	primitiveCodecs.RegisterPrimitiveCodecs(rb)
	return rb
}

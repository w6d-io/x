package kafkax

// Credit to https://github.com/riferrei/srclient/issues/17

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"strings"

	"google.golang.org/protobuf/proto"

	"github.com/w6d-io/x/logx"
)

var (
	ErrWrongInputType                  = errors.New("Wrong Input Type")
	ErrWhileSerializingProtobufMessage = errors.New("Failure while deserializing Protobuf Message")
	ErrDecodingMessageIndexArray       = errors.New("Unable to decode message index array")
)

func (e Event) Marshall(v interface{}) ([]byte, error) {

	// Check message is a proto
	pb, ok := v.(proto.Message)
	if !ok {
		return json.Marshal(v)
	}

	if e.SchemaRegistry != nil {
		schema, err := e.SchemaRegistry.GetLatestSchema(e.Topic)
		if err != nil {
			return nil, err
		}
		if schema != nil {
			// Check Schema format from content
			// Todo: let see if srclient lib will be improved to return SchemaType ...
			if strings.HasPrefix(strings.ReplaceAll(schema.Schema, " ", ""), "syntax=\"proto3\"") {
				value, _, err := serialize(*schema, pb)
				if err != nil {
					return nil, err
				}
				return value, nil
			}
		}
	}
	return proto.Marshal(pb)

}

func serialize(schema Schema, pb proto.Message) ([]byte, int, error) {
	log := logx.WithName(nil, "SchemaRegistry.Serialize")
	schemaIDBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(schemaIDBytes, uint32(schema.Id))

	// 10 bytes is sufficient for 64 bits in zigzag encoding, but array indexes cannot possibly be
	// 64 bits in a reasonable protobuf, so let's just make a buffer sufficient for a reasonable
	// array of 4 or 5 elements, each with relatively small index
	// varBytes := make([]byte, 16)
	// // array length 1
	// length := binary.PutVarint(varBytes, 1)
	// // index 0 element.  We could write array length 0 with no subsequent value, which is equivalent to writing 1, 0
	// length += binary.PutVarint(varBytes[length:], 0)

	// See https://docs.confluent.io/platform/current/schema-registry/serdes-develop/index.html#wire-format
	// We assume "... Also since most of the time the actual message type will be just the first message type ..."
	varBytes := make([]byte, 1)
	length := 1
	msgIndexBytes := varBytes[:length]

	bytes, err := proto.Marshal(pb)
	if err != nil {
		log.Error(ErrWhileSerializingProtobufMessage, "failed serialize", "err", err.Error())
		return nil, -1, ErrWhileSerializingProtobufMessage
	}

	var msgBytes []byte
	// schema serialization protocol version number
	msgBytes = append(msgBytes, byte(0))
	// schema id
	msgBytes = append(msgBytes, schemaIDBytes...)
	// zig zag encoded array of message indexes preceded by length of array
	msgBytes = append(msgBytes, msgIndexBytes...)

	offset := len(msgBytes)
	log.Info("msgBytes before proto", "len", len(msgBytes))
	msgBytes = append(msgBytes, bytes...)

	return msgBytes, offset, nil
}

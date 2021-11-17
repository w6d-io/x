package kafkax

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/jhump/protoreflect/dynamic"
)

type protoInfo struct {
	m   *dynamic.Message
	fqn string
}

var (
	cacheLock                        = sync.RWMutex{}
	cache                            = make(map[string]protoInfo)
	ErrProtoMarshallerNoElementFound = errors.New("No element found in proto")
)

func (e Event) Unmarshal(d map[string]interface{}) error {

	if e.SchemaRegistry != nil {
		schema, err := e.SchemaRegistry.GetLatestSchema(e.Topic)
		if err != nil {
			return err
		}
		offset, err := deserialize(e.Value)
		if err != nil {
			return err
		} else {
			e.Value = e.Value[offset:]
			cacheLock.RLock()
			_, ok := cache[e.Topic]
			if !ok {
				path := filepath.Join(os.TempDir(), e.Topic)
				if err := ioutil.WriteFile(path, []byte(schema.Schema), 0600); err != nil {
					return err
				}
				var parser protoparse.Parser
				fd, err := parser.ParseFiles(path)
				if err != nil {
					cacheLock.RUnlock()
					// Let see if it's a json file
					goto default_unmarshall
				}
				mt := fd[0].GetMessageTypes()
				md := fd[0].FindMessage(mt[0].GetFullyQualifiedName())
				cache[e.Topic] = protoInfo{
					m:   dynamic.NewMessage(md),
					fqn: mt[0].GetName(),
				}
			}
			cacheLock.RUnlock()

			m := cache[e.Topic].m
			if err := proto.Unmarshal(e.Value, m); err != nil {
				return err
			}

			err = convertProtoToMap(
				cache[e.Topic].fqn,
				m,
				m.GetMessageDescriptor().GetFields(),
				d,
			)
			return err
		}
	}

default_unmarshall:
	// Default format is Json
	return json.Unmarshal(e.Value, &d)
}

func deserialize(bytes []byte) (int, error) {
	// decode the number of elements in the array of message indexes
	arrayLen, bytesRead := binary.Varint(bytes[5:])
	if bytesRead <= 0 {
		return -1, ErrDecodingMessageIndexArray
	}
	totalBytesRead := bytesRead
	// msgIndexArray := make([]int64, arrayLen)
	// iterate arrayLen times, decoding another varint
	for i := int64(0); i < arrayLen; i++ {
		// idx, bytesRead := binary.Varint(bytes[5+totalBytesRead:])
		_, bytesRead := binary.Varint(bytes[5+totalBytesRead:])
		if bytesRead <= 0 {
			return -1, ErrDecodingMessageIndexArray
		}
		totalBytesRead += bytesRead
		// msgIndexArray[i] = idx
	}
	// we have msgIndexArray populated.  If we had a parsed schema, we could
	// look up the actual message type with that.  Or use it as a key into a
	// table of closures, each of which returns proto.Message of the correct
	// type.  Then unmarshall into that and return it. User can cast proto.Message
	// to the actual type safely. But for now, we simply deserialize into the
	// proto that was passed in.
	// err := proto.Unmarshal(bytes[5+totalBytesRead:], pb)
	// if err != nil {
	// 	fmt.Printf("failed deserialize: %v", err)
	// 	return err
	// }
	return 5 + totalBytesRead, nil
}

// Retrieve Field Value From Path
func retrieveFieldValueByPath(path []string, m *dynamic.Message) (interface{}, error) {

	desc := m.GetMessageDescriptor().FindFieldByJSONName(path[1])

	if desc == nil {
		return nil, ErrProtoMarshallerNoElementFound
	}

	if len(path) == 2 {
		return m.TryGetField(desc)
	}

	msg, err := m.TryGetField(desc)
	if err != nil {
		return nil, err
	}
	for i := 2; i < len(path)-1; i++ {
		desc = msg.(*dynamic.Message).GetMessageDescriptor().FindFieldByJSONName(path[i])
		if desc == nil {
			return nil, ErrProtoMarshallerNoElementFound
		}
		msg, err = msg.(*dynamic.Message).TryGetField(desc)
		if err != nil {
			return nil, err
		}
	}

	desc = msg.(*dynamic.Message).GetMessageDescriptor().FindFieldByJSONName(path[len(path)-1])
	if desc == nil {
		return nil, ErrProtoMarshallerNoElementFound
	}

	return msg.(*dynamic.Message).TryGetField(desc)
}

// Convert Proto To Map
func convertProtoToMap(fqn string, m *dynamic.Message, desc []*desc.FieldDescriptor, m1 map[string]interface{}) error {

	for _, d := range desc {

		v, err := retrieveFieldValueByPath(
			strings.Split(fqn+"."+d.GetJSONName(), "."),
			m,
		)
		if err != nil {
			return err
		}
		ret, ok := v.(*dynamic.Message)
		if ok {
			if ret != nil {
				tmp := make(map[string]interface{})
				err := convertProtoToMap(
					fqn+"."+d.GetJSONName(),
					m,
					ret.GetMessageDescriptor().GetFields(),
					tmp,
				)
				if err != nil {
					return err
				}
				m1[d.GetJSONName()] = tmp
			}
		} else {
			m1[d.GetJSONName()] = v
		}
	}

	return nil
}

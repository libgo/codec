package proto

import (
	"google.golang.org/protobuf/proto"

	"github.com/libgo/codec"
)

// Name is the name registered for the proto codec.
const Name = "proto"

func init() {
	codec.RegisterCodec(codecImpl{})
}

var c = codecImpl{}

// codecImpl is a Codec implementation with protobuf.
type codecImpl struct{}

func (codecImpl) Marshal(v interface{}) ([]byte, error) {
	return proto.Marshal(v.(proto.Message))
}

func (codecImpl) Unmarshal(data []byte, v interface{}) error {
	return proto.Unmarshal(data, v.(proto.Message))
}

func (codecImpl) Name() string {
	return Name
}

package proto

import "google.golang.org/protobuf/proto"

func Marshal(v interface{}) ([]byte, error) {
	return c.Marshal(v.(proto.Message))
}

func Unmarshal(data []byte, v interface{}) error {
	return c.Unmarshal(data, v.(proto.Message))
}

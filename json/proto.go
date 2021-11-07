package json

import (
	"bytes"
	"reflect"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
)

const (
	nullStr   = "null"
	objectStr = "{}"
	listStr   = "[]"
)

var (
	nullBytes   = []byte(nullStr)
	objectBytes = []byte(objectStr)
	listBytes   = []byte(listStr)
)

func MarshalProtoArray(v interface{}) ([]byte, error) {
	rv := reflect.Indirect(reflect.ValueOf(v))
	if rv.IsNil() {
		return []byte(nullStr), nil
	}

	buf := bytes.NewBuffer(nil)
	buf.WriteByte('[')

	l := rv.Len()

	for i := 0; i < l; i++ {
		m := rv.Index(i).Interface().(proto.Message)
		data, err := Marshal(m)
		if err != nil {
			return nil, err
		}
		if buf.Len() > 1 {
			buf.WriteByte(',')
		}
		buf.Write(data)
	}

	buf.WriteByte(']')

	return buf.Bytes(), nil
}

func MustMarshalProtoArray(v interface{}) []byte {
	data, err := MarshalProtoArray(v)
	if err != nil {
		panic(err)
	}
	return data
}

func UnmarshalProtoArray(data []byte, v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr {
		panic("v must be pointer")
	}

	e := rv.Elem()

	if data == nil || len(data) == 0 || bytes.Equal(data, nullBytes) {
		e.Set(reflect.Zero(e.Type()))
		return nil
	}

	// TODO optimize, maybe we can use something like json slice object reader
	lv := []map[string]interface{}{}
	err := Unmarshal(data, &lv)
	if err != nil {
		return err
	}

	slice := reflect.MakeSlice(e.Type(), len(lv), len(lv))
	e.Set(slice)

	et := e.Type().Elem()

	for i := range lv {
		data := MustMarshal(lv[i])

		if e.Index(i).IsNil() {
			e.Index(i).Set(reflect.New(et.Elem()))
		}

		err = Unmarshal(data, slice.Index(i).Interface().(proto.Message))
		if err != nil {
			return err
		}
	}

	return nil
}

// UnmarshalToProto is helper func to convert bytes to related proto.Message.
func UnmarshalToProto(data []byte, m proto.Message) (proto.Message, error) {
	if data == nil || len(data) == 0 || bytes.Equal(data, nullBytes) {
		return reflect.Zero(reflect.TypeOf(m)).Interface().(proto.Message), nil
	}
	err := Unmarshal(data, m)
	if err != nil {
		return reflect.Zero(reflect.TypeOf(m)).Interface().(proto.Message), err
	}
	return m, nil
}

// MustUnmarshalToProto is helper func to convert bytes to related proto.Message without error.
func MustUnmarshalToProto(data []byte, m proto.Message) proto.Message {
	m, _ = UnmarshalToProto(data, m)
	return m
}

func UnmarshalToProtoStruct(data []byte) (*structpb.Struct, error) {
	if len(data) == 0 || bytes.Equal(data, nullBytes) {
		return nil, nil
	}

	s := &structpb.Struct{}
	if bytes.Equal(data, objectBytes) {
		return s, nil
	}

	err := Unmarshal(data, s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func MustUnmarshalToProtoStruct(data []byte) *structpb.Struct {
	s, err := UnmarshalToProtoStruct(data)
	if err != nil {
		panic(err)
	}
	return s
}

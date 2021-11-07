package json

func Marshal(v interface{}) ([]byte, error) {
	return c.Marshal(v)
}

func MustMarshal(v interface{}) []byte {
	data, err := Marshal(v)
	if err != nil {
		panic(err)
	}
	return data
}

func MustMarshalString(v interface{}) string {
	data, err := Marshal(v)
	if err != nil {
		panic(err)
	}

	return string(data)
}

func Unmarshal(data []byte, v interface{}) error {
	return c.Unmarshal(data, v)
}

func MustUnmarshal(data []byte, v interface{}) interface{} {
	err := Unmarshal(data, v)
	if err != nil {
		panic(err)
	}
	return v
}

package json

import "github.com/tidwall/gjson"

type Result = gjson.Result

func Get(jsonStr string, path string) Result {
	return gjson.Get(jsonStr, path)
}

func GetBytes(jsonData []byte, path string) Result {
	return gjson.GetBytes(jsonData, path)
}

func Parse(jsonStr string) Result {
	return gjson.Parse(jsonStr)
}

func ParseBytes(jsonData []byte) Result {
	return gjson.ParseBytes(jsonData)
}

func Valid(jsonStr string) bool {
	return gjson.Valid(jsonStr)
}

func ValidBytes(jsonData []byte) bool {
	return gjson.ValidBytes(jsonData)
}

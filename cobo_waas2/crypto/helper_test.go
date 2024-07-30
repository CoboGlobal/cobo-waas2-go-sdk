package crypto

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonUnmarshal1(t *testing.T) {
	jsonString1 := `{"key": "value"}`
	var data1 map[string]json.RawMessage
	err := json.Unmarshal([]byte(jsonString1), &data1)
	if err != nil {
		fmt.Println("解析JSON字符串1出错:", err)
		return
	}
	t.Log(data1)

	jsonString2 := `["item1", "item2"]`
	var data2 []json.RawMessage
	err = json.Unmarshal([]byte(jsonString2), &data2)
	if err != nil {
		fmt.Println("解析JSON字符串1出错:", err)
		return
	}
	t.Log(data2)

	jsonString3 := `["item1", "item2"]`
	var data3 map[string]json.RawMessage
	err = json.Unmarshal([]byte(jsonString3), &data3)
	if err != nil {
		fmt.Println("解析JSON字符串1出错:", err)
		return
	}
	t.Log(data2)
}

func TestJsonUnmarshal2(t *testing.T) {
	jsonString1 := `{"key": {"key1": "value"}}`
	var data1 map[string]json.RawMessage
	err := json.Unmarshal([]byte(jsonString1), &data1)
	if err != nil {
		fmt.Println("parse json string err:", err)
		return
	}
	v1, err := data1["key"].MarshalJSON()
	fmt.Println(string(v1), err)
	t.Log(data1)
}

func TestJsonUnmarshal3(t *testing.T) {
	jsonString1 := `{"key": "value"}`
	var data1 map[string]json.RawMessage
	err := json.Unmarshal([]byte(jsonString1), &data1)
	if err != nil {
		fmt.Println("parse json string err:", err)
		return
	}
	v1, err := data1["key"].MarshalJSON()
	fmt.Println(string(v1), string(data1["key"]))
	t.Log(data1)
}

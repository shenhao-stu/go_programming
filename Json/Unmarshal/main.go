package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Birthday string
	Sal      float64
}

func unmarshalStruct() {
	str := `{"Name":"Tom","Age":23,"Birthday":"2000-1-1","Sal":8000}`

	var per Person

	err := json.Unmarshal([]byte(str), &per)
	if err != nil {
		fmt.Printf("反序列化错误 err=%v\n", err)
	}
	fmt.Printf("反序列化后 per=%v\n", per)
}

func unmarshalMap() {
	str := `{"name":"tom","age":23,"birthday":"2000-1-1","sal":8000}`
	var m map[string]interface{}

	// tips: 反序列化map，不需要make，因为make操作呗封装到 Unmarshal 函数中
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		fmt.Printf("反序列化错误 err=%v\n", err)
	}
	fmt.Printf("反序列化后 m=%v\n", m)
}

func unmarshalSlice() {
	str := `[{"name":"tom","age":23,"birthday":"2000-1-1","sal":8000},{"name":"jack","age":24,"birthday":"2000-1-1","sal":8000}]`
	var slice []map[string]interface{}
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("反序列化错误 err=%v\n", err)
	}
	fmt.Printf("反序列化后 slice=%v\n", slice)
}

func main() {
	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()
}

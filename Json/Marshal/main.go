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

func testMap() {
	var a = make(map[string]interface{})
	a["name"] = "tom"
	a["age"] = 23
	a["birthday"] = "2000-1-1"

	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Printf("map a序列化后=%v\n", string(data))
}

func testSlice() {
	var slice []map[string]interface{}
	slice = append(slice, map[string]interface{}{"name": "tom", "age": 23})
	m := make(map[string]interface{})
	m["name"] = "jack"
	m["age"] = 24
	slice = append(slice, m)

	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Printf("slice序列化后=%v\n", string(data))

}

func testStruct() {
	per := Person{
		Name:     "Tom",
		Age:      23,
		Birthday: "2000-1-1",
		Sal:      8000.0,
	}
	// 使用引用传递，更快速，不需要进行copy
	data, err := json.Marshal(&per)
	if err != nil {
		fmt.Printf("序列号错误 err=%v\n", err)
	}
	fmt.Printf("person序列化后=%v\n", string(data))
}
func main() {
	testStruct()
	testMap()
	testSlice()
}

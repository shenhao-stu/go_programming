package demo

import (
	"reflect"
	"testing"
)

type user struct {
	UserId string
	Name   string
}

// 使用反射操作任意结构体类型
func TestReflectStruct(t *testing.T) {
	var (
		model *user
		sv    reflect.Value
	)
	model = &user{}
	sv = reflect.ValueOf(model)
	t.Log("reflect.ValueOf", sv.Kind().String())
	sv = sv.Elem()
	t.Log("reflect.ValueOf.Elem", sv.Kind().String())
	sv.FieldByName("UserId").SetString("123")
	sv.FieldByName("Name").SetString("tom")
	t.Log("model", *model)

}

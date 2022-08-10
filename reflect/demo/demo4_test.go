package demo

import (
	"reflect"
	"testing"
)

type User struct {
	UserId string
	Name   string
}

// 使用反射创建并操作结构体
func TestReflectStructPtr(t *testing.T) {
	var (
		model *User
		st    reflect.Type
		elem  reflect.Value
	)
	st = reflect.TypeOf(model)                             // 获取类型*User
	t.Log("reflect.TypeOf", st.Kind().String())            // ptr
	st = st.Elem()                                         // st指向的类型
	t.Log("reflect.TypeOf.Elem", st.Kind().String())       // struct
	elem = reflect.New(st)                                 // New返回一个Value类型值，该值持有一个指向类型为st的新申请的零值的指针
	t.Log("reflect.New", elem.Kind().String())             // ptr
	t.Log("reflect.New.Elem", elem.Elem().Kind().String()) // struct
	// model就是创建的user结构体变量
	model = elem.Interface().(*User)
	elem = elem.Elem() // 取得elem指向的值
	elem.FieldByName("UserId").SetString("123")
	elem.FieldByName("Name").SetString("tom")
	t.Log("model.Name", model.Name)
}

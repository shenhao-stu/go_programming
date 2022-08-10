package demo

import (
	"reflect"
	"testing"
)

// 适配器函数用作统一处理接口
func TestReflectFunc(t *testing.T) {
	call1 := func(v1 int, v2 int, s string) {
		t.Log(v1, v2, s)
	}

	call2 := func(v1 int, v2 int) {
		t.Log(v1, v2)
	}

	var (
		function reflect.Value
		inValues []reflect.Value
		n        int
	)
	bridge := func(call interface{}, args ...interface{}) {
		n = len(args)
		inValues = make([]reflect.Value, n)
		for i := 0; i < n; i++ {
			inValues[i] = reflect.ValueOf(args[i])
			t.Logf("inValues[%d] = %v", i, inValues[i])
		}
		function = reflect.ValueOf(call)
		// t.Log(inValues, function)
		function.Call(inValues)
	}

	bridge(call1, 1, 2, "test1")
	bridge(call2, 1, 2)
}

package file

import (
	"fmt"
	"io/ioutil"
	"testing"
)

// 将一个文件的内容导入另一个文件
// 两个文件已存在
func TestWRFile(t *testing.T) {
	src_file := "demo.txt"
	dst_file := "test.txt"

	data, err := ioutil.ReadFile(src_file)
	if err != nil {
		fmt.Println("read file error: ", err)
	}
	err = ioutil.WriteFile(dst_file, data, 0666)
	if err != nil {
		fmt.Println("write file error: ", err)
	}
}

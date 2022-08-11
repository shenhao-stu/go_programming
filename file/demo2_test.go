package file

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func WriteFile(path string) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file error: ", err)
	}
	defer file.Close()
	// 写入内容
	str := "hello world\n"
	// 写入是使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}

	// 因为writer是带缓存的，因此在调用WriterString方法时
	// 其实内容是先写入到缓存的，所以需要调用Flush方法，将缓存的内容全部写入到文件中
	// 否则文件中会没有数据！！
	writer.Flush()
}
func TestWriteFile(t *testing.T) {
	WriteFile("demo.txt")
}

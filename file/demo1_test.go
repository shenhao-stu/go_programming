package file

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

// 读取文件的内容并显示在终端（带缓冲区的方式）

func ReadBufferFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("open file error: ", err)
	}
	defer file.Close()

	// 创建一个 *Reader，它提供了带缓冲的文件读取。
	// 默认缓冲区大小是 4096。
	reader := bufio.NewReader(file)

	for {
		// ReadString方法从文件中读取一个字符串，直到遇到换行符（\n）或者EOF。
		// 如果遇到错误，则返回错误。
		line, err := reader.ReadString('\n')
		// io.EOF 表示文件的末尾
		// 因此文件最后一行必须要换行符，否则最后一行内容会丢失，ReadFile则会全部读取
		if err == io.EOF {
			break
		}
		// ReadString方法会包括换行符（\n）。
		fmt.Print(line)

	}
	fmt.Println("文件读取完毕")
}

// 使用ioutil一次将整个文件读入到内存中

func ReadFileOnce(path string) {
	// content []byte
	// 文件的open和close被封装到ReadFile函数中
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("read file error: ", err)
	}
	fmt.Println(string(content))
}

func TestReadFile(t *testing.T) {
	ReadBufferFile("test.txt")
	ReadFileOnce("test.txt")
}

package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
)

// copy file
// func Copy(dst Writer, src Reader) (written int64, err error)

func CopyFile(dst string, src string) (written int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Printf("open file error: %v\n", err)
	}
	defer srcFile.Close()
	// 通过srcFile，获取 Reader
	reader := bufio.NewReader(srcFile)

	dstFile, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("open file error: %v\n", err)
	}
	defer dstFile.Close()
	// 通过dstFile，获取 Writer
	writer := bufio.NewWriter(dstFile)

	return io.Copy(writer, reader)
}

func TestCopyFile(t *testing.T) {
	_, err := CopyFile("copy_pic.png", "pic.png")
	if err == nil {
		fmt.Println("copy file success")
	} else {
		fmt.Println("copy file error: ", err)
		return
	}
}

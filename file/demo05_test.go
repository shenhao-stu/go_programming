package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
)

// 统计英文、数字、空格和其他字符数量

type CharCount struct {
	EngCount   int
	NumCount   int
	SpaceCount int
	OtherCount int
}

func Count(path string) CharCount {
	var charcount CharCount

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("open file error: ", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		res, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		str := []rune(res)
		for _, char := range str {
			switch {
			case char >= 'a' && char <= 'z':
				fallthrough
			case char >= 'A' && char <= 'Z':
				charcount.EngCount++
			case char >= '0' && char <= '9':
				charcount.NumCount++
			case char == ' ' || char == '\t':
				charcount.SpaceCount++
			default:
				charcount.OtherCount++
			}
		}
	}
	return charcount
}

func TestCount(t *testing.T) {
	charcount := Count("count.txt")
	fmt.Println(charcount)
}

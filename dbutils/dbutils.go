package dbutils

import "fmt"

// Initial capitalization can be accessed by other packages
func GetConnect(str string) {
	fmt.Println("dbutils package: ", str)
}

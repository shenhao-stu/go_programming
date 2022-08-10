package main

import (
	"flag"
	"fmt"
)

// os.Args[0] is the path to the executable.
// os.Args[1] is the first argument to the program.
// os.Args[2] is the second argument to the program.
// ...
// os.Args[n] is the nth argument to the program.

// go build -o test flag/main.go
// ./test -u root -p root -h localhost -port 3306
func main() {
	var user string
	var pwd string
	var host string
	var port int

	flag.StringVar(&user, "u", "", "用户名 默认为空")
	flag.StringVar(&pwd, "p", "", "密码 默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名 默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口 默认为3306")

	flag.Parse()

	fmt.Printf("user=%v pwd=%v host=%v port=%v\n", user, pwd, host, port)
}

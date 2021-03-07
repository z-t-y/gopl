// echo4输出其命令行参数
package main

import (
	"flag"
	"fmt"
	"strings"
)

// 变量n和sep是指向标识变量的指针
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

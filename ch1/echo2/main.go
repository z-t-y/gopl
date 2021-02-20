// echo2输出其命令行参数
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] { // range每次产生一对值：索引和这个索引元素的值
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}


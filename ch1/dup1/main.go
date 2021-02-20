package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// 忽略input.Error()中的错误
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\n%s\n", n, line)
		}
	}
}

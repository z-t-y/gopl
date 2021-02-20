package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	time1 := time.Now().UTC().UnixNano()
	// less effective code
	s := ""
	for i := 97; i < 102; i++ {
		s += string(rune(i)) + " "
	}
	fmt.Println(s)
	time2 := time.Now().UTC().UnixNano()
	l2 := []string{"a", "b", "c", "d", "e"}
	s2 := strings.Join(l2, " ")
	fmt.Println(s2)
	time3 := time.Now().UTC().UnixNano()
	fmt.Println(time2 - time1)
	fmt.Println(time3 - time2)
}

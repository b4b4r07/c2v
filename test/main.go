package main

import (
	"fmt"
	"strings"
)

func main() {
	a := ` I'm 10 yaers old. ;-)`

	fmt.Println(strings.Split(a, " ").TrimSpace())
}

package main

import (
	"strings"
)

func main() {

	println(strings.Repeat("a", 0) + strings.Repeat("a", 1) + strings.Repeat("a", 0) + strings.Repeat("a", 2) + strings.Repeat("a", 3))
}

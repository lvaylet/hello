package main

import (
	"fmt"

	"github.com/user/stringutil"
)

// T is only used to test gofmt
type T struct {
	name  string // name of the object
	value int    // its value
}

func main() {
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
}

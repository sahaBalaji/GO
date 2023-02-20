package main

import (
	"fmt"
	"strings"
)

func joinstr(elements ...string) string {
	return strings.Join(elements, "-")
}

func main() {
	fmt.Println(joinstr())
	fmt.Println(joinstr("Hi\n", "How\tare\tyou?"))
	fmt.Println(joinstr("I", "am", "fine"))

}

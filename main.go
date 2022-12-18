package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var count int
	for _, value := range os.Args[1:] {
		strcount := strings.Split(value, " ")
		for _, strvalue := range strcount {
			if len(strvalue) != 0 {
				count++
			}
		}
	}
	fmt.Println(count)
}

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("arg not a valid")
	} else {
		var raw = strings.Trim(os.Args[1], " ")
		fmt.Println(raw)
		material := strings.Split(raw, "=")
		if len(material) != 2 {
			fmt.Println("arg not a valid")
			return
		}
		key := material[0]
		value := material[1]
		key = strings.Trim(key, " ")
		value = strings.Trim(value, " ")
		fmt.Println(key)
		fmt.Println(value)
		os.Setenv(key, value)
	}
}

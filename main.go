package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	h      bool
	remove bool
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.BoolVar(&remove, "remove", false, "remove environment variable")
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, `syntax: commands separated by whitespace
Usage: gexport [-remove] var
Options:
- remove  remove environment variable
`)
}

func setEnv(raw string) {
	raw = strings.Trim(raw, " ")
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
	os.Setenv(key, value)
	fmt.Println("set env success")
}

func unsetEnv(key string) {
	key = strings.Trim(key, " ")
	fmt.Println(key)
	os.Unsetenv(key)
	fmt.Println("unset env success")
}

func main() {
	flag.Parse()
	if h {
		flag.Usage()
		return
	}

	if len(flag.Args()) != 1 {
		fmt.Println("arg not a valid")
		return
	}

	if remove {
		var raw = flag.Args()[0]
		unsetEnv(raw)
	} else {
		var raw = flag.Args()[0]
		setEnv(raw)
	}
}

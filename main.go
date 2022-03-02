package main

import (
	"flag"
	"fmt"
)

const ABC = 100

func main() {
	var name string
	flag.StringVar(&name, "name", "golang", "名字")
	flag.Parse()
	fmt.Printf("hello %v", name)
}

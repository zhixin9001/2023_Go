package main

import (
	"fmt"
	"flag"
)

var name string

func init(){
	flag.StringVar(&name, "name", "everyone", "The greet object")
}

func main(){
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}
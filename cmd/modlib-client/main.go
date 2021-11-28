package main

import (
	"fmt"
	"github.com/coledrain/modlib"
	"github.com/coledrain/modlib/clientlib"
)

func main() {
	fmt.Println("Running client")
	fmt.Println("Config:", modlib.Config())
	fmt.Println(clientlib.Hello())
}

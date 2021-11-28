package main

import (
	"fmt"
	"github.com/coledrain/modlib"
	"github.com/coledrain/modlib/internal/auth"
	"github.com/coledrain/modlib/serverlib"
)

func main() {
	fmt.Println("Running Server")
	fmt.Println("Config:", modlib.Config())
	fmt.Println("Auth:", auth.GetAuth())
	fmt.Println(serverlib.Hello())
}

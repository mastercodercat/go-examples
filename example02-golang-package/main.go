package main

import (
	"fmt"

	"github.com/go-training/training/example02-golang-package/controller"
)

func main() {
	fmt.Println("Hello World")

	hi := controller.HelloWorld("mastercodercat")
	fmt.Println(hi)
}

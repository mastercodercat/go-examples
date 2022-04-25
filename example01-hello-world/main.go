package main

import "fmt"

func main() {
	a := 1
	fmt.Printf(HelloWorld("mastercodercat"))
	fmt.Println("Hello World")

	if (a >= 1) { fmt.Println("a >= 1") }
}

func HelloWorld(user_name string) string {
	return fmt.Sprintf("Hi, %s ", user_name)
}

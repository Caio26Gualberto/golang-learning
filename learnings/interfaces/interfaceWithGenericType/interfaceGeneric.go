package main

import "fmt"

func main() {
	generic("String")
	generic(1)
	generic(true)
}

func generic(interf interface{}) {
	fmt.Println(interf)
}

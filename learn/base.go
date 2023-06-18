package learn

// Path: learn\base.go
// this is my first go program
/*
this some line comment
多行注释
anthor line comment
date 2019-11-11
discrption: this is a test program

*/

import (
	"fmt"
)

func HelloWorld(a string) string {
	fmt.Println("Hello World")
	return a + " World"
}

func add(a int, b int) int {
	return a + b
}

func multiply(a int, b int) int {
	return a * b
}

package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	fmt.Println("helloworld!!!")
	fmt.Println(math.Floor(2.75))
	fmt.Println(strings.Title("hello \n\tworld!!!"))
	fmt.Println("\"\"")
	fmt.Println("\\")
	fmt.Println('B') //rune
	fmt.Println(reflect.TypeOf(42))
	fmt.Println(reflect.TypeOf(3.14))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("helloworld!!!"))
}

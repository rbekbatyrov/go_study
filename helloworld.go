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
	var variable_int int
	variable_int = 1
	fmt.Println(variable_int)
	var variable_float_1, variable_float_2 float64 = 1.2, 3.4
	fmt.Println(variable_float_1, variable_float_2)
	var variable_hello, variable_world string
	variable_hello = "hello"
	variable_world = "world"
	fmt.Println(strings.Title(variable_hello), strings.Title(variable_world), "!!!")
	var variable_bool bool
	fmt.Println(variable_bool)
}

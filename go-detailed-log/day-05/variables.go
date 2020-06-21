package main

import "fmt"

// declare a variable with the identifier of a to be of type integer with a zero value
// declare a variable with the identifier of b to be of type string  with a zero value
var a int
var b string

func main(){
    // we are assigning the value of 1 to our variable which is defined as type int
    a = 1
    // we are assigning the value of "1" to our variable which is defined as type string
    b = "1"
    fmt.Println(a)
    fmt.Printf("%T\n", a)
    fmt.Println(b)
    fmt.Printf("%T\n", b)
    // once we declared our variable with identifier of a is a integer, we cannot chage the
    // type to string, as it will error out
    // a = "foo" will fail
    a = "foo"
    fmt.Println(a)
    // ./app.go:22:7: cannot use "foo" (type untyped string) as type int in assignment
}

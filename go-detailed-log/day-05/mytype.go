package main

import "fmt"

var a string
type football string
var ball football

func main(){
    ball = "round"
    fmt.Println(ball)
    fmt.Printf("%T\n", ball)
}

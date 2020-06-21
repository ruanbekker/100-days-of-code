package main

import "fmt"

var num int
type hotdog int
var nomnom hotdog

func main(){
    num = 1
    nomnom = 1
    fmt.Println(num)
    fmt.Printf("%T\n", num)
    fmt.Println(nomnom)
    fmt.Printf("%T\n", nomnom)
    //conversion from type hotdog to type which has int as type underlying
    //we can only convert to int as hotdog underlying type is int
    num = int(nomnom)
    fmt.Println(num)
    fmt.Printf("%T\n", num)
}

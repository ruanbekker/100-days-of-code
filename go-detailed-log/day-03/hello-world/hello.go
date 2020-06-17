package main

import "fmt"

func main(){
    fmt.Println("begin")
    count()
    fmt.Println("end")
}

func count(){
    fmt.Println("begin to count")
    for x := 0; x < 10; x++ {
        fmt.Println(x)
    }
    fmt.Println("end to count")
}

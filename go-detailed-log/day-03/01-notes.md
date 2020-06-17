## Day 3 - Notes

### Package Main

Package main is your entry point to your program, the program ends when the flow of execution exits func main.

Control Flow:

1. Sequence
2. Loop or Iterative
3. Conditional

Ref: https://en.wikipedia.org/wiki/Control_flow

### Undeclared Variables

You need to be explicit. If you dont define a variable the compiler will throw out a error.

### Hello World

Hello world application:

```
package main

import "fmt"

func main(){
    fmt.Println("Hello, World!")
}
```

When we call other functions from our main function:

```
package main

import "fmt"

func main(){
    fmt.Println("start")
    doStuff()
    fmt.Println("end")
}

func doStuff(){
    fmt.Println("print some stuff")
}
```

When we run that:

```
$ go run app.go
start
print some stuff
end
```

### Getting information from a Package

For `fmt` go to [godoc.org/fmt](https://godoc.org/fmt) and read up on how to use the package.

To return the size in bytes and catch a error on return if exists:

```
package main

import "fmt"

func main(){
    n, err := fmt.Println("hello", 42, true)
    fmt.Println(n)
    fmt.Println(err)
}
```

To throw away a return:

```
package main

import "fmt"

func main(){
    n, _ := fmt.Println("hello", 42, true)
    fmt.Println(n)
}
```

It tells your compiler to throw it away



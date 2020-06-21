## Static Variables

Go is a static programming language, once you declare that a variable is of type string you cant change the type 
to something else after being declared.

```
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
```

## Types

We can declare our own type, for example `football` and the underlying type can be of type string:

```
package main

import "fmt"

var a string
type football string
var ball football

func main(){
    ball = "round"
    fmt.Println(ball)
    fmt.Printf("%T\n", ball)
    // output:
    // round
    // main.football
}
```

## Conversion

In other languages we call it casting, in go we call it conversion:

```
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
    // output:
    // 1
    // int
    // 1
    // main.hotdog
    // 1
    // int
}
```

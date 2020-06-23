## Day 6 - Notes

Challenges

### Challenge 1

Using the short declaration operator, assign these values to variables with the identifiers x,y,z and the print the values in single print statements and then multiple print statements

42, james bond, true

```
package main

import "fmt"

func main(){
    x := 42
    y := "James Bond"
    z := true
    // multiple print statements
    fmt.Println(x)
    fmt.Println(y)
    fmt.Println(z)
    // single print statements
    fmt.Println(x, y, z)
}
```

## Challenge 2

Declare package level scope variables and dont assign values

```
package main

import "fmt"

// package level scope variables
// deckared as types with zero values
// so they can store values of that type
var x int
var y string
var z bool

func main() {
    fmt.Println(x)
    fmt.Println(y)
    fmt.Println(z)
}
```

What are empty values called?

Zero values

## Challenge 3

Assign values to variables at the package level scope, use Sprintf from the fmt package to print all the values 
in a single line after assigning to to a s variable using the short declaration operator

```
package main

import "fmt"

// package level scope variables
var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
    // lookup docs on usage for fmt package on godoc.org/fmt
    s := fmt.Sprintf("%v, %v, %v", x, y, z)
    fmt.Println(s)
}
```

## Challenge 4

```
package main

import "fmt"

// declare new type called nommer and set the underlying type to int
type nommer int
// declare a new variable with the identifier x and set the type to the new type and
//assign a zero value to your variable
var x nommer

func main() {
    // print the value for variable x
    fmt.Println(x)
    // print the type for varibale x
    fmt.Printf("%T\n", x)
    // assign the value 42 to the variable x with the "=" operator
    x = 42
    // print out the value for variable with the identifier of x using the Println function from the fmt package
    fmt.Println(x)
}
```

## Challenge 5

```
package main

import "fmt"

// declare new type called nommer and set the underlying type to int
type nommer int
// declare a new variable with the identifier x and set the type to the new type and
//assign a zero value to your variable
var x nommer
var y int

func main() {
    // print the value for variable x
    fmt.Println(x)
    // print the type for varibale x
    fmt.Printf("%T\n", x)
    // assign the value 42 to the variable x with the "=" operator
    x = 42
    // print out the value for variable with the identifier of x using the Println function from the fmt package
    fmt.Println(x)
    // use conversion to convert the type of the value stored in x to the underlying type
    // we dont use short declaration operator as we have already declared y
    y = int(x)
    fmt.Println(y)
    fmt.Printf("%T\n", y)
}
```

## Summary

1. The smallest standalone element of a program that expresses some action to be carried out.? statement
2. A combination of one or more explicit values, constants, variables, operators, and functions that the programming language interprets and computes to produce another value? expression
3. The "scope" of a variable is where you can access the variable, eg, write to it or read the value from it
4. A "primitive" data TYPE is one that is built into the language AND/OR just a basic data type which is built into the language
5. In Go, an "int" or "string" is a primitive data TYPE 
6. A "composite" data TYPE allows you to compose together values of other data TYPES
6. When a variable is declared in Go using the "var" keyword, and no VALUE is ASSIGNED to that variable, then the compiler assigns a default value to the variable. This is known as the "zero value"
7. Keywords are words that a reserved for use by the Go programming language; they have to be used in a certain way for a certain purpose.
8. Keywords are sometimes called “reserved words.” (https://golang.org/ref/spec#Keywords)
9 You can’t use a keyword for anything other than its purpose.
10. In “2 + 2” the “+” is the OPERATOR
11. In “2 + 2” the “2”s are OPERANDS 
12. For finding documentation, what is the difference between documentation found at golang.org and godoc.org? godoc.org has standard libraries + 3rd party
13. "package" + "var" is a keyword
14. The entry point for all programs is in func main() which needs to be inside package main
15. The "short declaration operator" can be used anywhere in a program, including at both the package level and at the block level.
16. What are the three words used to describe good package names in the "effective go" document? short, concise, evocative
17. When you see something like "fmt.Println()" this is calling the "Println()" function from the "fmt" package. 
18. An "identifier" is the name assigned to a variable or a function or a constant. 
19. To call a func, variable, or constant from a package, use the "package-dot-identifier" syntax. For example, like this, "fmt.Println()" 
20. What is "idiomatic Go code"? how you should write go code
21. Which character allows you to "throw away returns" or "send returns into the void"? Said another way, which character allows you to tell the compiler that you are not going to use a value returned by a function? _
22. In Go, you cannot have a variable which you do not use.
23. When you see that a func has a parameter of this type "...interface{}" this is called a "variadic parameter" and it means that the func can take as many values of that type as you want to pass in.
24. Every value in Go is also of type "empty interface" which is expressed like this: "interface{}"
25. A statement is an instruction that commands the computer to perform a specified action. Usually statements take up a line in a program.
26. An expression is a combination of one or more explicit values, constants, variables, operators, and functions that the programming language interprets and computes to produce another value. For example, 2+3 is an expression which evaluates to 5.
27. If I wanted to print to a string and then assign that value to a variable, I could use the "func Sprintf()" from the "fmt" package.
28. In Go, you can create your own TYPE
29. We don't say "casting" in Go, we say "conversion"
30. When you create our own TYPE in Go, that TYPE will have an "underlying TYPE".

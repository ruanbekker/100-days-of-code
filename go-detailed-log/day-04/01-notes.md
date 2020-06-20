## Day 4 - Notes

### Terminlogy

Short declaration operator (:=) (declare and assign)
Keywords (reserved keywords like var)
Expression (1+1 is the expression which evaluates to 2)
Operator (+ will be the operator in 1 + 1)
Operand (the 1's will be the operand)
Statement (a := 1+1)
Curly braces: {}
Parentheses: ()
Initialization: Declare and Assign
Scope: Where it exists and accessible (x := 1 <== all the way until its not usable anymore)
In:

```
package main

import "fmt"

func main(){
    x := 1
    fmt.Println(x)
}
````

We are declaring x as a variable and we assing a value to the variable

```
package main

import "fmt"

func main(){
    x := 1
    fmt.Println(x)
    x = 2
    fmt.Println(x)
}
```

We are declaring x as a variable and assign a value (initialization), then assign another value

Remember: we cannot declare a variable without using it, the compiler will error

Important to check the golang spec to understand what you are using


function body, from `{` to `}` in:

```
func main{
    ....
}
```

Difference between `var x = 10` and `x := 10`:

You cannot declare a variable outside a function body using a short declaration operator, if you want to declare a variable outside a function variable you need to use `var x = value`

You cannot do this:

```
package main

import (
	"fmt"
)

x := 10

func main() {
	fmt.Println(x)
}
```

You can do this:

```
package main

import (
	"fmt"
)

var x = 10

func main() {
	fmt.Println(x)
}
```

Declare a variable with the identifier "y", assign a zero value and set the type to int:

```
var y int
var y int = 0 (the same as above)
var y string = ""
var y float = 0.0
```

Printing data types:

```
package main

import "fmt"

var x string = "hello, world"

func main(){
    fmt.Println(x)
    fmt.Printf("%T\n", x)
}

// returns
// hello, world
// string
```

Go is a STATIC programming language (not dynamic), once you declared a data type as string, you cant change the type to the declared variable as int 

A VARIABLE is DECLARED to hold a VALUE of a certain TYPE

DECLARE variable of TYPE string with IDENTIFIER of x and ASSIGNING a VALUE of ""

String literal: `var a string = "a"`

Raw string literal: 

```
var a string = `james said "hello"`
```

Go, you need to be precise. You need to specify what you want to do

Primitive data types (basic or built in data types):

- integer
- string
- boolean
- float

Composite data types (can hold more than one data type / aggregate data type):

- struct
- slice
- array

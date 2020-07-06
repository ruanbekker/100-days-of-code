## Control Flow

Theres no while loops in Go

Init; Condition; Post

GoByExample:
- https://gobyexample.com/for

## Basic For Loop

Basic for loop

```
package main

import "fmt"

func main(){
    for x := 0; x <= 100; x++ {
        fmt.Println(x)
    }
}
```

Nested for loop / For statement with for clause:

```
package main

import "fmt"

func main() {
	for x := 0; x <= 100; x++ {
		for y := 0; y <= 3; y++ {
			fmt.Printf("X is %d\n", y)
		}
		fmt.Printf("Y is %d\n", x)
	}
}
```

For statement with single condition

```
package main

import (
	"fmt"
)

func main() {
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("done")
}

```

For loop with if statement:

```
package main

import (
	"fmt"
)

func main() {
	x := 0
	for {
		if x > 10 {
			break
		}
		fmt.Println(x)
		x++
	}
	fmt.Println("done")
}

```

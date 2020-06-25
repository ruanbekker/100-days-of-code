## Day 8

Numeric Types:
- https://golang.org/ref/spec#Numeric_types

### Integers

int => 12
float64 => 1.12

```
package main

import (
	"fmt"
)

var x int
var y float64

func main() {
	x = 10
	y = 1.2
	fmt.Println(x)
	fmt.Println(y)
}
```

We cannot store integer values in variables that was declared as another type such as float64

`uint8` and `int8`

int8 types can range from -127 to 127 and uint8 can range from 0 to 255

```
package main

import (
	"fmt"
)

var x int8
var y uint8

func main() {
	x = -127
	y = 255
	fmt.Println(x)
	fmt.Println(y)
}
```

Reference:
- https://godoc.org/builtin#uint

```
package main

import (
	"fmt"
)

var x uint8

func main() {
	x = 0
  x = 255
	fmt.Println(x)
  // unit8 types can only go from 0 to 255
}
```

### Runtimes

```
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}

// linux
// amd64
```

### Strings

```
package main

import (
	"fmt"
)

func main() {
	x := "hello, world!"
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
```

Raw string literals:

```
package main

import (
	"fmt"
)

func main() {
	x := `"hello, world!"`
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
```

Or:

```
package main

import (
	"fmt"
)

func main() {
	x := `"hello, 
  world!"`
	fmt.Println(x)
	fmt.Printf("%T\n", x)
  //   hello,
  // world!
}
```

Slice of bytes 

```
package main

import (
	"fmt"
)

func main() {
	x := "hello"
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	bs := []byte(x)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
}
```

A byte is an alias of a int8 (which is a type)


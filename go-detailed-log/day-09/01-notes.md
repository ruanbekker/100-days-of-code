## Day 09 Notes

### Numeral Systems

- Decimal
- Binary
- Hexadecimal
- base64

Using Go to print decimal, binary, hexadecimal:

```
package main

import (
	"fmt"
)

func main() {
	s := "H"
	fmt.Println(s)
	
	bs := []byte(s)
	fmt.Println(bs)
	
	n := bs[0]
	fmt.Println(n)
	fmt.Printf("%T\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%#x\n", n)
}
```

### Constants

Declare constants and print their values:

```
package main

import (
	"fmt"
)

const a = 1
const b = 1.2
const c = "hello"

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
```

We can also declare multiple constants using the parenteses way:

```
package main

import (
	"fmt"
)

const (
  // untyped constants (constant of kind - gives the compiler flexability)
	a = 1
	b = 1.1
	c = "hello"
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
```

The same, but typed constants (constants of a kind):

```
package main

import (
	"fmt"
)

const (
	// typed constant
	a int = 1
	b float64 = 1.1
	c string = "hello"
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
```

### Iota

Increment numbers with Go using constants:

```
package main

import (
	"fmt"
)

const (
	// untyped constant
	a = iota
	b = iota
	c 
	d
)

const (
	e = iota
	f 
)

func main() {
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}

// 0
// int
// 1
// 2
// 3
// 0
// 1
```

To start incrementing from a specific number:

```
package main

import (
	"fmt"
)

const (
	a = 10 + iota
	b
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
}

// 10
// 11
```

### Bit Shifting

```
package main

import (
	"fmt"
)


func main() {
	// best practice: limit the scope, if you dont need package level scope, dont do it.
	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)
	
	// bit shifting
	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)
	
	// response (decimals, binary):
	// 4		100
	// 8		1000
	
}
```

Printing kb, mb, gb:

```
package main

import (
	"fmt"
)

func main() {
	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb
	
	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}
// 1024       10000000000
// 1048576		100000000000000000000
// 1073741824	1000000000000000000000000000000
```

Using iota to return the same results:

```
package main

import (
	"fmt"
)

const (
	// throw away value
	_ = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {	
	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)
}
// 1024       10000000000
// 1048576		100000000000000000000
// 1073741824	1000000000000000000000000000000
```

### Excercises

Define a number, print the decimal, binary and hex number:

```
package main

import (
	"fmt"
)

func main() {
	x := 10
	fmt.Printf("%d, %b, %#x\n", x, x, x)
}
```

Use the following operators, and write expressions and assign their values to varibales


```
package main

import (
	"fmt"
)

func main() {
	a := ( 10 == 42)
	b := ( 10 <= 42)
	c := ( 10 >= 42)
	d := ( 10 != 42)
	e := ( 10 > 42)
	f := ( 10 < 42)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)	
}
```



# go commands

## go fmt

The go fmt command formats your code, which allows a consistent standard of formatting

The go fmt command and fmt package are two different things

### Pronounciation

go fmt is pronounced as "go fumt"

### Example formatting

Create a file with strange formatted style:

```
$ cat app.go
package main

import "fmt"

func main(){
    fmt.Println("hello")
    }
```

After we run `go fmt` on our code, we will notice its formatted:

```
$ go fmt app.go
app.go
```

Verify that the code was formatted:

```
$ cat app.go
package main

import "fmt"

func main() {
	fmt.Println("hello")
}
```

To run this with a bunch of files, in the current working directory:

```
$ go fmt ./..
```

## go run

go run, runs compiles your code and runs it (leaves no binary behind)

## go build

go build, builds a binary to the current working directory

## go install

go install, installs a binary to your `${GOPATH}/bin` directory. Note that `GOBIN` environment variable needs to be set to `${GOPATH}/bin`

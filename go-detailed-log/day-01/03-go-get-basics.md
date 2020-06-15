# Go Get Basics

The command `go get` retrieves the go code from the specified repository and downloads it into your workspace

## Example

Our go workspace is empty at the moment:

```
$ tree ~/go
/Users/ruan/go
├── bin
├── pkg
└── src
```

If we use `go get -d <remote/user/name>` it will get the source code and stop after downloading, meaning it will not install the binary:

```
$ go get -d github.com/ruanbekker/go-hello
``` 

We can see that the binary was not built:

```
$ tree ~/go
/Users/ruan/go
├── bin
├── pkg
└── src
    └── github.com
        └── ruanbekker
            └── go-hello
                └── hello.go
```

If we have to run it again, but removing `-d`, we will get the binary installed to `${GOPATH}/bin`:

```
$ go get github.com/ruanbekker/go-hello
$ tree ~/go
/Users/ruan/go
├── bin
│   └── go-hello
├── pkg
└── src
    └── github.com
        └── ruanbekker
            └── go-hello
                └── hello.go
```

And since we have `${GOPATH}/bin` in our PATH, we will be able to run the binary:

```
$ go-hello
Hello, World!
```

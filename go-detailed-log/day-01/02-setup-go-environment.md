# Go Workspaces

The following namespace applies to a go workspace:

```
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

- bin: all binaries
- pkg: archives
- src: source code

The `GOPATH` environment variable defines your workspace and the `GOROOT` environment variable defines the path where go is installed.

## Setup Go Environment on Mac

To install golang on mac:

```
$ brew update
$ brew install golang
```

Verify:

```
$ go version
go version go1.14.3 darwin/amd64
```

Create your go workspace directories:

```
mkdir ~/go/{pkg,src,bin}
```

Depending if you have zsh or bash, in my case zsh, append the environment variables in `~/.zshrc`:

```
export GOPATH=~/go
export GOROOT=/usr/local/opt/go/libexec
export PATH=${PATH}:${GOPATH}/bin:${GOROOT}/bin
```

Source your environment:

```
$ source ~/.zshrc
```

## Setup Go Environment on Linux

Get the latest version from [their downloads page](https://golang.org/dl/) and download:

```
$ wget https://dl.google.com/go/go1.14.3.linux-amd64.tar.gz
```

Extract:

```
$ sudo tar -xf go1.14.3.linux-amd64.tar.gz -C /usr/local
```

Create the go workspace directories:

```
$ mkdir -p ~/go/{bin,src,pkg}
```

Append the environment variables in `~/.bashrc`:

```
export GOROOT=/usr/local/go
export GOPATH=~/go
export PATH=${PATH}:${GOROOT}/bin:${GOPATH}/bin
```

Source the environment:

```
$ source ~/.bashrc
```

Verify:

```
$ go version
go version go1.14.3 linux/amd64
```

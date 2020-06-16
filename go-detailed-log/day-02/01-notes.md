## Day 2 - Notes

### Direct vs Indirect Dependencies

Direct Dependencies:

- Is the software retrieved from the package

Indirect Dependencies:

- The software retrieved is dependent on a dependency

### Go dependencies

From their [blog](https://blog.golang.org/using-go-modules), working with dependencies, upgrading/downgrading:

- `go get rsc.io/sampler`
- `go test`
- `go list -m all`
- `go list -m rsc.io/sampler
- `go list -m -versions rsc.io/sampler`
- `go get rsc.io/sampler@v1.3.1`

Reference:

- https://blog.golang.org/using-go-modules


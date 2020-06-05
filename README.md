## Embedding in Go
```go
package embedding

import "fmt"

type Base struct {
	 name string
}

type Entity struct {
	Base
}

func (b *Base) play() {
	fmt.Println("base")
}

func Test() {
	e := &Entity{}
	// https://hackthology.com/object-oriented-inheritance-in-go.html
	e.play() // will call Base's play without downwards-shadowing.
}
```

## One Sentence
- **Common**: `Interface First`
- **Strategy Pattern**: `Behaviour Class`
- **Observer Pattern**: `A Push B,then A should compose B`
- **Decorator Pattern**: `Nested doll`

## References
- [Design Patterns in Golang](https://golangbyexample.com/all-design-patterns-golang/)
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
- [Refactoring Guru](https://refactoringguru.cn/design-patterns)
- [Graphic Design Patterns](https://design-patterns.readthedocs.io/)
- [Design Patterns in Object Oriented Programming](https://www.youtube.com/playlist?list=PLrhzvIcii6GNjpARdnO4ueTUAVR9eMBpc)
- [设计模式之美](https://time.geekbang.org/column/intro/250)
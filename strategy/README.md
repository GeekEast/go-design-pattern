## Principle
- **开闭原则**: 用class来封装behaviour, **快速定义新的行为类而不影响已经存在的行为类**, 大大提高程序的**拓展性**。
- **依赖反转**: **面向接口**将方法的`调用`和方法的`实现`进行了解耦、这样做也可以实现**业务逻辑**和**算法逻辑**的分离。增强**可读性**和**拓展性**。

## Definition of Strategy
- You need to define
  - strategy `interface`
  - **behavirour** entity
```go
type Strategy interface {
	algorithm()
}

// strategy 1
type algorithm1Entity struct{}

func (e *algorithm1Entity) algorithm() {
	fmt.Println("algorithm 1")
}

// strategy 2
type algorithm2Entity struct{}

func (e *algorithm2Entity) algorithm() {
	fmt.Println("algorithm 2")
}
```
## Creation of Strategy `Factory`
- stateless
```go
var Store = map[string]Strategy{
	"behaviour1": &algorithm1Entity{},
	"behaviour2": &algorithm2Entity{},
}
```
- stateful
```go
func getStrategy(t string) (s Strategy, err error) {
	if t == "behaviour1" {
		return &algorithm1Entity{}, nil
	} else if t == "behaviour2" {
		return &algorithm2Entity{}, nil
	} else {
		return nil, errors.New("no such strategy")
	}
}
```
## Selection of Strategy
- stateless
```go
func use(t string) {
    Store[t].algorithm()
}
```
- stateful
```go
func use(t string) {
	behaviourEntity, err := getStrategy(t)
	if err != nil {
		fmt.Println(err)
	}
    behaviourEntity.algorithm()
}
```
## Diagram
<p align="center"><img style="display: block; width: 600px; margin: 0 auto;" src=img/2020-06-04-00-22-57.png alt="no image found"></p>

## Advantage
- **业务逻辑**和**算法逻辑**的清晰分层，`高层模块的实现不依赖于底层模块的实现，而依赖于底层模块实现的抽象`。能够在不改变高层调用的情况下，动态选择实现方式。
- 解决**复杂继承模式**下**方法复用**的问题。将方法都存在**策略库**中，通过组合的方式，在运行时灵活调用即可，也提高了代码的**复用程度**。
- **增强了程序的拓展性**。 通过**行为类**来封装具体实现方式，让这些行为类之间互不干扰；通过依赖接口而不是依赖具体实现, 客户端无需更改调用方式；因此，后台可以无限地增加具体实现方式！

## Insights
- 好的分层需要**依赖反转**来解耦**层与层**之间的依赖
- **继承**默认给了子类父类的全部，复用的粒度是整个父类
- **组合**通过成员的方式，可以自定义复用的粒度，更加灵活

## Guess
- 在组合模式下，**实体性质的**struct成员变量包含**属性**和**行为**，而行为则是以接口的形式存在，新建struct的时候可以传入行为性质的struct帮助实体性质的struct实现其方法。
- 在组合模式下, 将func绑定在某个struct上面仅仅适用于行为性质的struct, 这类struct存在的核心价值在于其方法，成员属性作为状态而服务于struct的方法。

## Reference
- [策略模式1](https://design-patterns.readthedocs.io/zh_CN/latest/behavioral_patterns/strategy.html)
- [策略模式2](https://refactoringguru.cn/design-patterns/strategy)
- [依赖反转](https://zhuanlan.zhihu.com/p/24175489)
  
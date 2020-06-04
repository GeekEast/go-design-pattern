## Context
- **问题**: 在`多层级`的继承背景下，如果各级子类的行为存在一定的共性，但是并不严格遵循垂直方向的共享关系，如果硬要使用继承来实现行为方式的code-reuse,会提高类间关系的复杂度以及代码的维护成本。
- **前提**: 类间的区别仅仅体现在**行为**上。
- **解决**: 抛弃多层级的继承关系，使用策略模式，仅需要一个类，在生成对象的时候动态选择行为方式。
- **本质**: 
  - **开闭原则**: 用class来封装behaviour, **快速定义新的行为类而不影响已经存在的行为类**
  - **依赖反转**: **面向接口**将方法的调用和方法的实现进行了解耦、这样做也可以实现业务逻辑和算法逻辑的分离。

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
- 使用**has-a**取代**is-a**，实现`更灵活`和`更细粒度`的**代码复用**。

## Insights
- 好的分层需要**依赖反转**来解耦**层与层**之间的依赖
- **继承**默认给了子类父类的全部，复用的粒度是整个父类
- **组合**通过成员的方式，可以自定义复用的粒度，更加灵活
- 策略模式

## Guess
- 在组合模式下，**实体性质的**struct成员变量包含**属性**和**行为**，而行为则是以接口的形式存在，新建struct的时候可以传入行为性质的struct帮助实体性质的struct实现其方法。
- 在组合模式下, 将func绑定在某个struct上面仅仅适用于行为性质的struct, 这类struct存在的核心价值在于其方法，成员属性作为状态而服务于struct的方法。

## Reference
- [策略模式1](https://design-patterns.readthedocs.io/zh_CN/latest/behavioral_patterns/strategy.html)
- [策略模式2](https://refactoringguru.cn/design-patterns/strategy)
- [依赖反转](https://zhuanlan.zhihu.com/p/24175489)
  
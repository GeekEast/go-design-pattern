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

## Principle
- **开闭原则**: 通过行为类来封装具体实现，能够在不修改现存行为类的情况下添加新的行为类，提高了水平层面的可拓展性。
- **依赖反转**: 通过interface来对方法进行调用，使得高层模块不依赖于低层模块的实现，而依赖于低层模块实现的抽象。提高了垂直方向上的可拓展性。
- **组合大于继承**: 继承的复用粒度是整个父类，组合将复用粒度自定义化，而且不限类型(属性或者行为), 提高了代码复用的灵活性，有利于在创建新类的时候最大程度的利用已有代码。(前提: code split做的好,这又依赖于单一原则)


## Guess
- struct分为两种,一种是实体性质的，例如对应数据库的一张表；另外是行为性质的，成员中的属性完全是为了维护行为的状态而存在，而改struct最好是实现了某种接口.
- 实体类的具体行为可以通过compose interface和delegation来实现。

## Summary
- 分层思维 -> 解耦以易于拓展
- 组合而非继承 -> 降低复用粒度, 解除复杂关系
- 类的封装 -> 开闭原则, 在拓展的时候没有副作用
- `is-a`是狭义的`has-a`:
  - Employ has sth belonging to Person (has-a)
  - Employ has all things belonging to Person (is-a)
  - What if Employ has sth belongs to Bird?
    - No mutilple inheritance
    - Now we can use composition

## Reference
- [策略模式1](https://design-patterns.readthedocs.io/zh_CN/latest/behavioral_patterns/strategy.html)
- [策略模式2](https://refactoringguru.cn/design-patterns/strategy)
- [依赖反转](https://zhuanlan.zhihu.com/p/24175489)
  
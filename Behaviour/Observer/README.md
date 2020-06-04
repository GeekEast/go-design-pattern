## Diagram
<p align="center"><img style="display: block; width: 600px; margin: 0 auto;" src=img/2020-06-04-20-00-54.png alt="no image found"></p>

## Summary
- **观察者模式**解决的是**对象**之间**1对多**的依赖关系
- **依赖接口**是提高**可拓展性**的关键
- 实现接口的可以使某个实体类，也可以是行为类
  - 实体类: 这样使得数据与行为实现了耦合, 不推荐
  - **行为类**: 实体类可以`组合`行为类，通过`委托`实现功能, **推荐**


## References
- [Observer Pattern in Go](https://golangbyexample.com/observer-design-pattern-golang/)
- [Association, Aggregation, Composition](http://www.srcmini.com/37743.html)
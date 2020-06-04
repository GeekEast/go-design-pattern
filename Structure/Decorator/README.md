## Motivation
一般有两种方式可以实现给一个类或对象增加功能：

- `静态`: **继承机制**，使用继承机制是给现有类添加功能的一种有效途径，通过继承一个现有类可以使得子类在拥有自身方法的同时还拥有父类的方法。但是这种方法是静态的，100个子类需要增加同一个功能，则需要派生出100个该功能的子类，更不要说想添加两个功能。
- `动态`: **关联机制**，即将一个类的对象嵌入另一个对象中，由另一个对象来决定是否调用嵌入对象的行为以便扩展自己的行为，我们称这个嵌入的对象为装饰器(Decorator)
  
装饰模式以对客户透明的方式动态地给一个对象附加上更多的责任，换言之，客户端并不会觉得对象在装饰前和装饰后有什么不同。**装饰模式可以在不需要创造更多子类的情况下，将对象的功能加以扩展**。这就是装饰模式的模式动机。

## Example
```java

InputStream in = new FileInputStream("/user/wangzheng/test.txt");
InputStream bin = new BufferedInputStream(in);
DataInputStream din = new DataInputStream(bin);
int data = din.readInt();
```

## References
- [Decorator Pattern](https://www.youtube.com/watch?v=GCraGHx6gso&list=PLrhzvIcii6GNjpARdnO4ueTUAVR9eMBpc)
- [Decorator](https://refactoringguru.cn/design-patterns/decorator)
- [装饰器模式](https://time.geekbang.org/column/article/204845)
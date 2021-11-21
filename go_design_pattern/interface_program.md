## 设计模式的学习与总结
如果说 数据结构与算法教你如何写出高效的代码，那么设计模式就是告诉你如何写出可扩展、可读、可维护的高质量代码，它们跟平时的编码会有直接的关系，也会直接影响你的开发能力。
### 捋清面向对象、设计原则、设计模式、编程规范、重构五者的关系
1. 面向对象<br>
   面向对象具有四大特性 封装、继承、多态、抽象<br>
2. 设计原则<br>
   五大设计原则 SOLID 分别是 SRP(单一职责原则) OCP(开闭原则) LSP(里氏替换原则) ISP(接口隔离原则) DIP(依赖倒置原则)<br>
   另外还有DRY(Don't repeated yourself)，KISS(Keep is simple and stupid)，YAGNI(You aren't gonna need it)，LOD(Law of Demeter)
3. 设计模式<br>
   涵盖创建型、结构型、行为型的设计模式<br>
4. 编程规范<br>
   编程规范主要解决的是代码的可读性问题，相对于设计原则、设计模式更加具体并注重代码细节。例如变量、类、函数命名，代码注释规范、函数长短、参数不能过多等等。<br>
5. 代码重构<br>
   目的、对象、时机和方法都需要进行细量评估，保证重构后运行不出错，代码可测性，以及应用设计模式、设计原则进行不同规模的重构。重构也是保证代码质量不下降，避免或者补偿破窗效应的一个有效手段。<br>
   大重构包含多个规模以及高层次，小重构小规模低层次。相比于大重构，小重构可以持续进行，对函数命名，注释规范，抽象升级都属于小重构范围。<br>

### 讲讲设计原则
设计原则对于代码设计具有指导性意义。

#### SRP(single responsibility principle) 单一职责原则
A class or module should have a single responsibility。(一个类或者模块只负责完成一个职责(或功能))。<br>
1. 如何理解？
其实单一职责很好理解，就是不要设计一个大而全的类，一个不知道怎么命名才合适的类。可以理解为，当设计一个类的时候，尽量保持这个类足够单纯，足够具有可被上层组合使用，粒度更加适宜。举个简单的例子，一个领域既包含用户操作的功能，又包括订单的操作。 用户操作和订单应该属于两个独立的业务领域模型，更应该被设计成单一的用户类和订单类。<br>
2. 如何判断？
真实的开发场景很难拿捏如何设计一个单一的职责的一个类。举个实际的业务例子，比如客户报备这个场景。
```golang

type Customer struct{
    Name string
    Tel string
    Proj string
    Belong string
    ExtraInfo string
    Age int
    Address string

    Recommend func(context)error
}

```
针对这个场景，上面的类的属性和方法的确归属于客户，操作也属于客户的操作，其实也满足单一职责的原则。但是实际上，由于业务变化，客户归属是一个经常变化的阶段，`Belong`这个字段跟客户其实是一种很强的绑定关系，但是又不能单纯的属于客户操作，而是通过其他操作完成客户阶段变化，归属也会变化，那么这个时候其实我们应该给客户独立出一个类来承载这个变化。<br>
```golang
type CustomerBelong struct{
    BelongNow string
    BelongBefore string
    CustomerID string
    ...
}
```
其实在业务发展初期可以设计一个粗粒度的类，随着业务的深度发展，可以拆分成更加细致的类。 那么在重构阶段，应该根据当前业务，设计一个更加合适，细粒度的类来完成适合当前业务变化类，来承载业务的变化。<br>
3. 职责越单一越好？
没有银弹的设计，只有适合当前变化的设计。 如果一个类拆分的过于小，维护性就会成问题。因为过小的类，在这个代码中越多，整个代码就会变得很松散，没有了高内聚的味道了。

#### OCP(Open Closed Principle) 开闭原则
software entities (modules, classes, functions, etc.) should be open for extension , but closed for modification. 添加一个新的功能应该是，在已有代码基础上扩展代码（新增模块、类、方法等），而非修改已有代码（修改模块、类、方法等。<br>

1. 如何理解？
开闭原则可以理解为代码的扩展性是否得到满足，是判断一段代码是否易扩展准则。也就是在应对未来的需求的变化的的时候，是否能够做到对扩展开放，对修改关闭。提到扩展，顺便提一下，在设计代码的时候，应该具有偏向顶层的抽象思维，时刻保持着扩展意识、抽象意识、封装意识。<br>

2. 如何设计？
举个具体的例子，假设我们需要设计一个Kafka发送异步消息。那么我们在设计接口的时候，需要将其设计成与kafka无关的异步消息接口，使其通过依赖注入的方式来调用，方便后续如果需就能替换成rocketMQ发异步消息。<br>
```golang

type MsgQueue interface{
   SendNotify(string)error
   ...
}

type KafkaQueue struct{
}

func (k *KafkaQueue) SendNotify(string)error{
   // do notify
   return nil
}

type RockMQ struct{}
func (r *RockMQ) SendNotify(string)error{
   // do notify
   return nil
}

// 设计异步发送消息
func AsyncSend(m MsgQueue)error{
   m.SendNotify("hello,trans")
   return nil
}

func main(){
   AsyncSend(&KafkaQueue{})
}
```

3. 如何灵活使用？
其实上述举的例子只是开闭原则的冰山一角，开闭原则主要的一个理念是留好扩展点，适应业务变化，毕竟唯一不变的只有变化本身。开闭原则也并不是完全避免对代码的修改，其实可以对代码适当的修改，只要不会引起测试用例的变化，或者调用方的变化，那么这种修改就是适当的，它本身也是一种扩展。提高代码扩展性的方式有很多，例如多态、依赖注入、基于接口而非实现编程还有大部分的设计模式，装饰，策略，模版，职责链，状态等。

#### LSP(Liskov Substitution Principle) 里氏替换原则
Functions that use pointers of references to base classes must be able to use objects of derived classes without knowing it.
子类对象能够替换程序中父类对象出现的任何地方，并且保证原来程序的逻辑行为不变及正确性不被破坏。<br>

1. 如何理解？
Design By contract。 函数行为应该遵循父类定义的行为协议，对输入输出异常都应该与父类保持一致，使用父类单元测试的用例应该对子类是保持有效且正常的。<br>
由于go语言没有继承特性，使用组合的例子来说明里氏替换原则
```golang
// 假设我们需要设计考勤的类
// 考虑到考勤打卡功能
type Attendance interface {
	Pinch(emplyee string, t time.Time) bool
}

type AttendanceImpl struct {
	pinchTime time.Time
}
// 父类方法
func (ai *AttendanceImpl) Pinch(emplyee string, t time.Time) bool {
	// 简单的规则
	if emplyee != "" && t.Before(time.Now()) {
		return true
	}
	return false
}

// 现假设考勤打卡需要增加时间判断，超过特定的时间就不能打卡
// 重写方法
func (ai *AttendanceImpl) Pinch(emplyee string, t time.Time) bool {
	// 简单的规则
	if emplyee != "" && ai.pinchTime.After(t) {
		return true
	}
	return false
}

func attendancePinch(a Attendance) bool {
	return a.Pinch("vito", time.Now())
}

func Test_LSP(t *testing.T) {
	attendancePinch(&AttendanceImpl{})
}
```

#### ISP(Interface Segregation Principle) 接口隔离原则
Clients should not be forced to depend upon interfaces that they do not use。 调用者/使用者不应该被强迫依赖它不需要的接口。<br>

1. 如何理解？

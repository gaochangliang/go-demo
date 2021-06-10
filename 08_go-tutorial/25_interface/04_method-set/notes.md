## english
```
"Polymorphism is the ability to write code that can take on different behavior through the
 implementation of types. Once a type implements an interface, an entire world of
 functionality can be opened up to values of that type."
 - Bill Kennedy

"Interfaces are types that just declare behavior. This behavior is never implemented by the
 interface type directly, but instead by user-defined types via methods. When a
 user-defined type implements the set of methods declared by an interface type, values of
 the user-defined type can be assigned to values of the interface type. This assignment
 stores the value of the user-defined type into the interface value.

 If a method call is made against an interface value, the equivalent method for the
 stored user-defined value is executed. Since any user-defined type can implement any
 interface, method calls against an interface value are polymorphic in nature. The
 user-defined type in this relationship is often called a concrete type, since interface values
 have no concrete behavior without the implementation of the stored user-defined value."
  - Bill Kennedy

Receivers       Values
-----------------------------------------------
(t T)           T and *T
(t *T)          *T

Values          Receivers
-----------------------------------------------
T               (t T)
*T              (t T) and (t *T)


SOURCE:
Go In Action
William Kennedy
/////////////////////////////////////////////////////////////////////////

Interface types express generalizations or abstractions about the behaviors of other types.
By generalizing, interfaces let us write functions that are more flexible and adaptable
because they are not tied to the details of one particular implementation.

Many object-oriented lagnuages have some notion of interfaces, but what makes Go's interfaces
so distinctive is that they are SATISIFIED IMPLICITLY. In other words, there's no need to declare
all the interfaces that a given CONCRETE TYPE satisifies; simply possessing the necessary methods
is enough. This design lets you create new interfaces that are satisifed by existing CONCRETE TYPES
without changing the existing types, which is particularly useful for types defined in packages that
you don't control.

All the types we've looked at so far have been CONCRETE TYPES. A CONCRETE TYPE specifies the exact
representation of its values and exposes the intrinsic operations of that representation, such as
arithmetic for numbers, or indexing, append, and range for slices. A CONCRETE TYPE may also provide
additional behaviors through its methods. When you have a value of a CONCRETE TYPE, you know exactly
what is IS and what you can DO with it.

There is another kind of type in Go called an INTERFACE TYPE. An interface is an ABSTRACT TYPE. It doesn't
expose the representation or internal structure of its values, or the set of basic operations they support;
it reveals only some of their methods. When you have a value of an interface type, you know nothing about
what it IS; you know only what it can DO, or more precisely, what BEHAVIORS ARE PROVIDED BY ITS METHODS.

-------------------

type ReadWriter interface {
    Reader
    Writer
}

This is called EMBEDDING an interface.


-------------------

A type SATISFIES an interface if it possesses all the methods the interface requires.

-------------------

Conceptually, a value of an interface type, or INTERFACE VALUE, has two components,
    a CONCRETE TYPE and a
    VALUE OF THAT TYPE.
These are called the interface's
    DYNAMIC TYPE and
    DYNAMIC VALUE.

For a statically typed language like Go, types are a compile-time concept, so a type is not a value.
In our conceptual model, a set of values called TYPE DESCRIPTORS provide information about each type,
such as its name and methods. In an interface value, the type component is represented by the appropriate
type descriptor.


var w io.Writer
w = os.Stdout
w = new(bytes.Buffer)
w = nil


var w io.Writer
w
type: nil
value: nil

w = os.Stdout
w
type: *os.File
value: the address where a value of type os.File is stored

w = new(bytes.Buffer)
w
type: *bytes.Buffer
value: the address where a value of type bytes.Buffer is stored

w = nil
w
type: nil
value: nil

-------------------
The Go Programming Language
Donovan and Kernighan

Caplitalization and identation mine.
```

## chinese
```
"多态性是指编写代码的能力，它可以通过类型的实现而采取不同的行为。
 实现类型的能力。一旦一个类型实现了一个接口，整个世界的
 功能就可以向该类型的值开放。"
 - 比尔-肯尼迪

"接口是只是声明行为的类型。这种行为从来不是由
 直接实现，而是由用户定义的类型通过方法实现。当一个
 用户定义的类型实现了接口类型所声明的一系列方法时，用户定义的类型的值就可以分配给接口类型的值。
 用户定义的类型的值可以被分配给接口类型的值。这种赋值
 将用户定义的类型的值存储到接口的值中。

 如果对一个接口值进行方法调用，那么存储的用户定义值的等效方法将被执行。
 存储的用户定义的值被执行。由于任何用户定义的类型可以实现任何
 接口，所以针对接口值的方法调用在本质上是多态的。这种关系中的
 在这种关系中，用户定义的类型通常被称为具体类型，因为接口值
 如果没有存储的用户定义值的实现，就没有具体的行为。
  - 比尔-肯尼迪

接收者    值类型
-----------------------------------------------
(t T)     T和*T
(t *T)    *T

价值       接收者
-----------------------------------------------
T (t T)
*T (t T) 和 (t *T)


源于此。
Go In Action
威廉-肯尼迪
/////////////////////////////////////////////////////////////////////////

接口类型表达了对其他类型的行为的概括或抽象。
通过概括，接口让我们编写的函数更加灵活，适应性更强
因为它们不会被某个特定的实现的细节所束缚。

许多面向对象的语言都有一些接口的概念，但是Go的接口
的独特之处在于它们是以非强制性的方式实现的。换句话说，我们没有必要去声明
的所有接口；只要拥有必要的方法
就足够了。这种设计让你可以在不改变现有CONCRETE TYPE的情况下，创建新的、被现有CONCRETE TYPE所满足的接口。
而不需要改变现有的类型，这对于定义在你无法控制的包中的类型来说特别有用。
你不控制的包中定义的类型特别有用。

到目前为止，我们所看到的所有类型都是 CONCRETE TYPES。一个 CONCRETE 类型指定了其值的确切
它的值的精确表示，并暴露了该表示的内在操作，例如
数字的算术，或索引，附加，和切片的范围。CONCRETE类型也可以通过它的方法提供
通过它的方法提供额外的行为。当你有一个CONCRETE类型的值时，你会清楚地知道
是什么以及你能用它做什么。

Go中还有一种类型叫做界面类型。一个接口是一个抽象类型。它并不
暴露其值的表示或内部结构，或其支持的基本操作集。
它只揭示了它们的一些方法。当你有一个接口类型的值时，你对它是什么一无所知。
它是什么；你只知道它能做什么，或者更准确地说，它的方法提供了什么行为。

-------------------

Type ReadWriter 接口 {
    读者
    写入者
}

这被称为EMBEDDING接口。


-------------------

如果一个类型拥有一个接口所要求的所有方法，那么它就能满足这个接口。

-------------------

从概念上讲，一个接口类型的值，或者说接口值，有两个组成部分。
    一个CONCRETE类型和一个
    该类型的值。
这些被称为接口的
    动态类型和
    DYNAMIC VALUE。

对于像Go这样的静态类型语言，类型是一个编译时的概念，所以类型不是一个值。
在我们的概念模型中，一组叫做TYPE DESCRIPTORS的值提供了关于每个类型的信息。
比如它的名字和方法。在一个接口值中，类型组件是由适当的
类型描述符来表示。


var w io.Writer
w = os.Stdout
w = new(byte.Buffer)
w = nil


var w io.Writer
w
类型: nil
值: nil

w = os.Stdout
w
类型。*os.File
值：存储os.File类型的值的地址

w = new(byte.Buffer)
w
类型。*bytes.Buffer
值：存储byte.Buffer类型的值的地址

w = nil
w
类型：nil
值：nil

-------------------
Go编程语言
多诺万和克尼根

大写字母和标识是我的。
```
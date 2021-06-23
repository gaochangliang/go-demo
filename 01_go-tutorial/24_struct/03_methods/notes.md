## english 
```
 We used shorthand notation:
 to create a variable named p1 of type person
 to create a variable named p2 of type person
 We initialized those variables with specific values
 We used the short variable declaration operator with a struct literal to initialize
 ----------------------------------------
 here is how we talk about structs:
 -- user defined type
 -- we declare the type
 -- the type has fields
 -- the type can also have "tags"
 ----  we haven't seen this yet
 -- the type has an underlying type
 ---- in this case, the underlying type is struct
 -- we declare variables of the type
 -- we initialize those variables
 ---- initialize with a specific value, or
 ---- or, initiliaze to the zero value
 -- a struct is a composite type
 ----------------------------------------
 Bill Kennedy:
 Go allows us the ability to declare our own types.
 Struct types are declared by composing a fixed set of unique fields together.
 Each field in a struct is declared with a known type.
 This could be a built-in type or another user defined type.
 Once we have a type declared, we can create values from the type
 When we declare variables, the value that the variable represents is always initialized.
 The value can be initialized with a specific value or it can be initialized to its zero value
 For numeric types, the zero value would be 0; for strings it would be empty;
 and for booleans it would be false.
 In the case of a struct, the zero value would apply to all the different fields in the struct.
 Anytime a variable is created and initialized to its zero value, it is idiomatic to use the keyword var.
 Reserve the use of the keyword var as a way to indicate that a variable is being set to its zero value.
 If the variable will be initialized to something other than its zero value,
 then use the short variable declaration operator with a struct literal
```
中文
```
 我们使用速记符号。
 创建一个名为p1的人物类型的变量
 创建一个名为p2的人物类型的变量。
 我们用特定的值初始化了这些变量
 我们使用简短的变量声明操作符和一个结构字来初始化
 ----------------------------------------
 以下是我们谈论结构的方式。
 --用户定义的类型
 --我们声明该类型
 --该类型有字段
 --该类型也可以有 "标签"
 ----，我们还没有看到这个
 --该类型有一个底层类型
 ---- 在这种情况下，底层类型是struct
 -- 我们声明该类型的变量
 -- 我们初始化这些变量
 ----，用一个特定的值初始化，或者
 ---- 或者，初始化为零值
 -- 结构是一种复合类型
 ----------------------------------------
 比尔-肯尼迪。
 Go让我们有能力声明自己的类型。
 结构类型是通过将一组固定的独特字段组合在一起而声明的。
 结构中的每个字段都用一个已知的类型来声明。
 这可能是一个内置类型或另一个用户定义的类型。
 一旦我们声明了一个类型，我们就可以从该类型中创建值
 当我们声明变量时，该变量所代表的值总是被初始化。
 这个值可以被初始化为一个特定的值，也可以被初始化为其零值
 对于数字类型，零值是0；对于字符串，零值是空。
 而对于布尔运算，它将是假的。
 如果是结构体，零值将适用于结构体中的所有不同字段。
 任何时候，当一个变量被创建并初始化为零值时，使用关键字var是一种习惯。
 保留使用关键字var的方式，以表明变量被设置为零值。
 如果变量将被初始化为零值以外的东西。
 那么请使用短变量声明操作符和结构文字
```
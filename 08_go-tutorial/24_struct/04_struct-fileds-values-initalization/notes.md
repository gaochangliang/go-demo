english
```
(p person) is the "receiver"
it is another parameter
not idiomatic to use "this" or "self"


"Not many people know this, but method notation, i.e. v.Method() is actually syntactic sugar and Go also understands the de-sugared version of it: (T).Method(v). You can see an example here. Naming the receiver like any other parameter reflects that it is, in fact, just another parameter quite well.
This also implies that the receiver-argument inside a method may be nil. This is not the case with this in e.g. Java."
SOURCE:
https://www.reddit.com/r/golang/comments/3qoo36/question_why_is_self_or_this_not_considered_a/?utm_source=golangweekly&utm_medium=email
```

中文
```
(p人)是 "接受者"。
它是另一个参数
使用 "this "或 "self "不是成语。


"没有多少人知道，但方法符号，即v.Method()实际上是句法糖，Go也能理解它的去糖版本。(T).Method(v)。你可以在这里看到一个例子。像其他参数一样命名接收方，很好地反映了它实际上只是另一个参数。
这也意味着方法中的接收方参数可能是空的。而在Java中就不是这样了。
源于此。
https://www.reddit.com/r/golang/comments/3qoo36/question_why_is_self_or_this_not_considered_a/?utm_source=golangweekly&utm_medium=email
```
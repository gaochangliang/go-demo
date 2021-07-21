## mapstructure

mapstructure is a Go library for decoding generic map values to structures and vice versa, while providing helpful error handling.


This library is most useful when decoding values from some data stream (JSON, Gob, etc.) where you don't quite know the structure of the underlying data until you read a part of it. You can therefore read a map[string]interface{} and use this library to decode it into the proper underlying native Go structure.


If the received JSON string has only one data, 
it is natural to use the JSON unmarshal deserialization structure.
But if the JSON string has more than one data, 
and it is possible that each data comes from a different structure, 
then you need to use the mapstructure package to map structures, and then determine which structure is based on the fields, and finally JSON.Unmarshal


有时候我们并不知道数据的类型，也许有多个数据类型，
因此我们先转化为map结构，然后根据约定的某个字段再去序列化成结构体

## 参考

https://segmentfault.com/a/1190000023442894
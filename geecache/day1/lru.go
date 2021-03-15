package day1

import (
	"container/list"
)

type Cache struct {
	maxBytes  int64                    //允许最大使用的内存
	nbytes    int64                    //当前已经使用的内存
	ll        *list.List               //标准库实现的双向链表
	cache     map[string]*list.Element //存储的是双向链表的每一个元素
	OnEvicted func(key string, value Value)
}

//表示链表有多少个节点
func (c *Cache) Len() int {
	return c.ll.Len()
}

//element是链表的每一个元素，
//element单节点存储的值是Value接口类型，这里存储的是entry这个结构体
/*
值存储key的好处在于很容易得到key从而删除对应的映射
因为通常来说只是让你删除链表的一个元素，你并不知道key,
只知道要删除的是链表的最后一个元素
*/

type entry struct {
	key   string
	value Value //注意这里的数据类型
}

type Value interface {
	Len() int //用于返回值所占用的内存大小
}

func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

//查找
/*
队头是最近比较常使用的节点
队尾是最近比较少使用的节点
*/
func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		//链表元素移到队头
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

func (c *Cache) RemoveOldest() {
	//返回链表的最后一个元素
	ele := c.ll.Back()
	if ele != nil {
		//删除该元素
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		//根据元素中存储的key删除元素
		delete(c.cache, kv.key)
		//删除键和值占用的空间
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())
		//回调函数不为nil,调用回调函数
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

//新增/修改元素
func (c *Cache) Add(key string, value Value) {
	//存在就更新，不存在就添加元素
	if ele, ok := c.cache[key]; ok {
		//元素移到队头
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		//更新会替换掉旧的值，所以需要差值
		c.nbytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		//添加新的元素到队头
		ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.nbytes += int64(len(key)) + int64(value.Len())
	}
	//如果内存超出最大值，移除最少访问的元素
	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest()
	}
}

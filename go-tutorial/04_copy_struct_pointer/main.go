package _4_copy_struct_pointer

import "sync"

type HandlerFunc func(*Context)

type HandlersChain []HandlerFunc

type Param struct {
	Key   string
	Value string
}

type Params []Param

type Context struct {
	name     string
	Params   Params
	mu       sync.RWMutex
	handlers HandlersChain
	Keys     map[string]interface{}
}

/*
There is a pointer structure that requires a copy of non-pointer data in order to prevent external changes,
but of course we can also return a non-pointer data
*/

/*
map[string]interface{}  Instantiation

obj = map[string]interface{}{
    "name": "gin",
    "password": "1234",
}
*/

func (c *Context) Copy() *Context {
	cp := Context{
		Params: c.Params, //
		name:   c.name,
	}
	cp.handlers = nil
	//Note the way it is written here，He writes similarly to this     struct{}
	cp.Keys = map[string]interface{}{}
	for k, v := range c.Keys {
		cp.Keys[k] = v
	}
	//Because slices are reference types, they need to be copied here
	paramCopy := make([]Param, len(cp.Params))
	copy(paramCopy, cp.Params)
	cp.Params = paramCopy
	return &cp
}

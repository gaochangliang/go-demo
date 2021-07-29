package store

import (
	"fmt"
	"log"
	"time"
)

type Opt struct {
	ConnStr     string
	IdleTimeout time.Duration
}

type CreateStore func(Opt) (IStore, error)

// IStore 这里只是模拟两个操作数据库的方法，具体看实际的结构
type IStore interface {
	SavePersonInfo() error
	GetPersonInfoById(id string)
}

var providerMap = make(map[string]CreateStore)

// Register 注册提供程序
func Register(provider string, fn CreateStore) {
	providerMap[provider] = fn
}

func NewStorer(provider string, opt Opt) IStore {
	fn, has := providerMap[provider]
	if !has {
		log.Fatal(fmt.Errorf("Not support provider %s", provider))
	}
	storer, err := fn(opt)
	if err != nil {
		log.Fatal(err)
	}

	return storer
}

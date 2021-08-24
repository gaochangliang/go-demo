package core

import (
	"fmt"
	"gin/05/global"
	"gin/05/initalize"
	"log"
)

func RunServer() {
	router := initalize.Routers()
	address := fmt.Sprintf(":%d", global.GLOBAL_CONFIG.System.Addr)
	s := initServer(address, router)
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal("server err", err)
		return
		log.
	}
}

type server interface {
	ListenAndServe() error
}

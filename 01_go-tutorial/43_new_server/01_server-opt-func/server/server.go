package server

import "fmt"

//这里创建的时候忽略plugins,主要体现opts参数传递的思路
type Server struct {
	opts    *ServerOptions
	plugins string
}

func NewServer(opts ...ServerOption) *Server {
	s := &Server{
		opts: &ServerOptions{},
	}

	fmt.Println(opts)

	for _, o := range opts {
		o(s.opts)
	}
	return s
}

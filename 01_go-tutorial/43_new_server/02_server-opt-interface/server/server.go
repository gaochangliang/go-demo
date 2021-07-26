package server

//这里创建的时候忽略plugins,主要体现opts参数传递的思路
type Server struct {
	opts    *ServerOptions
	plugins string
}

func NewServer(opts ...Option) *Server {
	s := &Server{
		opts: &ServerOptions{},
	}
	s.WithOptions(opts...)
	return s
}

func (s *Server) WithOptions(opts ...Option) *Server {
	c := s.clone()
	for _, opt := range opts {
		opt.apply(c.opts)
	}
	return c
}

func (s *Server) clone() *Server {
	copyServer := *s
	return &copyServer
}

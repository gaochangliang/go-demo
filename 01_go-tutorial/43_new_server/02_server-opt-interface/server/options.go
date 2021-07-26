package server

type ServerOptions struct {
	address  string // listening address, e.g. :( ip://127.0.0.1:80、 dns://www.google.com)
	network  string // network type, e.g. : tcp、udp
	protocol string // protocol type, e.g. : proto、json
}

// An Option configures a Logger.
type Option interface {
	apply(*ServerOptions)
}

type ServerOption func(options *ServerOptions)

func (s ServerOption) apply(options *ServerOptions) {
	s(options)
}

//之前返回的是函数，这里返回的是接口
func WithAddress(address string) Option {
	return ServerOption(func(options *ServerOptions) {
		options.address = address
	})
}

func WithNetwork(network string) Option {
	return ServerOption(func(options *ServerOptions) {
		options.network = network
	})
}

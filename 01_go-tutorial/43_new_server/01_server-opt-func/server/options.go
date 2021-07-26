package server

type ServerOptions struct {
	address  string // listening address, e.g. :( ip://127.0.0.1:80、 dns://www.google.com)
	network  string // network type, e.g. : tcp、udp
	protocol string // protocol type, e.g. : proto、json
}

type ServerOption func(options *ServerOptions)

func WithAddress(address string) ServerOption {
	return func(options *ServerOptions) {
		options.address = address
	}
}

func WithNetwork(network string) ServerOption {
	return func(options *ServerOptions) {
		options.network = network
	}
}

func WithProtocol(protocol string) ServerOption {
	return func(options *ServerOptions) {
		options.protocol = protocol
	}
}

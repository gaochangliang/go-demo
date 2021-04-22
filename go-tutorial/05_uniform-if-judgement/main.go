package main

type handler func(c *context)

type context struct {
}

func main() {
	path := "/"
	method := "GET"
	var handlers []handler

	if path[0] != '/' {
		panic("path must begin with '/'")
	}

	if method == "" {
		panic("HTTP method can not be empty")
	}

	if len(handlers) == 0 {
		panic("there must be at least one handler")
	}

	//We can rewrite it like this
	assert1(path[0] == '/', "path must begin with '/'")
	assert1(method != "", "HTTP method can not be empty")
	assert1(len(handlers) > 0, "there must be at least one handler")

}

func assert1(guard bool, text string) {
	if !guard {
		panic(text)
	}
}

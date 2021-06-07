package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	msg := "we love kobe"
	rdr := strings.NewReader(msg)
	_, err := io.Copy(os.Stdout, rdr)
	if err != nil {
		fmt.Println("rdr copy error", err)
		return
	}

	rdr2 := bytes.NewBuffer([]byte(msg))
	_, err = io.Copy(os.Stdout, rdr2)
	if err != nil {
		fmt.Println("bytes error", err)
		return
	}

	res, err := http.Get("http://www.mcleods.com")
	if err != nil {
		fmt.Println("http get error", err)
		return
	}

	io.Copy(os.Stdout, res.Body)
	if err := res.Body.Close(); err != nil {
		fmt.Println("res body err", err)
	}

}

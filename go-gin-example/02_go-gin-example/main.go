package main

import (
	"fmt"
	"gin-blog/pkg/setting"
	"gin-blog/routers"
	"net/http"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService https://www.topgoer.com
// @contact.name www.topgoer.com
// @contact.url https://www.topgoer.com
// @contact.email me@razeen.me
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 127.0.0.1:80
// @BasePath /api/v1
func main() {
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    0,
		WriteTimeout:   0,
		MaxHeaderBytes: 0,
	}
	_ = s.ListenAndServe()
}

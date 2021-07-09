package main

import (
	routers "example/01_qrcode/pkg"
	"example/01_qrcode/pkg/setting"
	"fmt"
	"log"
	"net/http"
)

func init() {
	setting.SetUp()
}

func main() {

	routersInit := routers.InitRouter()

	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()

}

package main

import (
	"log"
	"time"
)

func main() {
	nowTime := time.Now()
	log.Printf("nowTime %s  unix %d\n", nowTime, nowTime.Unix())

	expireTime := nowTime.Add(3 * time.Hour)
	log.Printf("expireTime %s unix %d\n", expireTime, expireTime.Unix())
}

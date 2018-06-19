package main

import (
	"fmt"
	"github.com/15125505/zlog/log"
	"time"
)

func main() {
	// log.Info("hello")
	// fmt.Println("hellosss")
	// log.Log.SetLogFile("aa.log")
	info := fmt.Sprintf("%s %s %s", "hello", "world", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(info)
	log.Info(info, "buhaowan")
	aa("hhh", "bbb", "aaa", "dddd", "eeee")
}

func aa(s ...string) {
	log.Info(s)
	fmt.Println("world", s)
}

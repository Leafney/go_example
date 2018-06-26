package main

//https://github.com/15125505/zlog

/*

go get -u github.com/15125505/zlog/log


*/


import (
	"fmt"
	"github.com/15125505/zlog/log"
	"time"
    golog "log"
)

func main() {

    log_for_log()

    //log_for_zlog()
}



func log_for_log() {
    //golog.Println("Hello","World") // 2018/06/26 11:49:25 Hello World

    // 设置
    golog.SetFlags(golog.Ldate|golog.Lshortfile)

    golog.Println("This is","Hahaha") //2018/06/26 logs.go:35: This is Hahaha

    golog.SetFlags(golog.Ldate|golog.Ltime|golog.Llongfile)
    golog.Println("www.baidu.com") //2018/06/26 11:54:20 /Users/xue/Learn/Go/xgo_workspace/src/go_example/logs.go:38: www.baidu.com

    /*
    const (
        Ldate         = 1 << iota     //日期示例： 2009/01/23
        Ltime                         //时间示例: 01:23:23
        Lmicroseconds                 //毫秒示例: 01:23:23.123123.
        Llongfile                     //绝对路径和行号: /a/b/c/d.go:23
        Lshortfile                    //文件和行号: d.go:23.
        LUTC                          //日期时间转为0时区的
        LstdFlags     = Ldate | Ltime //Go提供的标准抬头信息
    )
    */

    //设置日志的前缀
    golog.SetPrefix("【UserLog】")
    golog.SetFlags(golog.Ldate|golog.Ltime|golog.Lshortfile)
    golog.Println("this is a log") //【UserLog】2018/06/26 12:00:40 logs.go:55: this is a log

}


func log_for_zlog() {
    // log.Info("hello")
    // fmt.Println("hellosss")
    log.Log.SetLogFile("aa.log")
    info := fmt.Sprintf("%s %s %s", "hello", "world", time.Now().Format("2006-01-02 15:04:05"))
    fmt.Println(info)
    log.Info(info, "buhaowan")
    aa("hhh", "bbb", "aaa", "dddd", "eeee")

}


func aa(s ...string) {
	log.Info(s)
	fmt.Println("world", s)
}

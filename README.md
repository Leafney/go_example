### golang examples

#### 参考

* [Go 标准库 中文参考](https://wizardforcel.gitbooks.io/golang-stdlib-ref/content/)


#### 常用

##### os.exec
* [golang os.exec包学习 - 简书](https://www.jianshu.com/p/450e9f088a97)
* [golang执行系统命令os exec - 快乐编程](http://www.01happy.com/golang-os-exec/)
* [go-gitlab-webhook/gitlab-webhook.go at master · soupdiver/go-gitlab-webhook](https://github.com/soupdiver/go-gitlab-webhook/blob/master/gitlab-webhook.go)

```
_, err2 := exec.Command("CMD", "/C", "DATE", convertTime.Format("2006-01-02")).Output()
  if err2 != nil {
    log.Error(err2.Error())
  }
```

##### strings
* [golang中strings包用法 - CSDN博客](https://blog.csdn.net/chenbaoke/article/details/40318423)

##### HMAC的SHA1加密 md5 -- [sha1md5.go](sha1md5.go)
* [Golang之HMAC的SHA1加密](http://www.nljb.net/default/Golang%E4%B9%8BHMAC%E7%9A%84SHA1%E5%8A%A0%E5%AF%86/)
* [Golang对字符串进行SHA1哈希运算的代码,Go语言代码片段分享, - 脚本分享网](http://www.sharejs.com/codes/go/7303)

##### log
* [Go语言实战笔记（十八）| Go log 日志 | 飞雪无情的博客](http://www.flysnow.org/2017/05/06/go-in-action-go-log.html)
* [go语言记log：log,logger | Legendtkl](http://legendtkl.com/2016/03/11/go-log/)

##### 交叉编译 -- crosscompile.go

* [golang mac 交叉编译linux， linux无法运行 - SegmentFault 思否](https://segmentfault.com/q/1010000011760900)
* [Golang 在 Mac、Linux、Windows 下如何交叉编译 - CSDN博客](https://blog.csdn.net/panshiqu/article/details/53788067)

##### email

* [Golang：使用自定义模板发送邮件 | Oopsguy](http://oopsguy.com/2017/10/11/sending-emails-with-golang/)
* [golang 通过 qq 开通smtp 发送邮件的 全步骤（tls +ssl） 原 golang 通过 qq 开通smtp 发送邮件的 全步骤（tls +ssl） lwtcumt – 技术成就梦想](http://sparkgis.com/uncategorized/2018/01/golang-%E9%80%9A%E8%BF%87-qq-%E5%BC%80%E9%80%9Asmtp-%E5%8F%91%E9%80%81%E9%82%AE%E4%BB%B6%E7%9A%84-%E5%85%A8%E6%AD%A5%E9%AA%A4%EF%BC%88tls-ssl%EF%BC%89-%E5%8E%9F-golang-%E9%80%9A%E8%BF%87-qq-%E5%BC%80/)
* [smtp - The Go Programming Language](https://golang.org/pkg/net/smtp/#PlainAuth)
* [smtp Golang: 简单smtp邮件发送样例 - 为程序员服务](http://outofmemory.cn/code-snippet/6766/Golang-jiandan-smtp-email-send-yangli)
* [Golang Email - 薛定谔的猫_ - 博客园](http://www.cnblogs.com/juepei/p/4612067.html)

##### gjson

##### 命令行参数的解析

###### flag

flag包实现了命令行参数的解析

```
var (
	configFile string
)

func init() {
	flag.StringVar(&configFile, "c", "./config.json", "config file")
}

func main() {
	flag.Parse()
	processes, err := config.Load(configFile)
	if err != nil {
		panic(err)
	}
```

* [package flag | Go 标准库 中文参考](https://wizardforcel.gitbooks.io/golang-stdlib-ref/content/54.html)
* [Golang flag包使用详解(一) | 黄瓜蘸酱](https://faberliu.github.io/2014/11/12/Golang-flag%E5%8C%85%E4%BD%BF%E7%94%A8%E8%AF%A6%E8%A7%A3-%E4%B8%80/)
* [smtp Golang: 简单smtp邮件发送样例 - 为程序员服务](http://outofmemory.cn/code-snippet/6766/Golang-jiandan-smtp-email-send-yangli) **使用flag接收命令行参数**

###### pflag

* [spf13/pflag: Drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags.](https://github.com/spf13/pflag)
* [Golang之使用Flag和Pflag - 陈健的博客 | ChenJian Blog](https://o-my-chenjian.com/2017/09/20/Using-Flag-And-Pflag-With-Golang/)

###### Cobra

* [flw-cn/go-smart-config: 这是一个“零配置”的配置文件加载模块和命令行参数处理模块](https://github.com/flw-cn/go-smart-config)
* [spf13/cobra: A Commander for modern Go CLI interactions](https://github.com/spf13/cobra)
* [Golang之使用Cobra - 陈健的博客 | ChenJian Blog](https://o-my-chenjian.com/2017/09/20/Using-Cobra-With-Golang/)
* [Cobra - 一个 Golang 命令行项目生成工具 | Jsharkc's secret](https://jsharkc.github.io/2017/07/17/cobra%E5%85%A5%E9%97%A8%E5%B0%8F%E6%95%99%E7%A8%8B/)

###### cli

* [urfave/cli: A simple, fast, and fun package for building command line apps in Go](https://github.com/urfave/cli)
* [用golang来编写cli程序吧，Happy~](https://zhuanlan.zhihu.com/p/30649549)
* [使用Cli构建Go的命令行应用 | jasper的blog](http://www.opscoder.info/cli.html)

##### fmt.Sprintf


##### UPX

1. -ldflags ???


* [UPX: the Ultimate Packer for eXecutables - Homepage](https://upx.github.io/)
* [centos下安装UPX+压缩golang可执行程序 - Go语言中文网 - Golang中文社区](https://studygolang.com/articles/10763)

##### 定时任务

###### cron

* [robfig/cron: a cron library for go](https://github.com/robfig/cron)
* [cron - GoDoc](https://godoc.org/github.com/robfig/cron)
* [golang中执行定时任务的包—cron - CSDN博客](https://blog.csdn.net/skh2015java/article/details/78223951)
* [Go实战--golang中执行cron job定时任务(robfig/cron和jasonlvhit/gocron) - 程序园](http://www.voidcn.com/article/p-ckjvqtxh-bnu.html) **☆**
* [golang crontab的计划任务及定时任务使用 | 峰云就她了](http://xiaorui.cc/2016/03/03/golang-crontab%E7%9A%84%E8%AE%A1%E5%88%92%E4%BB%BB%E5%8A%A1%E5%8F%8A%E5%AE%9A%E6%97%B6%E4%BB%BB%E5%8A%A1%E4%BD%BF%E7%94%A8/)


###### gocron

* [ouqiang/gocron: 定时任务管理系统](https://github.com/ouqiang/gocron)

##### time

```
currentTime = time.Now().Format("2006-01-02 15:04:05")
```

##### service 服务

* [btcsuite/winsvc: Windows service library written in go - Forked from http://code.google.com/p/winsvc/](https://github.com/btcsuite/winsvc)
* [kardianos/service: Run go programs as a service on major platforms.](https://github.com/kardianos/service)
* [我第1个可用的golang小程序 - 老匡的个人空间 - 开源中国](https://my.oschina.net/u/130746/blog/226050) **☆**
* [golang以服务方式运行 - Minho - 开源中国](https://my.oschina.net/idufei/blog/710885) 

##### windows SC命令详解

* [windows SC命令详解-成功不仅是个人荣誉，更是对家人责任-51CTO博客](http://blog.51cto.com/hukunlin/235229)

##### 如何将golang程序编译成windows下具有管理员权限的程序

golang windows程序获取管理员权限（UAC ）

在windows上执行有关系统设置命令的时候需要管理员权限才能操作，比如修改网卡的禁用、启用状态。双击执行是不能正确执行命令的，只有右键以管理员身份运行才能成功。

* [mozey/run-as-admin: Golang example of using a manifest file to prompt "Run as administrator"](https://github.com/mozey/run-as-admin) **☆**
* [How to ask for administer privileges on Windows with Go - Stack Overflow](https://stackoverflow.com/questions/31558066/how-to-ask-for-administer-privileges-on-windows-with-go?utm_medium=organic&utm_source=google_rich_qa&utm_campaign=google_rich_qa)
* [akavel/rsrc: Tool for embedding .ico & manifest resources in Go programs for Windows.](https://github.com/akavel/rsrc)
* [What is a Manifest (in Windows)? (Technical Article)](http://www.samlogic.net/articles/manifest.htm)

`gowins.exe.manifest` :

```
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<assembly xmlns="urn:schemas-microsoft-com:asm.v1" manifestVersion="1.0">
<assemblyIdentity
    version="1.0.0.0"
    processorArchitecture="x86"
    name="gowins.exe"
    type="win32"
/>
<description>Golang example of using a manifest file to prompt "Run as administrator"</description>
<trustInfo xmlns="urn:schemas-microsoft-com:asm.v3">
    <security>
        <requestedPrivileges>
            <requestedExecutionLevel level="requireAdministrator" uiAccess="false"/>
        </requestedPrivileges>
    </security>
</trustInfo>
</assembly>
```

注意：只需要更改 `name="gowins.exe"` 这一项即可，其他的无需更改，无论x86系统还是x64系统均适用，否则无法正常执行。

```
go get github.com/akavel/rsrc

./rsrc -manifest gowins.exe.manifest -o gowins.syso

CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o gowins.exe
```


Todo 用golang写windows服务时需要获取管理员权限才能进行操作

##### 日志 log

* [在Github中最受欢迎的Go日志库集合 - Go开发社区 | CTOLib码库](https://www.ctolib.com/topics-123640.html)
* go-logging -- [op/go-logging: Golang logging library](https://github.com/op/go-logging)
* [issue9/logs: 基于 xml 配置的日志系统](https://github.com/issue9/logs)

###### log

* [Go语言实战笔记（十八）| Go log 日志 | 飞雪无情的博客](http://www.flysnow.org/2017/05/06/go-in-action-go-log.html)
* [Golang 优化之路——自己造一个日志轮子](http://blog.cyeam.com/golang/2017/07/14/go-log)

###### zlog

* zlog -- [15125505/zlog: 一个用于go语言的log模块。](https://github.com/15125505/zlog)

安装：

```
go get github.com/15125505/zlog/log

go get -u github.com/15125505/zlog/log
```

使用:

```
package main

import "github.com/15125505/zlog/log"

func main() {
	log.Log.SetLogFile("logs/filename") // 加入这行代码，就可以写日志文件了
	log.Debug("Some debug info...")
	log.Info("Some infomation...")
	log.Notice("Some notice...")
	log.Error("Error info...")
}
```

##### 后台执行

* [CodyGuo/godaemon: Run golang app as background program, 以后台形式运行golang](https://github.com/codyguo/godaemon)
* [golang 另类方法后台运行程序（linux、windows） - CSDN博客](https://blog.csdn.net/CodyGuo/article/details/53939291)

##### 热更新

* [Golang热加载配置实践 - CSDN博客](https://blog.csdn.net/chenwenhao0304/article/details/49977395)
* [golang配置文件熱更新 - IT閱讀](http://www.itread01.com/articles/1488092299.html)
* 前面两个结合着看
* [John/gf: 模块化、松耦合、轻量级、高性能的Go语言Web开发框架。支持热重启、热更新、多域名、多端口、多服务、HTTP/HTTPS、动态路由等特性，并提供了Web服务开发的系列核心组件，如：MVC、Cookie、Session、服务注册、配置管理、模板引擎、数据校验、分页管理、数据库ORM等等等等，并且提供了数十个实用开发模块集，如：缓存、日志、命令行、二进制、文件锁、数据编码、进程管理、进程通信、并发安全容器、Goroutine池等等等等等等。](https://gitee.com/johng/gf)
* [Golang热更新 · Issue #39 · dongjun111111/blog](https://github.com/dongjun111111/blog/issues/39)
* [facebookgo/grace: Graceful restart & zero downtime deploy for Go servers.](https://github.com/facebookgo/grace)
* [rcrowley/goagain: Zero-downtime restarts in Go](https://github.com/rcrowley/goagain)
* [用Go自己实现配置文件热加载功能 - 云+社区 - 腾讯云](https://cloud.tencent.com/developer/article/1079571) **感觉这个可以**
* [silenceper/gowatch: go程序热编译工具，提升开发效率](https://github.com/silenceper/gowatch)

#### email

##### net/mail


##### gomail

* [Golang 使用gomail包发送邮件 - CSDN博客](https://blog.csdn.net/wj199395/article/details/75206501)
* [go-gomail/gomail: The best way to send emails in Go.](https://github.com/go-gomail/gomail)

***





### golang examples

#### 参考

* [Go 标准库 中文参考](https://wizardforcel.gitbooks.io/golang-stdlib-ref/content/)
* [Golang好用的库](https://jimmysong.io/go-practice/docs/golang_library.html)


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
* [jakecoffman/cron: a cron library for go, updated to have removable jobs](https://github.com/jakecoffman/cron) **优化版本：把任务的结构体里增加了name string字段，这就让所有任务都有了一个唯一标示，借助他的RemoveJob function可以轻松根据唯一标示动态删除计划任务。**

###### robfig/cron扩展

* [golang /robfig/cron库 的学习笔记 – 一个买错域名的站点](http://chuquanl.com/golang-cron%E7%AE%80%E4%BB%8B%E5%8F%8A%E4%BD%BFcron%E6%94%AF%E6%8C%81%E5%B8%A6%E5%8F%82%E6%95%B0%E4%BB%BB%E5%8A%A1%E8%B0%83%E7%94%A8/) **带参数任务的实现，测试后发现不能执行,缺少一行方法的定义**
* [让golang的cron库支持带事件参数的回调 - CSDN博客](https://blog.csdn.net/sryan/article/details/50129133)
* [outman/dcron: Web crontab based on Golang. (Just for study, don't use it on production env.)](https://github.com/outman/dcron)

Todo :动态新增和移除定时任务

###### gocron

* [ouqiang/gocron: 定时任务管理系统](https://github.com/ouqiang/gocron)

###### PPGo_Job

对源 robfig/cron 做了修改

* [george518/PPGo_Job: 定时任务管理-支持多台服务器](https://github.com/george518/PPGo_Job)

###### wechat_pusher

* [hundredlee/wechat_pusher: a wechat message push framework base on golang](https://github.com/hundredlee/wechat_pusher)
* [wechat_pusher : 基于Golang开发的微信消息定时推送框架 - Go中国技术社区 - golang](https://gocn.vip/article/359)

###### cronsun

* [shunfei/cronsun: A Distributed, Fault-Tolerant Cron-Style Job System.](https://github.com/shunfei/cronsun)

###### jobrunner

* [bamzi/jobrunner: Framework for performing work asynchronously, outside of the request flow](https://github.com/bamzi/jobrunner)

##### time

```
currentTime = time.Now().Format("2006-01-02 15:04:05")
```

##### service 服务

* [btcsuite/winsvc: Windows service library written in go - Forked from http://code.google.com/p/winsvc/](https://github.com/btcsuite/winsvc)
* [kardianos/service: Run go programs as a service on major platforms.](https://github.com/kardianos/service)
* [我第1个可用的golang小程序 - 老匡的个人空间 - 开源中国](https://my.oschina.net/u/130746/blog/226050) **☆**
* [golang以服务方式运行 - Minho - 开源中国](https://my.oschina.net/idufei/blog/710885) 
* [golang程序在windows上，注册为服务 - CSDN博客](https://blog.csdn.net/yang8023tao/article/details/53332984)

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

###### seelog

* [cihub/seelog: Seelog is a native Go logging library that provides flexible asynchronous dispatching, filtering, and formatting.](https://github.com/cihub/seelog)

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

* [go-gomail/gomail: The best way to send emails in Go.](https://github.com/go-gomail/gomail)
* [models/Mailer.go · qiqi/mailsrv - 码云 - 开源中国](https://gitee.com/zhucheer/mailsrv/blob/master/models/Mailer.go)
* [Golang 使用gomail包发送邮件 - CSDN博客](https://blog.csdn.net/wj199395/article/details/75206501)


***

#### 配置文件

##### TOML

* [toml-lang/toml: Tom's Obvious, Minimal Language](https://github.com/toml-lang/toml)
* [TOML： GitHub 这是要革 YAML 的命呀！ - SegmentFault 业界资讯 - SegmentFault 思否](https://segmentfault.com/a/1190000000477752)
* [BurntSushi/toml: TOML parser for Golang with reflection.](https://github.com/BurntSushi/toml) **不错**
* [Golang TOML Configs Example: MySQL Connection - JonathanMH](https://jonathanmh.com/golang-toml-configs-example-mysql-connection/)
* [koding/multiconfig: Load configuration from multiple sources in Go](https://github.com/koding/multiconfig)

##### viper

viper太重量级，使用viper时你需要pull另外20个viper依赖的第三方包 ???

* [spf13/viper: Go configuration with fangs](https://github.com/spf13/viper)
* [Golang的配置信息处理框架Viper-RecallSong的博客-51CTO博客](http://blog.51cto.com/13599072/2072753)
* [Golang程序配置方案小结 | Tony Bai](https://tonybai.com/2015/07/01/config-solutions-for-golang-app/)
* [PhalGo-Viper获取配置 - Go语言中文网 - Golang中文社区](https://studygolang.com/articles/4244)


##### yaml

* [go-yaml/yaml: YAML support for the Go language.](https://github.com/go-yaml/yaml)
* [Go 处理yaml类型的配置文件 - 后端 - 掘金](https://juejin.im/entry/5ae29326518825670c45a5b7)
* [kylelemons/go-gypsy: Go YAML Parser for Simple YAML](https://github.com/kylelemons/go-gypsy)
* [Go实战--go语言中使用YAML配置文件(与json、xml、ini对比) - CSDN博客](https://blog.csdn.net/wangshubo1989/article/details/73784907)


##### go test

* [go test 初始化--- TestMain的使用 - CSDN博客](https://blog.csdn.net/lanyang123456/article/details/79776410)


##### os

```
// check directory exist
func IsDirExist(path string) bool {
	file, err := os.Stat(path)

	if err != nil {
		return os.IsExist(err)
	} else {
		return file.IsDir()
	}
}

// check file exist
func IsFileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil || os.IsExist(err)
}

```

##### fsnotify

* [fsnotify/fsnotify: Cross-platform file system notifications for Go.](https://github.com/fsnotify/fsnotify)
* [golang 通过fsnotify监控文件，并通过文件变化重启程序 - 怀素真 - 博客园](https://www.cnblogs.com/jkko123/p/7256927.html)
* [viper/viper.go at master · spf13/viper](https://github.com/spf13/viper/blob/master/viper.go#L261)
* [GoLang fsnotify 实现目录中日志文件大小监控 - 川川籽 - 博客园](https://www.cnblogs.com/yzhch/p/6420625.html)
* [Examples: File monitoring with Golang](https://noypi-linux.blogspot.com/2014/07/file-monitoring-with-golang_4.html)
* [golang 通过fsnotify监控文件，并通过文件变化重启程序 - 怀素真 - 博客园](https://www.cnblogs.com/jkko123/p/7256927.html)


##### howeyc/fsnotify

howeyc/fsnotify fork version 增加了新的api

For each event:

Name
IsCreate()
IsDelete()
IsModify()
IsRename()

* [howeyc/fsnotify: File system notification for Go](https://github.com/howeyc/fsnotify)
* [Go语言监控文件变化小程序. - Go开发社区 | CTOLib码库](https://www.ctolib.com/topics-5472.html)

##### 配置文件热更新

* [Golang configuration file hot update](http://www.fatalerrors.org/a/golang-configuration-file-hot-update.html) **例子**
* [Golang hot configuration reload](http://openmymind.net/Golang-Hot-Configuration-Reload/)
* [golang配置文件热更新 - 挖坑笔记 - SegmentFault 思否](https://segmentfault.com/a/1190000008487440)
* [Golang学习--TOML配置处理 - 疯狂的原始人 - 博客园](http://www.cnblogs.com/CraryPrimitiveMan/p/7928647.html)


##### golang中文件以及文件夹路径相关操作

* [Go实战--golang中文件以及文件夹路径相关操作 - CSDN博客](https://blog.csdn.net/wangshubo1989/article/details/77933654)

***

##### orm

###### xorm


* [首页 - xorm: 简单而强大的 Go 语言ORM框架](http://xorm.io/)
* [首页 |](http://gobook.io/read/github.com/go-xorm/manual-zh-CN/)
* [go-xorm/xorm: Simple and Powerful ORM for Go, support mysql,postgres,tidb,sqlite3,mssql,oracle](https://github.com/go-xorm/xorm)
* [Golang解决XORM的时区问题 - Go中国技术社区 - golang](https://gocn.vip/article/566)

###### gorm

* [GORM - The fantastic ORM library for Golang, aims to be developer friendly.](http://gorm.io/)
* [golang orm 框架之 gorm - 个人文章 - SegmentFault 思否](https://segmentfault.com/a/1190000013216540)
* [jinzhu/gorm: The fantastic ORM library for Golang, aims to be developer friendly](https://github.com/jinzhu/gorm)

###### Beego orm

* [使用beedb库进行ORM开发 · Build web application with Golang](https://astaxie.gitbooks.io/build-web-application-with-golang/zh/05.5.html)
* [ORM 使用 - beego: 简约 & 强大并存的 Go 应用框架](https://beego.me/docs/mvc/model/orm.md)

##### string.Format

* [Equivalent of Python string.format in Go? - Stack Overflow](https://stackoverflow.com/questions/40811117/equivalent-of-python-string-format-in-go)


##### http请求

* [rest - How do I send a JSON string in a POST request in Go - Stack Overflow](https://stackoverflow.com/questions/24455147/how-do-i-send-a-json-string-in-a-post-request-in-go)
* [【GoLang】golang HTTP GET/POST JSON的服务端、客户端示例，包含序列化、反序列化 - junneyang - 博客园](https://www.cnblogs.com/junneyang/p/6211190.html)
* [Golang语言Post发送 json形式的请求 - Corwien - SegmentFault 思否](https://segmentfault.com/a/1190000008624328)
* [golang net/http包使用 · Golang学习室 · 看云](https://www.kancloud.cn/digest/batu-go/153529)

##### json操作

* 

##### govendor

* [用 govendor 管理 Golang 第三方套件](https://seans.tw/2017/govendor-intro/)
* [[Golang] govendor的使用 - 胡伟煌 | Blog](http://www.huweihuang.com/article/golang/govendor-usage/)
* [go 依赖管理利器 -- govendor - CSDN博客](https://blog.csdn.net/yeasy/article/details/65935864)
* [Go包管理工具Vendor使用教程 - CSDN博客](https://blog.csdn.net/benben_2015/article/details/80614873)
* [govendor Golang依赖管理 | 能孩子](http://boyneng.com/2017/09/03/govendor%20golang%E4%BE%9D%E8%B5%96%E7%AE%A1%E7%90%86/)
* [kardianos/govendor: Go vendor tool that works with the standard vendor file.](https://github.com/kardianos/govendor)

```
go get -u -v github.com/kardianos/govendor
```


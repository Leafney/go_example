package main

/*
载入默认配置文件和自定义配置文件

* [BurntSushi/toml: TOML parser for Golang with reflection.](https://github.com/BurntSushi/toml)
* [Golang TOML Configs Example: MySQL Connection - JonathanMH](https://jonathanmh.com/golang-toml-configs-example-mysql-connection/)
* [akkagao/echodemo: 学习使用传说中比其他Go web框架性能高10倍的ECHO](https://github.com/akkagao/echodemo)
* [TOML： GitHub 这是要革 YAML 的命呀！ - SegmentFault 业界资讯 - SegmentFault 思否](https://segmentfault.com/a/1190000000477752)

*/


import (
	"bytes"
	"fmt"
	"github.com/15125505/zlog/log"
	"github.com/BurntSushi/toml"
	"github.com/fsnotify/fsnotify"
	"go_service/govpn/util"
	"io/ioutil"
	"path/filepath"
)

/*
type

*/

type vpnConfig struct {
	Title       string
	LogPath     logPath
	Runtime     runTime
	CheckStatus checkStatus
	Email       email
	DingDing    dingDing
}

type logPath struct {
	Log string
}

// 运行时间
type runTime struct {
	StartCron string `toml:"cron_start"`
	StopCron  string `toml:"cron_stop"`
}

// 检测运行状态
type checkStatus struct {
	Enable    bool   `toml:"enable"`
	DbPath    string `toml:"db_path"`
	CheckCron string `toml:"cron_check"`
	StartCron string `toml:"check_start"`
	StopCron  string `toml:"check_stop"`
	DefStatus int    `toml:"def_status"`
}

// 邮件通知
type email struct {
	Enable    bool     `toml:"enable"`
	EServer   string   `toml:"server"`
	EPort     int      `toml:"port"`
	ESsl      bool     `toml:"ssl"`
	EUserName string   `toml:"username"`
	EPassWord string   `toml:"password"`
	ETos      []string `toml:"to_users"`
}

// 钉钉通知
type dingDing struct {
	Enable       bool   `toml:"enable"`
	DRootUrl     string `toml:"robot_url"`
	DAccessToken string `toml:"access_token"`
	DMsgPrefix   string `toml:"msg_prefix"`
}

// 配置文件的默认名称
const (
	DefConfigName string = "vpnconfig.toml"
)

var Config vpnConfig

/*
设置默认必需配置项
*/
func defSetting() {
	// 参考自：https://github.com/BurntSushi/toml/issues/47

	Config = vpnConfig{
		Title: "vpn_monitor",
		LogPath: logPath{
			Log: "logs/vpn.log",
		},
		Runtime: runTime{
			StartCron: "0 0 6 * * *",
			StopCron:  "0 0 22 * * *",
		},
		CheckStatus: checkStatus{
			Enable: false,
		},
		Email: email{
			Enable: false,
		},
		DingDing: dingDing{
			Enable: false,
		},
	}
}

func InitConfig() {

	// 载入默认必需配置项
	defSetting()

	// 判断是否存在配置文件，如果不存在，则创建一个默认的配置文件，同时载入默认配置项
	if !util.Exist(DefConfigName) {
		fmt.Println("*** 未发现配置文件，将创建默认配置文件！ ***")
		// 创建
		createConfigforDefault()
	}

	// 载入配置文件
	if _, err := toml.DecodeFile(DefConfigName, &Config); err != nil {
		fmt.Println(err)
		panic("*** 配置文件读取异常！ ***")
		return
	}

	fmt.Println("*** 配置文件载入成功！ ***")
}

// 默认配置文件格式内容
func createConfigforDefault() {
	/*
		参考：https://github.com/BurntSushi/toml/blob/master/encode_test.go
	*/

	var config = map[string]interface{}{
		"title": "vpn_monitor",
		"logpath": map[string]string{
			"log": "logs/vpn.log",
		},
		"runtime": map[string]string{
			"cron_start": "0 0 6 * * *",
			"cron_stop":  "0 0 22 * * *",
		},
		"checkstatus": map[string]interface{}{
			"enable":      false,
			"db_path":     "",
			"cron_check":  "0/1 * * * * *",
			"check_start": "0 0 6 * * *",
			"check_stop":  "0 0 22 * * *",
			"def_status":  0,
		},
		"email": map[string]interface{}{
			"enable":   false,
			"server":   "",
			"port":     465,
			"ssl":      true,
			"username": "",
			"password": "",
			"to_users": []string{""},
		},
		"dingding": map[string]interface{}{
			"enable":       false,
			"robot_url":    "https://oapi.dingtalk.com/robot/send",
			"access_token": "",
			"msg_prefix":   "",
		},
	}

	// 参考自：https://github.com/BurntSushi/toml/blob/master/encode_test.go#L603
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(config); err != nil {
		fmt.Println(err)
	}

	// 打印输出
	// fmt.Println(buf.String())

	// 创建默认配置文件
	bts := []byte(buf.String())
	if err := ioutil.WriteFile(DefConfigName, bts, 0644); err != nil {
		panic(err)
	}
}

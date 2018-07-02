package main

/*
热更新配置文件的测试

* [GoLang fsnotify 实现目录中日志文件大小监控 - 川川籽 - 博客园](https://www.cnblogs.com/yzhch/p/6420625.html)
* [viper/viper.go at master · spf13/viper](https://github.com/spf13/viper/blob/master/viper.go#L261)
* [golang 通过fsnotify监控文件，并通过文件变化重启程序 - 怀素真 - 博客园](https://www.cnblogs.com/jkko123/p/7256927.html)

*/


func WatchConfig(configFile string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Error(err)
	}
	defer watcher.Close()

	//要监控的文件绝对路径
	filename, err := filepath.Abs(configFile)
	if err != nil {
		log.Error(err)
		return
	}

	cfgFile := filepath.Clean(filename)
	//要监控的配置文件所在目录的绝对路径
	cfgDir, _ := filepath.Split(cfgFile)

	log.Info(cfgFile)
	log.Info(cfgDir)

	done := make(chan bool)
	// 另起一个goroutine来处理监控对象的事件
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Info("file:", event.Name, event.Op.String())

				// we only care about the config file
				if filepath.Clean(event.Name) == cfgFile {
					// 新增或读取时
					if event.Op&fsnotify.Create == fsnotify.Create {
						log.Info("新增配置文件")
						fmt.Println("新增了")
					}
					if event.Op&fsnotify.Write == fsnotify.Write {
						log.Info("配置文件被修改")
						fmt.Println("修改了")
					}
					// 删除时
					if event.Op&fsnotify.Remove == fsnotify.Remove {
						log.Info("配置文件被删除")
						fmt.Println("删除了")
						// 判断文件是否还存在，不存在了就表示删除了
						if !util.Exist(cfgFile) {
							watcher.Remove(event.Name)
							fmt.Println("删除监控")
						}
					}
					//重命名
					if event.Op&fsnotify.Rename == fsnotify.Rename {
						//Mac系统下 “移到废纸篓” 会触发Rename事件

						log.Info("配置文件被重命名")
						fmt.Println("重命名了")
						// 判断文件是否还存在，不存在了就表示重命名了
						if !util.Exist(cfgFile) {
							watcher.Remove(event.Name)
							fmt.Println("删除监控")
						}
					}
				}

			case err := <-watcher.Errors:
				log.Error("error:", err)
			}
		}
	}()

	// 添加监控目录
	err = watcher.Add(cfgDir)
	if err != nil {
		log.Error(err)
	}
	<-done
}


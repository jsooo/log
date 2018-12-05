# log
golang log库

提供了按照日志等级`info`,`debug`,`error`,`warning`,`fatal`进行隔离的配置

```  golang
package main

import (
	golog "log"
	"os"

	"github.com/jsooo/log"
)

func main() {
	errorFile, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Errorf("打开日志文件失败：%v", err)
		return
	}
	//设置error等级日志Output
	log.SetLogOutput(errorFile, log.LevelError)
	//设置error等级日志Flag
	log.SetLogFlag(golog.Ldate|golog.Ltime, log.LevelError)
	//设置error等级日志Prefix
	log.SetLogPrefix("error - ", log.LevelError)

	log.Warn("log - warn")
	log.Warnf("log - %s", "warn")
	log.Error("log - error")
	log.Errorf("log - %s", "error")
	log.Info("log - info")
	log.Infof("log - %s", "info")
	log.Debug("log - debug")
	log.Debugf("log - %s", "debug")
	log.Fatal("log - fatal")
	log.Fatalf("log - %s", "fatal")

}

```

输出：
```
2018/12/05 21:28:05 new-make.go:20: [WARN] log - warn
2018/12/05 21:28:05 new-make.go:21: [WARN] log - warn
2018/12/05 21:28:05 new-make.go:24: [INFO] log - info
2018/12/05 21:28:05 new-make.go:25: [INFO] log - info
2018/12/05 21:28:05 new-make.go:26: [DEBUG] log - debug
2018/12/05 21:28:05 new-make.go:27: [DEBUG] log - debug
2018/12/05 21:28:05 new-make.go:28: [FATAL] log - fatal
```

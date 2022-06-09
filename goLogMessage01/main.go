package main

import (
	"fmt"
	tail "github.com/hpcloud/tail"
	"goLogMessage/controller"
	"goLogMessage/entity"
	"strings"
	"sync"
)

func main() {
	// 获取配置文件结构体
	var c entity.Conf
	conf := c.GetConf()
	fileList := strings.Split(conf.FileName, ",")
	//fmt.Println(conf)
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // 是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的哪个地方开始读
		MustExist: false,                                // 文件不存在不报错
		Poll:      true,                                 // 监听新行，使用tail -f，这个参数非常重要
	}
	// waitGroup等待所有协程退出
	var wg sync.WaitGroup
	wg.Add(len(fileList))
	for _, s := range fileList {
		tailFile, err := tail.TailFile(s, config)
		if err != nil {
			fmt.Printf("tail file failed,file: %v err: %v/n", s, err)
			return
		}
		go controller.GetMessage(tailFile, wg)
	}
	wg.Wait()
}

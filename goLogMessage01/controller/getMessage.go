package controller

import (
	"fmt"
	"github.com/hpcloud/tail"
	"goLogMessage/entity"
	"goLogMessage/service"
	"strings"
	"sync"
	"time"
)

func GetMessage(tailFile *tail.Tail, wg sync.WaitGroup) {
	var line *tail.Line
	var ok bool
	// for循环一直遍历执行，监听文件的变化
	for {
		line, ok = <-tailFile.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tailFile.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println(line.Text)
		text := line.Text
		if strings.Contains(text, "error") {
			var message entity.Message
			message.FileName = tailFile.Filename
			message.Body = text
			message.GetTime = time.Now()
			service.NewSendMessage("mail").Send(message)
			service.NewSendMessage("note").Send(message)
		}
	}
	wg.Done()
}

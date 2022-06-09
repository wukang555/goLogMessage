package main

import (
	"fmt"
	"net/smtp"
)

func main() {
	// 一些参数
	subject := fmt.Sprintf("Subject: %s\r\n", "日志异常提醒")
	send := fmt.Sprintf("From: %s 测试发件邮箱\r\n", "1205995782.com")
	receiver := fmt.Sprintf("To: %s\r\n", "67363085@qq.com")
	contentType := "Content-Type: text/plain" + "; charset=UTF-8\r\n\r\n"
	// 拼接body
	content := "你好，这里是邮件的内容..."
	msgMail := []byte(subject + send + receiver + contentType + content)
	addr := "smtp.qq.com:25"
	auth := smtp.PlainAuth("", "1205995782@qq.com", "rdzzvkpcuifyggdg", "smtp.qq.com")
	from := "1205995782@qq.com"
	to := []string{"67363085@qq.com"}
	// 发送邮件
	err := smtp.SendMail(addr, auth, from, to, msgMail)
	fmt.Println(err)
}

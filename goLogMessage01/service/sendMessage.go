package service

import (
	"fmt"
	"goLogMessage/entity"
	"net/smtp"
)

func NewSendMessage(str string) SendMessage {
	switch str {
	case "mail":
		return SendMail{}
	case "note":
		return SendDChat{}
	}
	return SendMail{}

}

// SendMessage 发送消息的接口
type SendMessage interface {
	Send(entity.Message) bool
}

// SendMail 邮件实现接口
type SendMail struct{}

// Send sendMail函数实现接口中的send方法
func (fMail SendMail) Send(msg entity.Message) bool {
	var c = entity.Conf{}
	emailConf := c.GetConf().Email

	fmt.Println(msg)
	// 发送邮件配置
	subject := fmt.Sprintf("Subject: %s\r\n", emailConf.Subject)
	send := fmt.Sprintf("From: %s 测试发件邮箱\r\n", emailConf.SenMail)
	receiver := fmt.Sprintf("To: %s\r\n", emailConf.AcceptMail)
	contentType := "Content-Type: text/plain" + "; charset=UTF-8\r\n\r\n"
	// 拼接body
	content := "文件：" + msg.FileName + ",存在error日志行为：" + msg.Body + "，出现时间：" + msg.GetTime.String()

	msgMail := []byte(subject + send + receiver + contentType + content)
	addr := "smtp.qq.com:25"
	auth := smtp.PlainAuth("", emailConf.SenMail, emailConf.Password, "smtp.qq.com")
	from := emailConf.SenMail
	to := []string{emailConf.AcceptMail}
	// 发送邮件
	err := smtp.SendMail(addr, auth, from, to, msgMail)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// SendDChat DChat实现接口
type SendDChat struct{}

func (fChat SendDChat) Send(msg entity.Message) bool {
	fmt.Println(msg)
	return true
}

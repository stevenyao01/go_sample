package smoke

import (
	"bytes"
	"fmt"
	"github.com/jordan-wright/email"
	"html/template"
	"net/smtp"
	"time"
)

type mail struct {
	from string
	to   []string
	sub  string
	att  []string
}

func (m *mail) SendMail(msg string) error {
	// NewEmail返回一个email结构体的指针
	e := email.NewEmail()
	// 发件人
	e.From = m.from
	// 收件人(可以有多个)
	//e.To = []string{toUser}
	e.To = m.to
	// 邮件主题
	e.Subject = m.sub
	// 解析html模板
	t, err := template.ParseFiles("config/email-template.html")
	if err != nil {
		return err
	}
	// Buffer是一个实现了读写方法的可变大小的字节缓冲
	body := new(bytes.Buffer)
	// Execute方法将解析好的模板应用到匿名结构体上，并将输出写入body中
	_ = t.Execute(body, struct {
		FromUserName string
		ToUserName   string
		TimeDate     string
		Message      string
	}{
		FromUserName: "LeapIot.Org",
		ToUserName:   "姚海平",
		TimeDate:     time.Now().Format("2006/01/02"),
		Message:      msg,
	})
	// html形式的消息
	e.HTML = body.Bytes()
	// 从缓冲中将内容作为附件到邮件中
	//_, _ = e.Attach(body, "email-template.html", "text/html")
	// 以路径将文件作为附件添加到邮件中
	for _, v := range m.att{
		if v != "" {
			_, errAttachFile := e.AttachFile(v)
			if errAttachFile != nil {
				fmt.Println("errAttachFile: ", errAttachFile.Error())
			}
		}
	}
	//_, errAttachFile := e.AttachFile("/home/steven/code/go/src/github.com/go_sample/src/autosmoke/agentSign/EdgeAgentlinux64/EdgeAgentlinux64.log")
	//if errAttachFile != nil {
	//	fmt.Println("errAttachFile: ", errAttachFile.Error())
	//}

	// 发送邮件(如果使用QQ邮箱发送邮件的话，passwd不是邮箱密码而是授权码)
	return e.Send("smtp.126.com:25", smtp.PlainAuth("", "leapiot@126.com", "IESSAVSGWFXNPQMO", "smtp.126.com"))
}

func MailNew(f string, t []string, s string, at []string) (*mail, error) {
	return &mail{
		from: f,
		to:   t,
		sub:  s,
		att:  at,
	}, nil
}

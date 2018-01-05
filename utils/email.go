package utils

import (
	"fmt"
	"strings"
	"net/smtp"
)


/*
 *  user : example@example.com login smtp server user
 *  password: xxxxx login smtp server password
 *  host: smtp.example.com:port   smtp.163.com:25
 *  to: example@example.com;example1@163.com;example2@sina.com.cn;...
 *  subject:The subject of mail
 *  body: The content of mail
 *  mailtyoe: mail type html or text
 */


func sendMail(user, password, host, to, subject, body, mailtype string) error{
	hp := strings.Split(host, ":")
	auth := unencryptedAuth{smtp.PlainAuth("", user, password, hp[0])}
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/"+ mailtype + "; charset=UTF-8"
	}else{
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<"+ user +">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

type unencryptedAuth struct{
	smtp.Auth
}
func (a unencryptedAuth) Start(server *smtp.ServerInfo)(string,[]byte,error){
	s :=*server
	s.TLS=true
	return a.Auth.Start(&s)
}

func SSendRegisterEmail(to string,resadd string) error{
	user := "longlee@mylonglee.com"
	password := "longlee11fF"
	host := "smtp.mylonglee.com:25"

	subject := "注册认证邮件"

	body := `
    <html>
    <body>
    <h3>
    <a href="http://localhost/email/reg?code=%s&email=%s">请点击这段文字进行注册确认</a>
	如无法打开请复制URL在浏览器中打开 http://localhost/email/reg?code=%s&email=%s
    </h3>
    </body>
    </html>
    `
	body=fmt.Sprintf(body,resadd,to,resadd,to)
	go sendMail(user, password, host, to, subject, body, "html")

	return nil
}


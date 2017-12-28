//package main
//
//import(
//	"fmt"
//	"net/smtp"
//	"strings"
//)
//
//func main(){
//	auth:=smtp.PlainAuth("","longlee@mylonglee.com","longlee11fF","smtp.mxhichina.com")
//	to :=[]string{"762370895@qq.com"}
//	nickname :="test"
//	user:="longlee@mylonglee.com"
//	subject :="test mail"
//	content_type :="Content-Type: text/plain; charset=UTF-8"
//
//	body:="this is the mail body"
//	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
//		"<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
//	err :=smtp.SendMail("smtp.mxhichina.com:25",auth,user,to,msg)
//	if err !=nil{
//		fmt.Printf("send mail error %v",err)
//	}
//}

package main
import (
	"net/smtp"
	"fmt"
	"strings"
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


func SendMail(user, password, host, to, subject, body, mailtype string) error{
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

func main() {
	user := "longlee@mylonglee.com"
	password := "longlee11fF"
	host := "smtp.mylonglee.com:25"
	to := "762370895@qq.com"

	subject := "Test send email by golang"

	body := `
    <html>
    <body>
    <h3>
    "Test send email by golang"
    </h3>
    </body>
    </html>
    `
	fmt.Println("send email")
	err := SendMail(user, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("send mail error!")
		fmt.Println(err)
	}else{
		fmt.Println("send mail success!")
	}

}

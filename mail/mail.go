package mail

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"
	"strings"
)

type Mail struct {
	senderId string
	toIds    []string
	subject  string
	template string
	data     interface{}
}

type SmtpServer struct {
	host string
	port string
}

func (s *SmtpServer) ServerName() string {
	return s.host + ":" + s.port
}

func (mail *Mail) BuildMessage() bytes.Buffer {

	var body bytes.Buffer
	headers := "MIME-version: 1.0;\nContent-Type: text/html;"
	body.Write([]byte(fmt.Sprintf("From: %s\r\n",  mail.senderId)))
	body.Write([]byte(fmt.Sprintf("To: %s\r\n", strings.Join(mail.toIds, ";"))))
	body.Write([]byte(fmt.Sprintf("Subject: %s\n%s\n\n", mail.subject, headers)))
	wd, _ := os.Getwd()
	t, _ := template.ParseFiles(wd+"/mail/templates/"+mail.template)

	t.Execute(&body, mail.data)

	return body
}

func SendMail(toIds []string, subject string, template string, data interface{}) {
	var Host = os.Getenv("MAIL_HOST")
	var Port = os.Getenv("MAIL_PORT")
	var Username = os.Getenv("MAIL_USERNAME")
	var Password = os.Getenv("MAIL_PASSWORD")

	mail := Mail{}
	mail.senderId = Username
	mail.toIds = toIds
	mail.subject = subject
	mail.template = template
	mail.data = data

	messageBody := mail.BuildMessage()

	smtpServer := SmtpServer{host: Host, port: Port}

	auth := smtp.PlainAuth("", mail.senderId, Password, smtpServer.host)

	client, err := smtp.Dial(smtpServer.ServerName())
	if err != nil {
		log.Panic(err)
	}
	if ok, _ := client.Extension("STARTTLS"); ok {
		config := &tls.Config{ServerName: smtpServer.host}
		if err = client.StartTLS(config); err != nil {
			log.Panic(err)
		}
	}

	if err = client.Auth(auth); err != nil {
		log.Panic(err)
	}

	if err = client.Mail(Username); err != nil {
		log.Panic(err)
	}
	for _, k := range toIds {
		if err = client.Rcpt(k); err != nil {
			log.Panic(err)
		}
	}

	w, err := client.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write(messageBody.Bytes())
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	client.Quit()

	log.Println("Mail sent successfully")
}
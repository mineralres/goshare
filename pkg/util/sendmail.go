package util

import (
	"encoding/base64"
	"net/smtp"
	"strings"
)

// SendMail 发送邮件的逻辑函数
func SendMail(user, password, host, to, subject, content string) error {
	mailtype := ""
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var contentType string
	if mailtype == "html" {
		contentType = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	// 演示base64编码
	encodeString := base64.StdEncoding.EncodeToString([]byte(subject))

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject:=?UTF-8?B?" + encodeString + "?=\r\n" + contentType + "\r\n\r\n" + content)
	sendTo := strings.Split(to, ";")
	var filter []string
	for _, t := range sendTo {
		if t != "" {
			filter = append(filter, t)
		}
	}
	err := smtp.SendMail(host, auth, user, filter, msg)
	return err
}

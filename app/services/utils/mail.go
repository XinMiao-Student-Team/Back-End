package utils

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

func SendMail(mailto, subject, body string) {
	m := gomail.NewMessage()
	// 发送人
	m.SetHeader("From", "3190381602@qq.com")
	// 接收人
	m.SetHeader("To", "13567713788@163.com")
	// 抄送人
	//m.SetAddressHeader("Cc", "xxx@qq.com", "xiaozhujiao")
	// 主题
	m.SetHeader("Subject", "a Text form ZeroHzzzz")
	// 内容
	m.SetBody("text/html", "a test email")
	// 附件
	//m.Attach("./myIpPic.png")
	// 连接
	d := gomail.NewDialer("smtp.qq.com", 587, "3190381602@qq.com", "pjfwbwwlldshdeee")

	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("DialAndSend err %v:", err)
		panic(err)
	}
	fmt.Printf("send mail success\n")
}

package userServices

import (
	"xinmiao/app/config"
	"xinmiao/app/models"
	"xinmiao/app/utils"

	"gopkg.in/gomail.v2"
)

func DecryptUserKeyInfo(user *models.User) {
	key := config.GetEncryptKey()
	if user.Password != "" {
		slt := utils.AesDecrypt(user.Password, key)
		user.Password = slt[0 : len(slt)-len(user.Password)]
	}
}

func SendMail(mailto, subject, body string) error {
	m := gomail.NewMessage()
	//发送人
	m.SetHeader("From", "3190381602@qq.com")
	//接收人
	m.SetHeader("To", "13567713788@163.com")
	//抄送人
	//m.SetAddressHeader("Cc", "xxx@qq.com", "xiaozhujiao")
	//主题
	m.SetHeader("Subject", subject)
	//内容
	m.SetBody("text/html", body)
	//附件
	//m.Attach("./myIpPic.png")

	//拿到token，并进行连接,第4个参数是填授权码
	d := gomail.NewDialer("smtp.qq.com", 587, "3190381602@qq.com", "pjfwbwwlldshdeee")

	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

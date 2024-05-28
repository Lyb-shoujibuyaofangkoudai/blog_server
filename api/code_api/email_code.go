package code_api

import (
	"blog_server/global"
	"blog_server/models/res"
	"blog_server/utils"
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
	"io/ioutil"
	"text/template"
	"time"
)

func GetEmailCode(c *gin.Context, email string) {
	// HTML文件的本地路径
	filePath := "api/code_api/email.html"
	// 读取HTML文件内容
	htmlContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		global.Log.Errorf("读取HTML文件失败：%s", err)
		res.FailWithMsg("未知错误", c)
		return
	}

	// 定义模板
	tmpl, err := template.New("html").Parse(string(htmlContent))
	if err != nil {
		global.Log.Errorf("解析HTML文件失败：%s", err)
		res.FailWithMsg("未知错误", c)
		return
	}

	code := global.Redis.Get(context.Background(), email).Val()
	if code != "" {
		res.FailWithMsg("验证码已发送，请1分钟后再试", c)
		return
	}

	// 随机生成6位数数字验证码
	code = utils.GenerateRandomCode(6)
	// 将验证码存入Redis 时效为1分钟
	global.Redis.Set(context.Background(), email, code, time.Duration(5)*time.Minute)

	// 创建替换后的HTML内容的缓冲区
	var buf bytes.Buffer

	// 执行模板替换占位符
	err = tmpl.Execute(&buf, map[string]string{
		"OPERATION": "邮箱注册",
		"CODE":      code,
	})
	if err != nil {
		res.FailWithMsg("读取HTML文件失败", c)
		return
	}
	message := buf.String()

	// QQ 邮箱：
	// SMTP 服务器地址：smtp.qq.com（SSL协议端口：465/994 | 非SSL协议端口：25）
	// 163 邮箱：
	// SMTP 服务器地址：smtp.163.com（端口：25）
	host := global.Config.Email.Host
	port := global.Config.Email.Port
	userName := global.Config.Email.User
	password := global.Config.Email.AuthorizationCode

	m := gomail.NewMessage()
	m.SetHeader("From", userName) // 发件人
	// m.SetHeader("From", "alias"+"<"+userName+">") // 增加发件人别名

	m.SetHeader("To", email)      // 收件人，可以多个收件人，但必须使用相同的 SMTP 连接
	m.SetHeader("Subject", "验证码") // 邮件主题

	// text/html 的意思是将文件的 content-type 设置为 text/html 的形式，浏览器在获取到这种文件时会自动调用html的解析器对文件进行相应的处理。
	// 可以通过 text/html 处理文本格式进行特殊处理，如换行、缩进、加粗等等
	m.SetBody("text/html", fmt.Sprintf(message, "YB博客"))

	// text/plain的意思是将文件设置为纯文本的形式，浏览器在获取到这种文件时并不会对其进行处理
	// m.SetBody("text/plain", "纯文本")
	// m.Attach("test.sh")   // 附件文件，可以是文件，照片，视频等等
	// m.Attach("lolcatVideo.mp4") // 视频
	// m.Attach("lolcat.jpg") // 照片

	d := gomail.NewDialer(
		host,
		port,
		userName,
		password,
	)
	// 关闭SSL协议认证
	d.TLSConfig = &tls.Config{InsecureSkipVerify: !global.Config.Email.UserSSL}

	if err := d.DialAndSend(m); err != nil {
		global.Log.Errorf("邮件发送失败：%s", err)
		res.FailWithMsg("邮件发送失败", c)
		return
	}
	res.OkWithMsg("邮件发送成功，请前往邮箱查看", c)
}

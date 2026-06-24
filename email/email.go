// Package email 提供邮件发送功能。
//
// 基于 gomail.v2 库实现 SMTP 邮件发送，支持 HTML 格式邮件。
// 通过自动配置（email.enabled=true）从 Environment 读取 SMTP 配置，
// 创建 EmailClient 并注册到 IoC 容器。
//
// 基本用法：
//
//	client := email.NewEmailClient(
//	    email.WithSmtp("smtp.example.com"),
//	    email.WithPort(465),
//	    email.WithUsername("user@example.com"),
//	    email.WithPassword("password"),
//	)
//	client.SendEmail("target@example.com", "Subject", "<h1>Hello</h1>")
package email

import (
	"crypto/tls"
	"fmt"

	"gopkg.in/gomail.v2"
)

// emailConfig SMTP 邮件配置
type emailConfig struct {
	Smtp     string `json:"smtp"`     // SMTP 服务器地址
	Port     int    `json:"port"`     // SMTP 服务器端口
	Username string `json:"username"` // SMTP 用户名
	Password string `json:"password"` // SMTP 密码或授权码
}

// EmailConfigOption 邮件配置选项函数类型
type EmailConfigOption func(*emailConfig)

// EmailClient 邮件客户端
//
// 封装 gomail.Dialer，提供简单的邮件发送功能。
type EmailClient struct {
	emailConfig *emailConfig
}

// WithSmtp 设置 SMTP 服务器地址
func WithSmtp(smtp string) EmailConfigOption {
	return func(config *emailConfig) {
		config.Smtp = smtp
	}
}

// WithPort 设置 SMTP 服务器端口
func WithPort(port int) EmailConfigOption {
	return func(config *emailConfig) {
		config.Port = port
	}
}

// WithUsername 设置 SMTP 用户名
func WithUsername(username string) EmailConfigOption {
	return func(config *emailConfig) {
		config.Username = username
	}
}

// WithPassword 设置 SMTP 密码或授权码
func WithPassword(password string) EmailConfigOption {
	return func(config *emailConfig) {
		config.Password = password
	}
}

// NewEmailClient 创建邮件客户端
//
// 默认配置为 163 邮箱 SMTP 服务器，请通过选项函数替换为实际配置。
func NewEmailClient(options ...EmailConfigOption) *EmailClient {
	config := &emailConfig{
		Smtp:     "smtp.163.com",
		Port:     25,
		Username: "examples@163.com",
		Password: "password",
	}
	for _, option := range options {
		option(config)
	}
	return &EmailClient{emailConfig: config}
}

// SendEmail 发送 HTML 格式邮件
//
// 参数：
//   - target: 收件人邮箱地址
//   - subject: 邮件主题
//   - content: 邮件内容（支持 HTML 格式）
//
// 返回：
//   - error: 发送失败时返回错误
func (email *EmailClient) SendEmail(target, subject, content string) error {
	// 创建一个新的邮件发送器
	d := gomail.NewDialer(email.emailConfig.Smtp, email.emailConfig.Port, email.emailConfig.Username, email.emailConfig.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// 创建邮件
	m := gomail.NewMessage()
	// 设置邮件头
	m.SetHeader("From", email.emailConfig.Username)
	// 设置收件人
	m.SetHeader("To", target)
	// 设置抄送人为自己
	// m.SetAddressHeader("Cc", username, "admin")
	// 设置邮件主题
	m.SetHeader("Subject", subject)
	// 设置邮件内容，支持html格式
	m.SetBody("text/html", content)
	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("邮件发送失败: %w", err)
	}
	return nil
}

// GetConfig 返回当前配置，用于测试目的
func (email *EmailClient) GetConfig() *emailConfig {
	return email.emailConfig
}

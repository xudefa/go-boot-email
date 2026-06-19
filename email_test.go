package email

import "testing"

func TestSend(t *testing.T) {
	emailClient := NewEmailClient(
		WithSmtp("smtp.163.com"),
		WithPort(25),
		WithUsername("xudefa_163mail@163.com"),
		WithPassword("CMLWVNWHKMHYKBSC"),
	)
	emailClient.SendEmail("1371055523@qq.com", "测试邮件", "这是一封测试邮件")
}

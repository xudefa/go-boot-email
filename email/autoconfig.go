// Package email 提供邮件客户端的自动配置。
//
// 当 email.enabled=true 时自动启用，从 Environment 中读取 email.smtp、email.port、
// email.username、email.password 等配置项，
// 创建并注册 EmailClient Bean 到 IoC 容器中（Bean ID: emailClient）。
package email

import (
	"github.com/xudefa/go-boot/boot"
	"github.com/xudefa/go-boot/condition"
	"github.com/xudefa/go-boot/constants"
	"github.com/xudefa/go-boot/core"
)

// EmailAutoConfiguration 邮件客户端的自动配置
//
// 从 Environment 中读取 email.smtp、email.port、email.username、email.password 等配置项，
// 创建 EmailClient 实例并注册到 IoC 容器中。
// 启用条件：email.enabled=true
type EmailAutoConfiguration struct{}

// init 注册邮件自动配置，由 email.enabled=true 条件控制
func init() {
	boot.RegisterAutoConfig(&EmailAutoConfiguration{},
		condition.OnProperty(constants.EmailEnabled, constants.ConditionTrue),
	)
}

// Configure 执行自动配置逻辑，创建 EmailClient 并注册为 Bean
func (e *EmailAutoConfiguration) Configure(ctx boot.ApplicationContext) error {
	env := ctx.Environment()

	// 构建邮件配置选项
	opts := []EmailConfigOption{
		WithSmtp(env.GetString(constants.EmailSMTP, constants.DefaultEmailSMTP)),
		WithPort(env.GetInt(constants.EmailPort, constants.DefaultEmailPort)),
		WithUsername(env.GetString(constants.EmailUsername, constants.DefaultEmailUsername)),
		WithPassword(env.GetString(constants.EmailPassword, constants.DefaultEmailPassword)),
	}

	// 创建邮件客户端
	emailClient := NewEmailClient(opts...)

	// 注册到 IoC 容器
	if err := ctx.Register(constants.EmailClientBeanID,
		core.Bean(emailClient),
		core.Singleton(),
	); err != nil {
		return err
	}

	return nil
}

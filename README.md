# go-boot-email

[![Go Version](https://img.shields.io/github/go-mod/go-version/xudefa/go-boot-email)](https://go.dev/) [![License](https://img.shields.io/github/license/xudefa/go-boot-email)](./LICENSE) [![Build Status](https://img.shields.io/github/actions/workflow/status/xudefa/go-boot-email/test.yml?branch=master)](https://github.com/xudefa/go-boot-email/actions) [![Go Reference](https://pkg.go.dev/badge/github.com/xudefa/go-boot-email.svg)](https://pkg.go.dev/github.com/xudefa/go-boot-email) [![Go Report Card](https://goreportcard.com/badge/github.com/xudefa/go-boot-email)](https://goreportcard.com/report/github.com/xudefa/go-boot-email)

基于 [go-boot](https://github.com/xudefa/go-boot) 的 SMTP 邮件发送集成模块。将 gomail 无缝集成到 go-boot 的 IoC 容器和自动配置体系中，提供声明式的邮件发送能力。

> 设计理念：遵循 go-boot 的开发规范，将 EmailClient 注册为 Bean，通过自动配置实现零代码启动邮件发送服务。

## 整体架构

```
┌───────────────────────────────────────────────────────────────────────┐
│                    go-boot ApplicationContext                         │
│  ┌───────────┐ ┌──────────────┐ ┌───────────┐ ┌───────────┐           │
│  │ Container │ │  Environment │ │ Lifecycle │ │ EventBus  │           │
│  └───────────┘ └──────────────┘ └───────────┘ └───────────┘           │
│                       ┌─────────────────────┐                         │
│                       │ AutoConfig Registry │                         │
│                       └─────────────────────┘                         │
└───────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
                    ┌───────────────────────────────┐
                    │    go-boot-email Starter      │
                    │  ┌─────────────────────────┐  │
                    │  │ EmailClient Bean        │  │
                    │  │ SMTP Configuration      │  │
                    │  │ TLS/SSL Support         │  │
                    │  └─────────────────────────┘  │
                    └───────────────────────────────┘
```

## 目录

- [快速开始](#快速开始)
- [功能特性](#功能特性)
- [邮件发送](#邮件发送)
- [配置选项](#配置选项)
- [项目结构](#项目结构)
- [开发指南](#开发指南)
- [贡献](#贡献)
- [许可证](#许可证)

## 快速开始

### 安装

```bash
# 安装核心框架
go get github.com/xudefa/go-boot

# 安装邮件集成模块
go get github.com/xudefa/go-boot-email
```

### 最小示例

```go
package main

import (
    "github.com/xudefa/go-boot/boot"
    "github.com/xudefa/go-boot-email"
)

func main() {
    app, err := boot.NewApplication(
        boot.WithAppName("my-email-app"),
        boot.WithVersion("1.0.0"),
    )
    if err != nil {
        panic(err)
    }
    defer app.Stop()

    // 获取 EmailClient（由自动配置注册）
    emailClient := app.Container().Get("emailClient").(*email.EmailClient)

    // 发送邮件
    emailClient.SendEmail(
        "recipient@example.com",
        "欢迎注册",
        "<h1>欢迎使用我们的服务！</h1><p>请点击链接验证邮箱...</p>",
    )

    // 启动应用
    app.Start()

    // 等待终止信号
    app.WaitForSignal()
}
```

## 功能特性

| 特性 | 说明 |
|------|------|
| SMTP 集成 | 基于 gomail.v2 实现 SMTP 邮件发送 |
| 自动配置 | 通过 `email.enabled=true` 自动启用 |
| HTML 邮件 | 支持 HTML 格式邮件内容 |
| 函数式选项 | 灵活的 SMTP 配置（服务器、端口、认证） |
| 依赖注入 | EmailClient 自动注册为 Bean |
| TLS 支持 | 支持 TLS/SSL 加密连接 |

## 邮件发送

### 基本用法

```go
// 通过依赖注入获取 EmailClient
type NotificationService struct {
    EmailClient *email.EmailClient `inject:"emailClient"`
}

func (s *NotificationService) SendWelcomeEmail(userEmail string) {
    s.EmailClient.SendEmail(
        userEmail,
        "欢迎注册",
        "<h1>欢迎！</h1><p>您的账户已创建成功。</p>",
    )
}
```

### 手动创建客户端

```go
client := email.NewEmailClient(
    email.WithSmtp("smtp.example.com"),
    email.WithPort(465),
    email.WithUsername("noreply@example.com"),
    email.WithPassword("your-password"),
)

client.SendEmail(
    "user@example.com",
    "测试邮件",
    "<p>这是一封测试邮件</p>",
)
```

## 配置选项

通过 `boot.WithProperty()` 或配置文件设置：

| 配置项 | 默认值 | 说明 |
|--------|--------|------|
| `email.enabled` | `false` | 是否启用邮件服务 |
| `email.smtp` | `smtp.163.com` | SMTP 服务器地址 |
| `email.port` | `25` | SMTP 服务器端口 |
| `email.username` | - | SMTP 用户名 |
| `email.password` | - | SMTP 密码或授权码 |

### 示例配置

```yaml
# application.yml
email:
  enabled: true
  smtp: smtp.qq.com
  port: 465
  username: noreply@qq.com
  password: your-authorization-code
```

## 项目结构

```
go-boot-email/
├── autoconfig.go           # 自动配置注册
├── email.go                # EmailClient 实现
├── email_test.go           # 单元测试
├── README.md
├── LICENSE
└── go.mod
```

## 开发指南

### 构建

```bash
go build ./...
```

### 测试

```bash
go test ./...
go test -cover ./...       # 带覆盖率
go test -race ./...        # 数据竞争检测
```

### 代码规范

```bash
go fmt ./...
golangci-lint run
```

## 贡献

欢迎提交 Issue 和 Pull Request！详细贡献指南请参阅 [CONTRIBUTING.md](./CONTRIBUTING.md)。

## 许可证

本项目采用 MIT 许可证 — 详情请参阅 [LICENSE](./LICENSE) 文件。
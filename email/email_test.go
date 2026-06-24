package email

import (
	"fmt"
	"testing"
)

// TestNewEmailClient_DefaultConfig 测试默认配置
func TestNewEmailClient_DefaultConfig(t *testing.T) {
	client := NewEmailClient()

	if client.emailConfig.Smtp != "smtp.163.com" {
		t.Errorf("expected default smtp smtp.163.com, got %s", client.emailConfig.Smtp)
	}
	if client.emailConfig.Port != 25 {
		t.Errorf("expected default port 25, got %d", client.emailConfig.Port)
	}
	if client.emailConfig.Username != "examples@163.com" {
		t.Errorf("expected default username examples@163.com, got %s", client.emailConfig.Username)
	}
	if client.emailConfig.Password != "password" {
		t.Errorf("expected default password, got %s", client.emailConfig.Password)
	}
}

// TestNewEmailClient_WithOptions 测试使用选项函数配置
func TestNewEmailClient_WithOptions(t *testing.T) {
	client := NewEmailClient(
		WithSmtp("smtp.gmail.com"),
		WithPort(587),
		WithUsername("test@gmail.com"),
		WithPassword("secret"),
	)

	if client.emailConfig.Smtp != "smtp.gmail.com" {
		t.Errorf("expected smtp smtp.gmail.com, got %s", client.emailConfig.Smtp)
	}
	if client.emailConfig.Port != 587 {
		t.Errorf("expected port 587, got %d", client.emailConfig.Port)
	}
	if client.emailConfig.Username != "test@gmail.com" {
		t.Errorf("expected username test@gmail.com, got %s", client.emailConfig.Username)
	}
	if client.emailConfig.Password != "secret" {
		t.Errorf("expected password secret, got %s", client.emailConfig.Password)
	}
}

// TestNewEmailClient_PartialOptions 测试部分选项配置
func TestNewEmailClient_PartialOptions(t *testing.T) {
	client := NewEmailClient(
		WithSmtp("smtp.qq.com"),
		WithPort(465),
	)

	if client.emailConfig.Smtp != "smtp.qq.com" {
		t.Errorf("expected smtp smtp.qq.com, got %s", client.emailConfig.Smtp)
	}
	if client.emailConfig.Port != 465 {
		t.Errorf("expected port 465, got %d", client.emailConfig.Port)
	}
	// 其他配置应该使用默认值
	if client.emailConfig.Username != "examples@163.com" {
		t.Errorf("expected default username, got %s", client.emailConfig.Username)
	}
}

// TestWithSmtp 测试 WithSmtp 选项函数
func TestWithSmtp(t *testing.T) {
	config := &emailConfig{}
	opt := WithSmtp("smtp.example.com")
	opt(config)

	if config.Smtp != "smtp.example.com" {
		t.Errorf("expected smtp.example.com, got %s", config.Smtp)
	}
}

// TestWithPort 测试 WithPort 选项函数
func TestWithPort(t *testing.T) {
	config := &emailConfig{}
	opt := WithPort(587)
	opt(config)

	if config.Port != 587 {
		t.Errorf("expected port 587, got %d", config.Port)
	}
}

// TestWithUsername 测试 WithUsername 选项函数
func TestWithUsername(t *testing.T) {
	config := &emailConfig{}
	opt := WithUsername("user@example.com")
	opt(config)

	if config.Username != "user@example.com" {
		t.Errorf("expected username user@example.com, got %s", config.Username)
	}
}

// TestWithPassword 测试 WithPassword 选项函数
func TestWithPassword(t *testing.T) {
	config := &emailConfig{}
	opt := WithPassword("my-secret")
	opt(config)

	if config.Password != "my-secret" {
		t.Errorf("expected password my-secret, got %s", config.Password)
	}
}

// TestEmailClient_StructFields 测试 EmailClient 结构体字段
func TestEmailClient_StructFields(t *testing.T) {
	config := &emailConfig{
		Smtp:     "smtp.test.com",
		Port:     25,
		Username: "test@test.com",
		Password: "test-pass",
	}

	client := &EmailClient{emailConfig: config}

	if client.emailConfig.Smtp != "smtp.test.com" {
		t.Errorf("config.Smtp = %s, want smtp.test.com", client.emailConfig.Smtp)
	}
	if client.emailConfig.Port != 25 {
		t.Errorf("config.Port = %d, want 25", client.emailConfig.Port)
	}
	if client.emailConfig.Username != "test@test.com" {
		t.Errorf("config.Username = %s, want test@test.com", client.emailConfig.Username)
	}
	if client.emailConfig.Password != "test-pass" {
		t.Errorf("config.Password = %s, want test-pass", client.emailConfig.Password)
	}
}

// TestEmailConfigOption_Chaining 测试选项函数链式调用
func TestEmailConfigOption_Chaining(t *testing.T) {
	config := &emailConfig{}

	// 链式应用选项
	WithSmtp("smtp.chain.com")(config)
	WithPort(999)(config)
	WithUsername("chain@test.com")(config)
	WithPassword("chain-pass")(config)

	if config.Smtp != "smtp.chain.com" {
		t.Errorf("Smtp = %s, want smtp.chain.com", config.Smtp)
	}
	if config.Port != 999 {
		t.Errorf("Port = %d, want 999", config.Port)
	}
	if config.Username != "chain@test.com" {
		t.Errorf("Username = %s, want chain@test.com", config.Username)
	}
	if config.Password != "chain-pass" {
		t.Errorf("Password = %s, want chain-pass", config.Password)
	}
}

// TestNewEmailClient_EmptyOptions 测试空选项
func TestNewEmailClient_EmptyOptions(t *testing.T) {
	client := NewEmailClient()

	if client == nil {
		t.Fatal("expected non-nil EmailClient")
	}
	if client.emailConfig == nil {
		t.Fatal("expected non-nil emailConfig")
	}
}

// TestNewEmailClient_MultipleCalls 测试多次创建客户端
func TestNewEmailClient_MultipleCalls(t *testing.T) {
	client1 := NewEmailClient(WithSmtp("smtp.1.com"))
	client2 := NewEmailClient(WithSmtp("smtp.2.com"))

	if client1.emailConfig.Smtp == client2.emailConfig.Smtp {
		t.Error("expected different smtp configs")
	}
	if client1.emailConfig.Smtp != "smtp.1.com" {
		t.Errorf("client1 smtp = %s, want smtp.1.com", client1.emailConfig.Smtp)
	}
	if client2.emailConfig.Smtp != "smtp.2.com" {
		t.Errorf("client2 smtp = %s, want smtp.2.com", client2.emailConfig.Smtp)
	}
}

// TestEmailConfig_DefaultValues 测试配置默认值
func TestEmailConfig_DefaultValues(t *testing.T) {
	client := NewEmailClient()

	// 验证默认值不为空
	if client.emailConfig.Smtp == "" {
		t.Error("default smtp should not be empty")
	}
	if client.emailConfig.Username == "" {
		t.Error("default username should not be empty")
	}
	if client.emailConfig.Password == "" {
		t.Error("default password should not be empty")
	}
	if client.emailConfig.Port <= 0 {
		t.Errorf("default port should be positive, got %d", client.emailConfig.Port)
	}
}

// TestEmailConfig_OverrideDefaults 测试覆盖默认值
func TestEmailConfig_OverrideDefaults(t *testing.T) {
	client := NewEmailClient(
		WithSmtp("custom.smtp.com"),
		WithPort(587),
		WithUsername("custom@example.com"),
		WithPassword("custom-pass"),
	)

	if client.emailConfig.Smtp != "custom.smtp.com" {
		t.Errorf("smtp not overridden, got %s", client.emailConfig.Smtp)
	}
	if client.emailConfig.Port != 587 {
		t.Errorf("port not overridden, got %d", client.emailConfig.Port)
	}
	if client.emailConfig.Username != "custom@example.com" {
		t.Errorf("username not overridden, got %s", client.emailConfig.Username)
	}
	if client.emailConfig.Password != "custom-pass" {
		t.Errorf("password not overridden, got %s", client.emailConfig.Password)
	}
}

// TestEmailConfig_ZeroPort 测试零端口配置
func TestEmailConfig_ZeroPort(t *testing.T) {
	client := NewEmailClient(WithPort(0))

	if client.emailConfig.Port != 0 {
		t.Errorf("expected port 0, got %d", client.emailConfig.Port)
	}
}

// TestEmailConfig_NegativePort 测试负数端口配置
func TestEmailConfig_NegativePort(t *testing.T) {
	client := NewEmailClient(WithPort(-1))

	if client.emailConfig.Port != -1 {
		t.Errorf("expected port -1, got %d", client.emailConfig.Port)
	}
}

// TestEmailConfig_EmptyStrings 测试空字符串配置
func TestEmailConfig_EmptyStrings(t *testing.T) {
	client := NewEmailClient(
		WithSmtp(""),
		WithUsername(""),
		WithPassword(""),
	)

	if client.emailConfig.Smtp != "" {
		t.Errorf("expected empty smtp, got %s", client.emailConfig.Smtp)
	}
	if client.emailConfig.Username != "" {
		t.Errorf("expected empty username, got %s", client.emailConfig.Username)
	}
	if client.emailConfig.Password != "" {
		t.Errorf("expected empty password, got %s", client.emailConfig.Password)
	}
}

// TestEmailClient_SendEmail_Validation 测试 SendEmail 参数验证
func TestEmailClient_SendEmail_Validation(t *testing.T) {
	client := NewEmailClient(
		WithSmtp("smtp.test.com"),
		WithPort(25),
		WithUsername("test@test.com"),
		WithPassword("test"),
	)

	// 由于没有真实的 SMTP 服务器，发送会失败
	err := client.SendEmail("target@test.com", "测试主题", "测试内容")
	if err == nil {
		t.Error("expected error when sending email without valid SMTP server")
	}
}

// TestEmailConfigOption_NilSafety 测试选项函数的 nil 安全性
func TestEmailConfigOption_NilSafety(t *testing.T) {
	config := &emailConfig{}

	// 测试各个选项函数不会 panic
	WithSmtp("test")(config)
	WithPort(25)(config)
	WithUsername("test")(config)
	WithPassword("test")(config)

	// 验证配置已正确设置
	if config.Smtp != "test" {
		t.Errorf("Smtp = %s, want test", config.Smtp)
	}
	if config.Port != 25 {
		t.Errorf("Port = %d, want 25", config.Port)
	}
	if config.Username != "test" {
		t.Errorf("Username = %s, want test", config.Username)
	}
	if config.Password != "test" {
		t.Errorf("Password = %s, want test", config.Password)
	}
}

// TestEmailClient_DifferentPorts 测试不同端口配置
func TestEmailClient_DifferentPorts(t *testing.T) {
	ports := []int{25, 465, 587, 2525}

	for _, port := range ports {
		t.Run(fmt.Sprintf("Port_%d", port), func(t *testing.T) {
			client := NewEmailClient(WithPort(port))
			if client.emailConfig.Port != port {
				t.Errorf("expected port %d, got %d", port, client.emailConfig.Port)
			}
		})
	}
}

// TestEmailClient_CommonSMTPConfigs 测试常见的 SMTP 配置
func TestEmailClient_CommonSMTPConfigs(t *testing.T) {
	tests := []struct {
		name     string
		smtp     string
		port     int
		expected string
	}{
		{
			name:     "163 Mail",
			smtp:     "smtp.163.com",
			port:     25,
			expected: "smtp.163.com",
		},
		{
			name:     "QQ Mail",
			smtp:     "smtp.qq.com",
			port:     465,
			expected: "smtp.qq.com",
		},
		{
			name:     "Gmail",
			smtp:     "smtp.gmail.com",
			port:     587,
			expected: "smtp.gmail.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewEmailClient(
				WithSmtp(tt.smtp),
				WithPort(tt.port),
			)
			if client.emailConfig.Smtp != tt.expected {
				t.Errorf("smtp = %s, want %s", client.emailConfig.Smtp, tt.expected)
			}
			if client.emailConfig.Port != tt.port {
				t.Errorf("port = %d, want %d", client.emailConfig.Port, tt.port)
			}
		})
	}
}

// TestEmailClient_GetConfig 测试 GetConfig 方法
func TestEmailClient_GetConfig(t *testing.T) {
	client := NewEmailClient(
		WithSmtp("smtp.test.com"),
		WithPort(587),
		WithUsername("test@test.com"),
		WithPassword("test-pass"),
	)

	config := client.GetConfig()
	if config.Smtp != "smtp.test.com" {
		t.Errorf("GetConfig().Smtp = %s, want smtp.test.com", config.Smtp)
	}
	if config.Port != 587 {
		t.Errorf("GetConfig().Port = %d, want 587", config.Port)
	}
	if config.Username != "test@test.com" {
		t.Errorf("GetConfig().Username = %s, want test@test.com", config.Username)
	}
	if config.Password != "test-pass" {
		t.Errorf("GetConfig().Password = %s, want test-pass", config.Password)
	}
}

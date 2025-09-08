package config

import (
	"testing"

	"os"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	// 重置 Viper 状态，确保测试独立性
	viper.Reset()

	// 创建一个临时配置文件
	configContent := `
auth:
  clientID: "clientID111"
  clientSecret: "clientSecret"
  openAPIHost: "openAPIHost"
plugins:
  - type: "mysql"
    host: "localhost"
    port: 3306
    username: "root"
    password: "root"
    database: "example"
    configKey: "default"
  - type: "mysql"
    host: "localhost"
    port: 3307
    username: "root"
    password: "root"
    database: "example"
    configKey: "default2"
`
	configFile, err := os.CreateTemp("", "config.yaml")
	require.NoError(t, err)
	defer os.Remove(configFile.Name())

	_, err = configFile.Write([]byte(configContent))
	require.NoError(t, err)
	configFile.Close()

	// 设置 Viper 读取临时配置文件
	viper.Reset() // 重置 Viper
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile.Name())

	// 调用 LoadConfig 函数
	err = LoadConfig("") // 使用空字符串测试默认配置
	auth := GetAuthClientConfig()
	require.NoError(t, err)

	// 验证配置内容
	require.Equal(t, "clientID111", auth.ClientID)
	require.Equal(t, "clientSecret", auth.ClientSecret)
	require.Equal(t, "openAPIHost", auth.OpenAPIHost)

	// require.Len(t, config.Plugins, 2)
	// require.Equal(t, "mysql", config.Plugins[0].Type)
	// require.Equal(t, "localhost", config.Plugins[0].ClientPluginOptions.(*v1.MySQLPluginOptions).Host)
	// require.Equal(t, 3306, config.Plugins[0].ClientPluginOptions.(*v1.MySQLPluginOptions).Port)
	// require.Equal(t, "root", config.Plugins[0].ClientPluginOptions.(*v1.MySQLPluginOptions).Username)
	// require.Equal(t, "root", config.Plugins[0].ClientPluginOptions.(*v1.MySQLPluginOptions).Password)
	// require.Equal(t, "example", config.Plugins[0].ClientPluginOptions.(*v1.MySQLPluginOptions).Database)
	// require.Equal(t, "default", config.Plugins[0].ClientPluginOptions.(*v1.MySQLPluginOptions).ConfigKey)

	// require.Equal(t, "mysql", config.Plugins[1].Type)
	// require.Equal(t, "localhost", config.Plugins[1].ClientPluginOptions.(*v1.MySQLPluginOptions).Host)
	// require.Equal(t, 3307, config.Plugins[1].ClientPluginOptions.(*v1.MySQLPluginOptions).Port)
	// require.Equal(t, "root", config.Plugins[1].ClientPluginOptions.(*v1.MySQLPluginOptions).Username)
	// require.Equal(t, "root", config.Plugins[1].ClientPluginOptions.(*v1.MySQLPluginOptions).Password)
	// require.Equal(t, "example", config.Plugins[1].ClientPluginOptions.(*v1.MySQLPluginOptions).Database)
	// require.Equal(t, "default2", config.Plugins[1].ClientPluginOptions.(*v1.MySQLPluginOptions).ConfigKey)
}

func TestLoadConfigWithEnvironment(t *testing.T) {
	// 重置 Viper 状态，确保测试独立性
	viper.Reset()

	// 测试开发环境配置
	devConfigContent := `
auth:
  clientID: "dev_client_id"
  clientSecret: "dev_client_secret"
  openAPIHost: "https://dev.api.example.com"
plugins:
  - type: "mysql"
    host: "dev-db"
    port: 3306
    username: "dev_user"
    password: "dev_pass"
    database: "dev_db"
    configKey: "dev"
`

	// 创建临时开发环境配置文件
	devConfigFile, err := os.CreateTemp("", "config.dev.yaml")
	require.NoError(t, err)
	defer os.Remove(devConfigFile.Name())

	_, err = devConfigFile.Write([]byte(devConfigContent))
	require.NoError(t, err)
	devConfigFile.Close()

	// 重置 Viper 实例
	viper.Reset()

	// 设置配置文件路径
	viper.SetConfigType("yaml")
	viper.SetConfigFile(devConfigFile.Name())

	// 调用 LoadConfig 函数，传入环境参数
	err = LoadConfig("dev")
	require.NoError(t, err)

	auth := GetAuthClientConfig()

	// 验证开发环境配置内容
	require.Equal(t, "dev_client_id", auth.ClientID)
	require.Equal(t, "dev_client_secret", auth.ClientSecret)
	require.Equal(t, "https://dev.api.example.com", auth.OpenAPIHost)
}

func TestLoadConfigDefault(t *testing.T) {
	// 重置 Viper 状态，避免受到之前测试的影响
	viper.Reset()

	// 创建一个默认的临时配置文件用于测试
	defaultConfigContent := `
auth:
  clientID: "default_client_id"
  clientSecret: "default_client_secret"
  openAPIHost: "https://api.dingtalk.com"
`
	defaultConfigFile, err := os.CreateTemp("", "config.yaml")
	require.NoError(t, err)
	defer os.Remove(defaultConfigFile.Name())

	_, err = defaultConfigFile.Write([]byte(defaultConfigContent))
	require.NoError(t, err)
	defaultConfigFile.Close()

	// 设置 Viper 使用临时配置文件
	viper.SetConfigType("yaml")
	viper.SetConfigFile(defaultConfigFile.Name())

	// 测试兼容性函数
	err = LoadConfigDefault()
	require.NoError(t, err)

	auth := GetAuthClientConfig()
	require.Equal(t, "default_client_id", auth.ClientID)
	require.Equal(t, "default_client_secret", auth.ClientSecret)
	require.Equal(t, "https://api.dingtalk.com", auth.OpenAPIHost)
}

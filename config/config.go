// config.go
package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type MySQLConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Database string `mapstructure:"database"`
}

type AppConfig struct {
	MySQLConfigs map[string]MySQLConfig `mapstructure:"mysql"`
}

func (c *AppConfig) InitConfig(filePath string) error {
	if filePath == "" {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
	} else {
		viper.SetConfigFile(filePath)
	}

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(c); err != nil {
		return err
	}
	return nil
}
func (c *AppConfig) GetEnvConfig() string {

	return viper.GetString("app_env")
}

func (c *AppConfig) GetMySQLConfig(env string) MySQLConfig {

	return MySQLConfig{
		Username: viper.GetString(env + ".mysql.username"),
		Password: viper.GetString(env + ".mysql.password"),
		Host:     viper.GetString(env + ".mysql.host"),
		Port:     viper.GetInt(env + ".mysql.port"),
		Database: viper.GetString(env + ".mysql.database"),
	}
}

func (c *AppConfig) GetMysqlConnectionString(env string) string {
	config := c.GetMySQLConfig(env)
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		config.Username, config.Password, config.Host, config.Port, config.Database)
}

// todo
func (c *AppConfig) SchedulerConfig(env string) MySQLConfig {
	return MySQLConfig{
		Username: viper.GetString(env + ".mysql.username"),
		Password: viper.GetString(env + ".mysql.password"),
		Host:     viper.GetString(env + ".mysql.host"),
		Port:     viper.GetInt(env + ".mysql.port"),
		Database: viper.GetString(env + ".mysql.database"),
	}
}

func (c *AppConfig) GetSchedulerConfig(env string) string {
	return ""
}

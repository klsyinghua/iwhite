// config.go
package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// LoadConfig loads the configuration from the specified file path.
func LoadConfig(filePath string) error {
	viper.SetConfigFile(filePath)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

// GetMySQLConfig returns the MySQL configuration parameters.
func GetMySQLConfig() string {
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	host := viper.GetString("mysql.host")
	port := viper.GetInt("mysql.port")
	database := viper.GetString("mysql.models")

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, database)
}

package config

import (
	"authService/pkg/logging"
	"fmt"
	"github.com/spf13/viper"
)

var logger = logging.GetLogger()

type Config struct {
	DatabaseURL string
	Port        string `mapstructure:"PORT"`
	JWTSecret   string `mapstructure:"JWT_SECRET"`
}

func LoadConfig() (*Config, error) {
	viper.AutomaticEnv()

	dbHost := viper.GetString("DB_HOST")
	dbPort := viper.GetString("DB_PORT")
	dbUser := viper.GetString("DB_USER")
	dbPassword := viper.GetString("DB_PASSWORD")
	dbName := viper.GetString("DB_NAME")
	port := viper.GetString("PORT")
	jwtSecret := viper.GetString("JWT_SECRET")

	databaseURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		dbUser, dbPassword, dbHost, dbPort, dbName,
	)

	fmt.Println(databaseURL)

	var cfg Config
	cfg.DatabaseURL = databaseURL
	cfg.Port = port
	cfg.JWTSecret = jwtSecret

	if err := viper.Unmarshal(&cfg); err != nil {
		logger.Errorf("Не удалось распаковать переменные окружения: %v", err)
		return nil, err
	}

	return &cfg, nil
}

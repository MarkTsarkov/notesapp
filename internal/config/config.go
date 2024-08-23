package config

import (
    "github.com/spf13/viper"
)

// Config структура для хранения всех параметров конфигурации
type Config struct {
    App      AppConfig
    Database DatabaseConfig
    Logger   LoggerConfig
}

// AppConfig структура для хранения настроек приложения
type AppConfig struct {
    Name string
    Port int
    Env  string
}

// DatabaseConfig структура для хранения настроек базы данных
type DatabaseConfig struct {
    Host     string
    Port     int
    User     string
    Password string
    DBName   string
    SSLMode  string
}

// LoggerConfig структура для хранения настроек логирования
type LoggerConfig struct {
    Level string
    File  string
}

// LoadConfig загружает конфигурацию из YAML-файла
func LoadConfig(path string) (*Config, error) {
    viper.SetConfigFile(path)
    viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err != nil {
        return nil, err
    }

    var config Config
    if err := viper.Unmarshal(&config); err != nil {
        return nil, err
    }

    return &config, nil
}

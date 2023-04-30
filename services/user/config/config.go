package config

import (
	"github.com/spf13/viper"
	"log"
	"path"
	"strings"
)

func Load(appPath, env string) {
	viper.AddConfigPath("$APP_HOME")
	viper.AddConfigPath(appPath)
	viper.AddConfigPath(".")

	envPath := path.Join(appPath, env+".env")
	viper.SetConfigFile(envPath)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("viper.ReadInConfig returns error: %s", err.Error())
	}

	viper.SetConfigName("override")
	if err := viper.MergeInConfig(); err != nil {
		if !strings.Contains(strings.ToLower(err.Error()), "not found") {
			log.Fatalf("viper.MergeInConfig returns error: %s", err.Error())
		}
	}
	viper.AutomaticEnv()
}

type AppConfig struct {
	Env         string
	Port        int
	GrpcPort    int
	ServiceName string
}

type DatabaseConfig struct {
	DatabaseName string
	DatabaseUrl  string
}

type JwtConfig struct {
	JwtSecretKey string
	JwtExpire    int
}

func GetAppConfig() AppConfig {
	return AppConfig{
		Env:         viper.GetString("APP_ENV"),
		Port:        viper.GetInt("APP_PORT"),
		GrpcPort:    viper.GetInt("APP_GRPC_PORT"),
		ServiceName: viper.GetString("APP_SERVICE_NAME"),
	}
}

func GetDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		DatabaseName: viper.GetString("DATABASE_NAME"),
		DatabaseUrl:  viper.GetString("DATABASE_URL"),
	}
}

func GetJwtConfig() JwtConfig {
	return JwtConfig{
		JwtSecretKey: viper.GetString("JWT_SECRET_KEY"),
		JwtExpire:    viper.GetInt("JWT_EXPIRED"),
	}
}

package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type ConfigEnv struct {
	Port           string
	Env            string
	FirebasePath   string
	DatabaseUrl    string
	RedisUrl       string
	JwtkeysAccess  string
	JwtkeysRefresh string
}

var AppConfig *ConfigEnv

func LoadEnv() {
	_ = godotenv.Load()
	viper.AutomaticEnv()
	AppConfig = &ConfigEnv{
		Port:           Getenv("PORT", "3000"),
		Env:            Getenv("ENV", "dev"),
		FirebasePath:   Getenv("FIREBASE_PATH", ""),
		DatabaseUrl:    Getenv("DB_URL", ""),
		RedisUrl:       Getenv("REDIS_URL", ""),
		JwtkeysAccess:  Getenv("ACCESS_TOKEN", ""),
		JwtkeysRefresh: Getenv("REFRESH_TOKEN", ""),
	}
}

func Getenv(key, defaultString string) string {
	if val := viper.GetString(key); val != "" {
		return val
	}
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultString
}

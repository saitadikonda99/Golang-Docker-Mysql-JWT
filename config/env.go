package config

import (
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port 	 string
	DBUser 	 string
	DBPasswd string
	DBAddr 	 string
	DBName 	 string
}

var GetConfig = InitConfig()

func InitConfig() *Config {
	godotenv.Load()
	return &Config{
		PublicHost: getEnv("PUBLIC_HOST", "localhost"),
		Port: getEnv("PORT", "8080"),
		DBUser: getEnv("DB_USER", "root"),
		DBPasswd: getEnv("DB_PASSWD", "sai@1234"),
		DBAddr: getEnv("DB_ADDR", "localhost:3306"),
		DBName: getEnv("DB_NAME", "go"),
	}
}

func getEnv(key, defaultVal string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	return defaultVal
}
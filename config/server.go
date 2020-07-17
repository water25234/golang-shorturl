package config

import (
	"net"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

var serverConfig *ServerConfig

var appConfig *AppConfig

type ServerConfig struct {
	IpAddress string
}

type AppConfig struct {
	AppLogPath   string
	DBConnection string
	DBHost       string
	DBPort       int
	DBDatabase   string
	DBUsername   string
	DBPassword   string
	RedisHost    string
	RedisPort    string
	RedisDB      int
}

func init() {
	godotenv.Load()
}

func SetServerGonfig() {

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				serverConfig = &ServerConfig{
					IpAddress: ipnet.IP.String(),
				}
			}
		}
	}
}

func SetAppConfig() {
	redisDb, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))

	appConfig = &AppConfig{
		AppLogPath:   os.Getenv("APP_LOG_PATH"),
		DBConnection: os.Getenv("DB_CONNECTION"),
		DBHost:       os.Getenv("DB_HOST"),
		DBPort:       dbPort,
		DBDatabase:   os.Getenv("DB_DATABASE"),
		DBUsername:   os.Getenv("DB_USERNAME"),
		DBPassword:   os.Getenv("DB_PASSWORD"),
		RedisHost:    os.Getenv("REDIS_HOST"),
		RedisPort:    os.Getenv("REDIS_PORT"),
		RedisDB:      redisDb,
	}
}

func GetServerConfig() *ServerConfig {
	return serverConfig
}

func GetAppConfig() *AppConfig {
	return appConfig
}

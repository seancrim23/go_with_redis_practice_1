package application

import "os"

type Config struct {
	RedisAddress  string
	RedisUsername string
	RedisPass     string
	ServerPort    uint16
}

func LoadConfig() Config {
	cfg := Config{
		RedisAddress:  "redis-16968.c89.us-east-1-3.ec2.redns.redis-cloud.com:16968",
		RedisUsername: "default",
		ServerPort:    3000,
	}

	if redisPass, exists := os.LookupEnv("REDIS_PASS"); exists {
		cfg.RedisPass = redisPass
	}

	return cfg
}

package model

import "os"

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	Name     string
}

func NewRedisConfig() RedisConfig {
	return RedisConfig{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     os.Getenv("REDIS_PORT"),
		Name:     os.Getenv("REDIS_DB"),
		Password: os.Getenv("REDIS_PASSWORD"),
	}
}

func NewStreamConfig() RedisConfig {
	return RedisConfig{
		Host:     os.Getenv("STREAM_HOST"),
		Port:     os.Getenv("STREAM_PORT"),
		Name:     os.Getenv("STREAM_DB"),
		Password: os.Getenv("STREAM_PASSWORD"),
	}
}

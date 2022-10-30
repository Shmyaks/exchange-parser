// Package config config from .env file
package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

// EnvStruct struct ENV file
type EnvStruct struct {
	RedisConnection string
	BackendPort     string
	ParseP2PDelay   time.Duration
	ParseSPOTDelay  time.Duration
}

// GetConfig get Env from .env file
func GetConfig() *EnvStruct {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	parseP2PDelay, err := time.ParseDuration(os.Getenv("PARSE_P2P_DELAY"))
	if err != nil {
		panic(err)
	}
	parseSPOTDelay, err := time.ParseDuration(os.Getenv("PARSE_SPOT_DELAY"))
	if err != nil {
		panic(err)
	}

	return &EnvStruct{
		RedisConnection: os.Getenv("REDIS_SERVER"),
		BackendPort:     os.Getenv("BACKEND_PORT"),
		ParseP2PDelay:   parseP2PDelay,
		ParseSPOTDelay:  parseSPOTDelay,
	}
}

// Env auto load
var Env = GetConfig()

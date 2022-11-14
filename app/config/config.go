// Package config config from .env file
package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

// EnvStruct struct ENV file
type EnvStruct struct {
	RedisConnection  string
	BackendPort      string
	ParseP2PDelay    time.Duration
	ParseSPOTDelay   time.Duration
	MinAmountRUB     int
	MinAmounUSD      int
	MinAmoutTRY      int
	MinOrders        int
	PostgresServer   string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
}

// GetConfig get Env from .env file
func GetConfig() *EnvStruct {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	return &EnvStruct{
		RedisConnection:  os.Getenv("REDIS_SERVER"),
		BackendPort:      os.Getenv("BACKEND_PORT"),
		ParseP2PDelay:    convertToTimeDuration(os.Getenv("PARSE_P2P_DELAY")),
		ParseSPOTDelay:   convertToTimeDuration(os.Getenv("PARSE_SPOT_DELAY")),
		MinAmountRUB:     convertToInt(os.Getenv("MIN_AMOUNT_RUB")),
		MinAmounUSD:      convertToInt(os.Getenv("MIN_AMOUNT_USD")),
		MinAmoutTRY:      convertToInt(os.Getenv("MIN_AMOUNT_TRY")),
		MinOrders:        convertToInt(os.Getenv("MIN_ORDERS")),
		PostgresServer:   os.Getenv("POSTGRES_SERVER"),
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresDB:       os.Getenv("POSTGRES_DB"),
	}
}

// Env auto load
var Env = GetConfig()

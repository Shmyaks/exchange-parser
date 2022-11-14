package config

import (
	"strconv"
	"time"
)

func convertToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

func convertToTimeDuration(str string) time.Duration {
	i, err := time.ParseDuration(str)
	if err != nil {
		panic(err)
	}
	return i
}

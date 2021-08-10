package configs

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
)

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func getJsonEnv(key, fallback string) map[string]interface{} {
	value := getEnv(key, fallback)
	var result map[string]interface{}
	err := json.Unmarshal([]byte(value), &result)
	if err != nil {
		log.Fatalln(key, " contains an invalid json: ", value)
	}
	return result
}

var DbType = getEnv("DB_TYPE", "csv")
var DbConfigs = getJsonEnv("DB_CONFIGS", `{"file_path":"ip_to_country/ip2location.csv"}`)
var LimiterTimeFrame, _ = strconv.Atoi(getEnv("LIMITER_TIME_FRAME", "5"))   // In seconds
var LimiterMaxCredits, _ = strconv.Atoi(getEnv("LIMITER_MAX_CREDITS", "5")) //In seconds

package env

import (
	"encoding/base64"
	"os"
	"strconv"
	"strings"
)

func GetEnv(key, defaultVal string) string {
	val, exists := os.LookupEnv(key)
	if exists {
		return strings.TrimSpace(val)
	}
	return defaultVal
}

func GetEnvFromBase64(key, defaultVal string) string {
	val, exists := os.LookupEnv(key)
	if exists {
		rval, err := base64.StdEncoding.DecodeString(val)
		if err != nil {
			return defaultVal
		}
		return strings.TrimSpace(string(rval))
	}
	return defaultVal
}

func GetEnvAsInt(key string, defaultVal int) int {
	val, exists := os.LookupEnv(key)
    val = strings.TrimSpace(val)
	if exists {
		nval, err := strconv.Atoi(val)
		if err != nil {
			return defaultVal
		}
		return nval
	}
	return defaultVal
}

func GetEnvAsBool(key string, defaultVal bool) bool {
	val, exists := os.LookupEnv(key)
    val = strings.TrimSpace(val)
	if exists {
		bval, err := strconv.ParseBool(val)
		if err != nil {
			return defaultVal
		}
		return bval
	}
	return defaultVal
}

func GetEnvAsSlice(key string, defaultVal []string, delim string) []string {
	val, exists := os.LookupEnv(key)
	if exists {
		sval := strings.Split(val, delim)
		return sval
	}
	return defaultVal
}

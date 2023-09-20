package utils

import (
	"os"
	"strconv"
)

func GetPort() int {
	env := os.Getenv("PORT")
	i, err := strconv.Atoi(env)
	Check(err)
	return i
}

func GetApplicationIdentifier() string {
	return os.Getenv("APP_ID")
}

func GetConfigurationProfileId() string {
	return os.Getenv("CONFIG_PROFILE_ID")
}

func GetConfigurationEnvironment() string {
	return os.Getenv("CONFIG_ENVIRONMENT")
}

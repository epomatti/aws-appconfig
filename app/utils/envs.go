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

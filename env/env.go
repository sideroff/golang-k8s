package env

import (
	"errors"
	"fmt"
	"os"
)

// Get - returns the environemnt variable if set; else empty string
func Get(key string) string {
	return os.Getenv(key)
}

// Require - returns the environemnt variable if set; else panics
func Require(key string) string {
	variable, isSet := os.LookupEnv(key)

	if(!isSet){
		panic(errors.New(fmt.Sprintf("Environment variable %s is not set", key)))
	}

	return variable
}

// WithDefault - returns the environemnt variable if set; else defaultValue
func WithDefault(key string, defaultValue string) (string) {
	variable, isSet := os.LookupEnv(key)

	if(!isSet){
		return defaultValue
	}

	return variable
}
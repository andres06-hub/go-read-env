package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// use viper package to load/read the config file or .env file and
// return the value of the key
func viperEnvVariable(key string) string {
	viper.SetConfigFile("db.env")
	// Find and read the config file
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	// viper.Get() returns an empty interface{}
	// to get the underlying type of the key,
	// we have to do the type assertion, we know the underlying value is string
	// if we type assert to other type it will throw an error
	value, ok := viper.Get(key).(string)
	// If the type is a string then ok will be true
	// ok will make sure the program not break
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	return value
}

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	// validate if there is an error reading the file
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func main() {
	value1 := goDotEnvVariable("GREETINGS")
	value2 := viperEnvVariable("URL_DB")

	fmt.Printf("value 1: %v\n", value1)
	fmt.Printf("value 2: %v\n", value2)
}

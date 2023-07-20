package main 

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// Using viper package to read .env file/files
// The aim is to return the value of the key

func viperEnvVariable(key string) string {
	// SetConfigFile explicitly defines the path, name and extension of the config file
	// viper will use this and not check any of the config paths
	// .env - It will search for the .env file in the current directory 


	viper.SetConfigFile("configs.env")

	// Finding and reading the config file 
	err := viper.ReadInConfig()


	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	// viper.Get() returns an empty interface{}
	// to get the underlying type of the key,
	// we have to do the type asertion, we know the underlying value is string
	// if we typ assert to other type it will throw an error

	value, ok := viper.Get(key).(string)

	// If the type is a string then ok will be true 
	// ok will make sure the program not breaks
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value 
}

func main() {
	viperenv := viperEnvVariable("STRONGEST_AVENGER")

	fmt.Printf("viper ; %s = %s \n", "STRONGEST_AVENGER", viperenv)
}
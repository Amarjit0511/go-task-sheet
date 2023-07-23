package main 

import (
	"fmt"
	"log"
	"github.com/spf13/viper"
)

// Using viper package to read the env file
func tomlEnvVariable(key string) string {

	viper.SetConfigFile("config.toml")
	viper.AddConfigPath(".")

	// Finding and reading from the config file 
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Sorry! Error while reading the config file %s", err)
	}
	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatalf("Invalid type")
	}
	return value 
}

func main() {
	toml := tomlEnvVariable("name")
	fmt.Printf("viper ; %s = %s \n", "name", toml)
}

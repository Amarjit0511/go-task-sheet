package main 

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func viperEnvVariable(key string) string {
	viper.SetConfigFile("configs.env")

	// Finding and reading the config file 
	err := viper.ReadInConfig()


	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	value, ok := viper.Get(key).(string)

	// ok will make sure the program not breaks
	if !ok {
		log.Fatalf("Invalid type")
	}
	return value 
}

func main() {
	viperenv := viperEnvVariable("name")
	fmt.Printf("%s\n", viperenv)
}

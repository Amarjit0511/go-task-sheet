package main

import (
	"os"
	"fmt"
	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {
	//Loading env file
	err := godotenv.Load("config.env")
	checkNilError(err)
	return os.Getenv(key)
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	dotenv := goDotEnvVariable("DB_CONNECTION_STRING")
	fmt.Println("DB_CONNECTION_STRING is %s\n", dotenv)
}
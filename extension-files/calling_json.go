package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Bio struct {
	Name     string            `mapstructure:"name"`
	Age      int               `mapstructure:"age"`
	Hometown string            `mapstructure:"hometown"`
	Intro    string            `mapstructure:"intro"`
	Hobbies  []string          `mapstructure:"hobbies"`
	Tasks    map[string]string `mapstructure:"tasks"`
	Task     []map[string]struct {
		OnTime          string `mapstructure:"ontime"`
		ReviewCompleted string `mapstructure:"review_completed"`
	} `mapstructure:"task"`
}

func main() {
	// Set the path to your bio.json file
	viper.SetConfigFile("bio.json")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %s", err)
		return
	}

	var bio Bio
	if err := viper.Unmarshal(&bio); err != nil {
		fmt.Printf("Error unmarshaling config file: %s", err)
		return
	}

	// Access the values from the bio struct
	fmt.Println("Name:", bio.Name)
	fmt.Println("Age:", bio.Age)
	fmt.Println("Hometown:", bio.Hometown)
	fmt.Println("Intro:", bio.Intro)
	fmt.Println("Hobbies:", bio.Hobbies)
	fmt.Println("Tasks:", bio.Tasks)
	fmt.Println("Task:", bio.Task)
}

package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	// Set the file name and path
	viper.SetConfigFile("bio.yaml")

	// Read the YAML file
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading YAML file:", err)
		return
	}

	// Read values from the YAML file
	name := viper.GetString("name")
	age := viper.GetInt("age")
	hometown := viper.GetString("hometown")
	intro := viper.GetString("intro")
	hobbies := viper.GetStringSlice("hobbies")
	tasks := viper.GetStringMapString("tasks")

	// Print the values
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Hometown:", hometown)
	fmt.Println("Intro:", intro)
	fmt.Println("Hobbies:", hobbies)
	fmt.Println("Tasks:", tasks)
}

package main

import (
	"book-management/app/config"
	"book-management/app/domain"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath("../../")      // arahkan ke root project
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	// viper.SetConfigFile(".env")
	// Find and read the config file
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

}

func main()  {

	DB := config.DBConnect()

	DB.AutoMigrate(&domain.Book{})

	fmt.Println("Migration Completed Successfully!")
}


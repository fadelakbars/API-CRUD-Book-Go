package main

import (
	"book-management/app/config"
	"book-management/app/router"
	"log"
	"strconv"

	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func init()  {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
}

func main()  {
	app_port, _ := strconv.Atoi(viper.Get("APP_PORT").(string))
	db := config.DBConnect()
	app := fiber.New()
	router.SetupRoutes(app, db)
	log.Fatal(app.Listen(fmt.Sprintf(":%d", app_port)))
}
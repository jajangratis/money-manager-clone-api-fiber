package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jajangratis/money-manager-clone-api-fiber/routers"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"
	"time"
)

func main() {

	config := viper.New()
	config.SetConfigFile(".env")
	config.AddConfigPath(".")
	// Add more configuration settings here as needed

	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}
	preforkValue := false
	//if os.Getenv("ENVIROMENT") != "local" {
	//	preforkValue = true
	//}
	//// GODOT ENV
	godotenv.Load()
	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		Prefork:      preforkValue,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			ctx.Status(fiber.StatusInternalServerError)
			return ctx.SendString("error " + err.Error())
		},
	})
	//app.Use(logger.New(logger.Config{
	//	Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	//}))
	api := app.Group("/")
	routers.AppRouter(api)
	fmt.Println("listening " + os.Getenv("IP_HOST") + ":" + os.Getenv("IP_PORT"))
	err := app.Listen(os.Getenv("IP_HOST") + ":" + os.Getenv("IP_PORT"))
	if err != nil {
		panic(err)
	}

}

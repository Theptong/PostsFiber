package main

import (
	"fmt"
	"strings"
	"workshop/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

// // Header —
// type Header struct {
// 	Key   string
// 	Value string
// }

// // Response —
// type Response struct {
// 	Headers []Header
// }

// func helloHandler(c *fiber.Ctx) {
// 	var response = &Response{
// 		Headers: []Header{
// 			{
// 				Key:   "Content-Type",
// 				Value: "application/json",
// 			},
// 		},
// 	}

// 	for _, responseHeader := range response.Headers {
// 		c.Set(responseHeader.Key, responseHeader.Value)
// 	}

// 	c.SendString("OKOK")
// }

func main() {
	initConfig()
	app := fiber.New()
	dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.database"))

	db, err := sqlx.Open(viper.GetString("db.driver"), dsn)
	if err != nil {
		panic(err)
	}
	_ = db

	// app.Post("/", func(c *fiber.Ctx) error {
	// 	todoForm := struct {
	// 		Text string `json:"text"`
	// 	}{}

	// 	c.BodyParser(&todoForm) // "{"Text":"do something"}"

	// 	return c.JSON(todoForm)
	// })

	api := app.Group("/api")

	routers.SetCollectionRoutes(api, db)

	app.Get("/error", func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusNotFound, "content not found")
	})
	app.Static("/", "./wwwroot")
	fmt.Println("Server Running on Port", viper.GetInt("app.port"))
	app.Listen(fmt.Sprintf(":%v", viper.GetString("app.port")))
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

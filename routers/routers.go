package routers

import (
	"fmt"
	"log"
	"workshop/repository"
	"workshop/service"
	"workshop/structs"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func SetCollectionRoutes(app fiber.Router, db *sqlx.DB) {
	// ctrls := controller.NewCollectionsRepositoryDB(db)
	customerRepository := repository.NewCustomerRepositoryDB(db)
	customerService := service.NewCustomerService(customerRepository)
	// customerHandler := handler.NewCustomerHandler(customerService)

	// var c *fiber.Ctx
	// Maw := customerHandler.GETABC(c)
	// fmt.Println("Maw::",Maw)
	//
	// maw := customerHandler.GetCustomers(c)
	// fmt.Println("maw:",maw)
	//
	// customerHandler.GetCustomers(c)
	app.Get("/collections", func(c *fiber.Ctx) error {
		list := customerService.GetCollectionService()
		err := c.JSON(list)
		if err != nil {
			panic(err)
		}
		return nil
	})

	app.Get("/collections/:id", func(c *fiber.Ctx) error {
		Params := c.Params("id")
		// var listPost []structs.Posts
		if Params != "" {
			post, err := customerService.GetCustomerServiceById(Params)
			if post != nil {
				c.JSON(post)
			}
			if err != nil {
				fmt.Println("error")
				// err := c.Next()
				return err
			}
		}
		// กรณี ส่งค่าเป็น true & false ต้องไปดึง arr มาแสดง
		if Params == "true" || Params == "false" {
			listPost, err := customerService.GetCustomerServiceByListId(Params)
			if listPost != nil {
				c.JSON(listPost)
			}
			if err != nil {
				fmt.Println("error")
				return err
			}
		} else { //กรณีดึงวันที่ ต้องไปดึง arr มาแสดง
			listPost, err := customerService.GetCustomerServiceByListId(Params)
			if listPost != nil {
				c.JSON(listPost)
			}
			if err != nil {
				fmt.Println("error")
				return err
			}

		}

		return nil
	})
	app.Post("/collections", func(c *fiber.Ctx) error {
		p := new(structs.Posts)
		if err := c.BodyParser(p); err != nil {
			return err
		}
		log.Println(p.Title)   // john
		log.Println(p.Content) // doe
		Post, err := customerService.CreateNewCollection(p.Title, p.Content, p.Published)
		if err != nil {
			panic(err)
		}
		if Post != nil {
			c.JSON(Post)
		}
		return nil
	})

	app.Patch("/collections/:id", func(c *fiber.Ctx) error {
		Params := c.Params("id")
		p := new(structs.Posts)
		if err := c.BodyParser(p); err != nil {
			return err
		}
		// var listPost []structs.Posts
		if Params != "" {
			post, err := customerService.UpdateCollection(Params, p.Title, p.Content, p.Published)
			if post != nil {
				c.JSON(post)
			}
			if err != nil {
				fmt.Println("error")
				// err := c.Next()
				return err
			}
		}
		return nil
	})

	app.Delete("/collections/:id", func(c *fiber.Ctx) error {
		Params := c.Params("id")
		customerService.DeleteCollection(Params)
		// s.customerRepository.DeleteCollection(id)

		c.BodyParser(&Params) // "{"Text":"do something"}"

		return c.JSON(Params)
	})

}

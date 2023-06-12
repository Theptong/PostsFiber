package routers

import (
	"workshop/handler"
	"workshop/repository"
	"workshop/service"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func SetCollectionRoutes(app fiber.Router, db *sqlx.DB) {
	customerRepository := repository.NewCustomerRepositoryDB(db)
	collectionService := service.NewCollectionService(customerRepository)
	handler := handler.NewCustomerHandler(collectionService)

	app.Post("/collections", func(c *fiber.Ctx) error {
		 handler.Create(c, db)
		return nil
	})

	app.Get("/collections/:id", func(c *fiber.Ctx) error {
		Params := c.Params("id")
		handler.GetById(Params, c, db) // byid
		handler.GetByListId(Params, c, db) //bylist id
		return nil
	})

	app.Get("/collections", func(c *fiber.Ctx) error {
		handler.GetCollection(c, db) // List
		limit := c.Query("limit")
		page := c.Query("page")
		if limit != "" || page == "" {
			handler.GetCollectionLimit(page,limit,c,db)
		}
		return nil
	})

	app.Patch("/collections/:id",func(c *fiber.Ctx) error {
		Params := c.Params("id")
		if Params != "" {
			handler.UpdateById(Params,c,db)
		}
		
		return nil
	})
	app.Delete("/collections/:id",func(c *fiber.Ctx) error {
		Params := c.Params("id")
		if Params != "" {
			handler.DeleteById(Params,c,db)
		}
		
		return nil
	})
}


package handler

import (
	"workshop/repository"
	"workshop/service"
	"workshop/structs"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type collectionHandler struct {
	collectionService service.CollectionService
}

func NewCustomerHandler(custSrv service.CollectionService) collectionHandler {
	return collectionHandler{collectionService: custSrv}
}

func (ch collectionHandler) Create(c *fiber.Ctx, db *sqlx.DB) error {
	customerRepository := repository.NewCustomerRepositoryDB(db)
	customerService := service.NewCollectionService(customerRepository)
	p := new(structs.Posts)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	if p.Title != "" {
		posts, err := customerService.CreateNewCollection(p.Title, p.Content, p.Published)
		if err != nil {
			// panic(err)
		}
		return c.JSON(posts)
	}

	required := "error : title is required"

	return c.JSON(required)
}
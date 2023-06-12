package handler

import (
	"strconv"
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

func (ch collectionHandler) Create(c *fiber.Ctx, db *sqlx.DB) (*structs.Posts, error) {
	customerRepository := repository.NewCustomerRepositoryDB(db)
	customerService := service.NewCollectionService(customerRepository)
	p := new(structs.Posts)
	if err := c.BodyParser(p); err != nil {
		return nil, err
	}
	if p.Title != "" {
		posts, err := customerService.CreateNewCollection(p.Title, p.Content, p.Published)
		if err != nil {
			panic(err)
		}
		return nil, c.JSON(posts)
	}

	required := "error : title is required"

	return p, c.JSON(required)
}

func (ch collectionHandler) GetById(id string, c *fiber.Ctx, db *sqlx.DB) (*structs.Posts, error) {
	customerRepository := repository.NewCustomerRepositoryDB(db)
	customerService := service.NewCollectionService(customerRepository)
	posts, err := customerService.GetCollectionServiceById(id)
	if posts != nil {
		if err != nil {
			panic(err)
		}
		return posts, c.JSON(&posts)
	}

	// required := "error : title is required"

	return posts, nil
}

func (ch collectionHandler) GetByListId(id string, c *fiber.Ctx, db *sqlx.DB) ([]structs.Posts, error) {
	customerRepository := repository.NewCustomerRepositoryDB(db)
	customerService := service.NewCollectionService(customerRepository)

	posts, err := customerService.GetCollectionServiceByListId(id)
	if posts != nil {
		if err != nil {
			panic(err)
		}
		return posts, c.JSON(&posts)
	}

	// required := "error : title is required"

	return nil, nil
}

func (ch collectionHandler) GetCollection(c *fiber.Ctx, db *sqlx.DB) (structs.ListPosts, error) {
	var null structs.ListPosts
	customerRepository := repository.NewCustomerRepositoryDB(db)
	customerService := service.NewCollectionService(customerRepository)

	posts, err := customerService.GetCollectionService()
	if &posts != nil {
		if err != nil {
			panic(err)
		}
		return posts, c.JSON(&posts)
	}

	return null, nil
}

func (ch collectionHandler) GetCollectionLimit(page, limit string, c *fiber.Ctx, db *sqlx.DB) ([]structs.Posts, error) {
	var dataList structs.ListPosts
	customerRepository := repository.NewCustomerRepositoryDB(db)
	customerService := service.NewCollectionService(customerRepository)
	Limit, _ := strconv.Atoi(limit)
	Page, _ := strconv.Atoi(page)
	if Page > 0 && Limit > 0 {
		Offset := 0
		if Page >= 0 {
			Offset = (Page - 1) * Limit
		} else {
			Offset = 0
		}
		dataCount, _ := customerService.GetCollectionService()
		dataList.Count = len(dataCount.Posts)

		rows, _ := customerService.GetServiceLimit(Offset, Limit)
		// if err != nil {
		// 	panic(err)
		// 	return nil, err
		// }
		dataList.Posts = append(dataList.Posts, rows...)
		dataList.Limit = Limit
		dataList.Page = Page
		total := (dataList.Count / dataList.Limit)

		remainder := (dataList.Count % dataList.Limit)
		if remainder == 0 {
			dataList.TotalPage = total
		} else {
			dataList.TotalPage = total + 1
		}
		return nil,c.JSON(dataList)
	}

	return nil, nil
}

func (ch collectionHandler) UpdateById(id string, c *fiber.Ctx, db *sqlx.DB) (*structs.Posts, error) {
	customerRepository := repository.NewCustomerRepositoryDB(db)
	customerService := service.NewCollectionService(customerRepository)
	p := new(structs.Posts)
	if err := c.BodyParser(p); err != nil {
		return nil, err
	}
	posts, err := customerService.UpdateCollection(id,p.Title,p.Content,p.Published)
	if posts != nil {
		if err != nil {
			panic(err)
		}
		return posts, c.JSON(&posts)
	}

	// required := "error : title is required"

	return posts, nil
}

func (ch collectionHandler) DeleteById(id string, c *fiber.Ctx, db *sqlx.DB)  error {
	Params := c.Params("id")
	customerRepository := repository.NewCustomerRepositoryDB(db)
	collectionService := service.NewCollectionService(customerRepository)
	collectionService.DeleteCollection(Params)

	Delete := "error : title is required" + id

	return c.JSON(Delete)
}
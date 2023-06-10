package routers

import (
	"fmt"
	"strconv"
	"workshop/repository"
	"workshop/service"
	"workshop/structs"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

// func SetCollectionRoutes(app fiber.Router, db *sqlx.DB) {
// 	// var c *fiber.Ctx
// 	customerRepository := repository.NewCustomerRepositoryDB(db)
// 	collectionService := service.NewCollectionService(customerRepository)
// 	handler := handler.NewCustomerHandler(collectionService)

// 	app.Post("/collections", func(c *fiber.Ctx) error {
// 		data,err := handler.Create(c, db)
// 		if err != nil {
// 			panic(err)
// 		}
// 		if &data != nil {
// 			c.JSON(data)
// 		}
// 		return nil
// 	})

// 	app.Post("/collection", func(c *fiber.Ctx) error {

// 		p := new(structs.Posts)
// 		if err := c.BodyParser(p); err != nil {
// 			return err
// 		}
// 		if p.Title != "" {
// 			posts, err := collectionService.CreateNewCollection(p.Title, p.Content, p.Published)
// 			if err != nil {
// 				panic(err)
// 			}
// 			return c.JSON(posts)
// 		}

// 		required := "error : title is required"

// 		return c.JSON(required)
// 	})

// }

func SetCollectionRoutes(app fiber.Router, db *sqlx.DB) {
	collectionRepository := repository.NewCustomerRepositoryDB(db)
	collectionService := service.NewCollectionService(collectionRepository)

	app.Get("/collections", func(c *fiber.Ctx) error {
		var dataList structs.ListPosts
		limit := c.Query("limit")
		page := c.Query("page")

		post, err := collectionService.GetCollectionService()
		// dataList.Count = len(post.Posts) //จะเอาไว้ตรงนี้ก็ได้ แต่จะย้อนกลับมาดูไม่รู้เรื่อง
		if &post != nil {
			c.JSON(post)
		}
		if err != nil {
			fmt.Println("error")
			// err := c.Next()
			return err
		}
		//
		if limit != "" || page == "" {
			Limit, _ := strconv.Atoi(limit)
			Page, _ := strconv.Atoi(page)
			Offset := 0
			if Page >= 0 {
				Offset = (Page - 1) * Limit
			} else {
				Offset = 0
			}
			if Limit > 0 {

				dataCount, _ := collectionService.GetCollectionService() // จำนวนทั้งหมด
				dataList.Count = len(dataCount.Posts)                    //ดึงจากก้อน
				rows, err := collectionService.GetServiceLimit(Offset, Limit)

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
				if &dataList != nil {
					c.JSON(dataList)
				}
				if err != nil {
					fmt.Println("error")
					return err
				}
			}
		}
		return nil
	})

	app.Get("/collections/:id", func(c *fiber.Ctx) error {
		Params := c.Params("id")
		// var listPost []structs.Posts
		if Params != "" {
			post, err := collectionService.GetCollectionServiceById(Params)
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
			listPost, err := collectionService.GetCollectionServiceByListId(Params)
			if listPost != nil {
				c.JSON(listPost)
			}
			if err != nil {
				fmt.Println("error")
				return err
			}
		} else { //กรณีดึงวันที่ ต้องไปดึง arr มาแสดง
			listPost, err := collectionService.GetCollectionServiceByListId(Params)
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
		if p.Title != "" {
			posts, err := collectionService.CreateNewCollection(p.Title, p.Content, p.Published)
			if err != nil {
				panic(err)
			}
			return c.JSON(posts)
		}

		required := "error : title is required"

		return c.JSON(required)
	})

	app.Patch("/collections/:id", func(c *fiber.Ctx) error {
		Params := c.Params("id")
		p := new(structs.Posts)
		if err := c.BodyParser(p); err != nil {
			return err
		}
		if p.Title != "" {
			collectionRepository := repository.NewCustomerRepositoryDB(db)
			collectionService := service.NewCollectionService(collectionRepository)
			post, _ := collectionService.UpdateCollection(Params, p.Title, p.Content, p.Published)
			// s.customerRepository.DeleteCollection(id)

			c.BodyParser(&post) // "{"Text":"do something"}"

			return c.JSON(post)
		}

		required := "error : title is required"
		return c.JSON(required)
	})

	app.Delete("/collections/:id", func(c *fiber.Ctx) error {
		Params := c.Params("id")
		collectionService.DeleteCollection(Params)
		// s.customerRepository.DeleteCollection(id)

		c.BodyParser(&Params) // "{"Text":"do something"}"

		return c.JSON(Params)
	})

}

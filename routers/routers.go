package routers

import (
	"workshop/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func SetCollectionRoutes(app fiber.Router, db *sqlx.DB) {
	// ctrls := v1.DBController{Database: db}
	ctrls := controller.NewCollectionsRepositoryDB(db)

	app.Get("/collections", func(c *fiber.Ctx) error {
		err := c.JSON(ctrls.GetCollection())
		if err != nil {
			panic(err)
		}
		return nil
	})
	// router.GET("collections", ctrls.GetCollection) // GET
	// router.GET("collections/:id", ctrls.GetCollectionById)   // GET BY ID
	// router.POST("collections", ctrls.CreateCollection)       // POST
	// router.PATCH("collections/:id", ctrls.UpdateCollection)  // PATCH
	// router.DELETE("collections/:id", ctrls.DeleteCollection) // DELETE
}

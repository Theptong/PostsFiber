package service_test

import (
	"testing"
	"time"
	"workshop/repository"
	"workshop/service"

	"github.com/magiconair/properties/assert"
)

func TestGetCollectionById(t *testing.T) {
	// dateString := "2023-06-14"
	// date, _ := time.Parse("2006-01-02", dateString)
	type testPosts struct {
		Id        string
		Title     string
		Content   string
		Published bool
		ViewCount int
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	cases := []testPosts{
		{Id: "123", Title: "title123", Content: "Content123", Published: true, ViewCount: 0, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Id: "456", Title: "title456", Content: "Content456", Published: true, ViewCount: 0, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		// {Id: "", Title: "title789", Content: "Content789", Published: true, ViewCount: 0, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		// {Id: "", Title: "", Content: "", Published: false, ViewCount: 0, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	for _, c := range cases {

		t.Run(c.Id, func(t *testing.T) {
			//Arrage
			promoRepo := repository.NewPromotionRepositoryMock()
			promoRepo.On("CheckCollectionById").Return(
				repository.MockPosts{
					Id:        "123",
					Title:     "title123",
					Content:   "Content123",
					Published: true,
					ViewCount: 0,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
			data, _ := promoRepo.CheckCollectionById()
			promoService := service.NewPromotionService(promoRepo)
			// fmt.Println("promoService::",promoService)
			// // Act
			check, _ := promoService.GetCollectionById(data.Id)
			// // Assert
			assert.Equal(t, c.Id, check)
		})

	}

}

func TestGetCollection(t *testing.T) {
	dateString := "2023-06-14"
	date, _ := time.Parse("2006-01-02", dateString)
	type testPosts struct {
		Id        string
		Title     string
		Content   string
		Published bool
		ViewCount int
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	cases := []repository.MockPosts{
		{Id: "123", Title: "title123", Content: "Content123", Published: true, ViewCount: 0, CreatedAt: date, UpdatedAt: date},
		{Id: "456", Title: "456", Content: "456", Published: true, ViewCount: 0, CreatedAt: date, UpdatedAt: date},
		// {Id: "", Title: "title789", Content: "Content789", Published: true, ViewCount: 0, CreatedAt: date, UpdatedAt: date},
		// {Id: "", Title: "", Content: "", Published: false, ViewCount: 0, CreatedAt: date, UpdatedAt: date},
	}

	for _, c := range cases {
		t.Run(c.Id, func(t *testing.T) {
			promoRepo := repository.NewPromotionRepositoryMock()
			promoRepo.On("GetCollection").Return([]repository.MockPosts{
				{Id: "123", Title: "title123", Content: "Content123", Published: true, ViewCount: 0, CreatedAt: date, UpdatedAt: date},
				{Id: "456", Title: "456", Content: "456", Published: true, ViewCount: 0, CreatedAt: date, UpdatedAt: date},
			}, nil)
			data, _ := promoRepo.GetCollection(cases)

			promoService := service.NewPromotionService(promoRepo)

			// Act
			check, _ := promoService.GetCollection(data)
			// }
			// fmt.Println("check::", check)

			// Assert
			assert.Equal(t, cases, check)

		})
	}

}

// func TestUpdateCollection(t *testing.T) {

// 	dateString := "2023-06-14"
// 	date, _ := time.Parse("2006-01-02", dateString)
// 	type testPosts struct {
// 		Id        string
// 		Title     string
// 		Content   string
// 		Published bool
// 		ViewCount int
// 		CreatedAt time.Time
// 		UpdatedAt time.Time
// 	}

// 	cases := repository.MockPosts{
// 		Id: "123",
// 		Title: "title123",
// 		 Content: "Content123",
// 		Published: true,
// 		ViewCount: 0,
// 		CreatedAt: date,
// 		UpdatedAt: date,
// 	}

// 		t.Run(cases.Id, func(t *testing.T) {
// 			promoRepo := repository.NewPromotionRepositoryMock()
// 			promoRepo.On("UpdateCollection").Return(repository.MockPosts{
// 				Id: "123",
// 				Title: "title123",
// 				Content: "Content123",
// 				Published: true,
// 				ViewCount: 0,
// 				CreatedAt: date,
// 				UpdatedAt: date,
// 			})
// 			// fmt.Println("promoRepo::",promoRepo)
// 			promoService := service.NewPromotionService(promoRepo)
// 			// fmt.Println("promoService:",promoService)
// 			data, _ := promoService.UpdateCollection()
// 			// if
// 			// // Act
// 			// check, _ := promoService.GetCollection(c.Id)
// 			// // // Assert
// 			assert.Equal(t, cases, data)

// 		})

// }

// func TestDeleteCollection(t *testing.T) {

// 	dateString := "2023-06-14"
// 	date, _ := time.Parse("2006-01-02", dateString)
// 	type testPosts struct {
// 		Id        string
// 		Title     string
// 		Content   string
// 		Published bool
// 		ViewCount int
// 		CreatedAt time.Time
// 		UpdatedAt time.Time
// 	}

// 	cases := repository.MockPosts{
// 		Id: "123",
// 		Title: "title123",
// 		 Content: "Content123",
// 		Published: true,
// 		ViewCount: 0,
// 		CreatedAt: date,
// 		UpdatedAt: date,
// 	}

// 		t.Run(cases.Id, func(t *testing.T) {
// 			promoRepo := repository.NewPromotionRepositoryMock()
// 			promoRepo.On("DeleteCollection").Return(repository.MockPosts{
// 				Id: "123",
// 				Title: "title123",
// 				Content: "Content123",
// 				Published: true,
// 				ViewCount: 0,
// 				CreatedAt: date,
// 				UpdatedAt: date,
// 			})
// 			// fmt.Println("promoRepo::",promoRepo)
// 			promoService := service.NewPromotionService(promoRepo)
// 			// fmt.Println("promoService:",promoService)
// 			data, _ := promoService.DeleteCollection()
// 			// if
// 			// // Act
// 			// check, _ := promoService.GetCollection(c.Id)
// 			// // // Assert
// 			assert.Equal(t, cases, data)

// 		})

// }

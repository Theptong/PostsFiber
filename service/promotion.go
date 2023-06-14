package service

import (
	"errors"
	"time"
	"workshop/repository"
)

type PromotionService interface {
	GetCollectionById(id string) (string, error)
}

type promotionService struct {
	promoRepo repository.PromotionRepository
}

func NewPromotionService(promoRepo repository.PromotionRepository) promotionService {
	return promotionService{promoRepo: promoRepo}
}

func (p promotionService) GetCollectionById(id string) (string, error) {

	if id == "" {
		return "", errors.New("don't found id")
	}
	if id == "123" {
		return id, nil
	}
	if id == "456" {
		return id, nil
	}

	return id, nil
}

func (p promotionService) GetCollection(id string) ([]repository.MockPosts, error) {
	dateString := "2023-06-14"
	date, _ := time.Parse("2006-01-02", dateString)
	mockData := []repository.MockPosts{}
	cases := []repository.MockPosts{
		{Id: "123", Title: "title123", Content: "Content123", Published: true, ViewCount: 0, CreatedAt: date, UpdatedAt: date},
		// {Id: "456", Title: "456", Content: "456", Published: true, ViewCount: 0, CreatedAt: date, UpdatedAt: date},
		// {Id: "789", Title: "789", Content: "789", Published: true, ViewCount: 0, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		// {Id: "", Title: "", Content: "", Published: false, ViewCount: 0, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	if id != "" {
		for _, obj := range cases {
			if obj.Id == id {
				mockData = append(mockData, obj)
				return mockData, nil
			}
		}
	}

	return cases, nil
}

func (p promotionService) CreateCollection() (repository.MockPosts, error) {
	var data repository.MockPosts
	dateString := "2023-06-14"
	date, _ := time.Parse("2006-01-02", dateString)
	Id := "123"
	Title := "title123"
	Content := "Content123"
	Published := true
	ViewCount := 0
	CreatedAt := date
	UpdatedAt := date

	data.Id = Id
	data.Title = Title
	data.Content = Content
	data.Published = Published
	data.ViewCount = ViewCount
	data.CreatedAt = CreatedAt
	data.UpdatedAt = UpdatedAt

	return data, nil
}

func (p promotionService) UpdateCollection() (repository.MockPosts, error) {
	var data repository.MockPosts
	dateString := "2023-06-14"
	date, _ := time.Parse("2006-01-02", dateString)
	Id := "123"
	Title := "title123"
	Content := "Content123"
	Published := true
	ViewCount := 0
	CreatedAt := date
	UpdatedAt := date

	data.Id = Id
	data.Title = Title
	data.Content = Content
	data.Published = Published
	data.ViewCount = ViewCount
	data.CreatedAt = CreatedAt
	data.UpdatedAt = UpdatedAt

	return data, nil
}
func (p promotionService) DeleteCollection() (repository.MockPosts, error) {
	var data repository.MockPosts
	dateString := "2023-06-14"
	date, _ := time.Parse("2006-01-02", dateString)
	Id := "123"
	Title := "title123"
	Content := "Content123"
	Published := true
	ViewCount := 0
	CreatedAt := date
	UpdatedAt := date

	data.Id = Id
	data.Title = Title
	data.Content = Content
	data.Published = Published
	data.ViewCount = ViewCount
	data.CreatedAt = CreatedAt
	data.UpdatedAt = UpdatedAt

	return data, nil
}
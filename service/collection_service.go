package service

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
	"workshop/repository"
	"workshop/structs"

	"github.com/google/uuid"
)

// ตัวเชื่อม
type collectionService struct {
	collectionRepository repository.CollectionRepository
}

func NewCollectionService(custRepo repository.CollectionRepository) collectionService {
	return collectionService{collectionRepository: custRepo}
}

// GET ALL
func (s collectionService) GetCollection() ([]structs.Posts, error) {
	PostsRes := []structs.Posts{}
	postsDB, err := s.collectionRepository.GetAll()
	if err != nil {
		log.Panicln(err)
		// panic(err)
		return nil, err
	}

	var Posts structs.Posts
	if len(postsDB) > 0 {
		for _, obj := range postsDB {
			id := uuid.MustParse(fmt.Sprint(&obj.Id)) // uuid ไม่สามารถยัดค่าไปตรงๆได้ ต้องแปลงเป็น สตริงแล้วแปลงกลับเป็น uuid
			Posts.Id = &id
			Posts.Title = obj.Title
			Posts.Content = obj.Content
			Posts.Published = obj.Published
			Posts.ViewCount = obj.ViewCount
			Posts.CreatedAt = &obj.CreatedAt
			Posts.UpdatedAt = &obj.UpdatedAt
			PostsRes = append(PostsRes, Posts)
		}
	}

	return PostsRes, nil
}

// GET
func (s collectionService) GetCollectionService() (structs.ListPosts, error) {
	var dataList structs.ListPosts
	// app := fiber.New()
	rows, _ := s.GetCollection()
	if len(rows) > 0 {

		dataList.Posts = rows
		dataList.Count = len(rows)
		dataList.Limit = len(rows)
		total := (dataList.Count / dataList.Limit)
		dataList.Page = total
		remainder := (dataList.Count % dataList.Limit)
		if remainder == 0 {
			dataList.TotalPage = total
		} else {
			dataList.TotalPage = total + 1
		}
		return dataList, nil
	} else {
		dataList := structs.ListPosts{}
		return dataList, nil
	}
}

// GET ID
func (s collectionService) GetCollectionServiceById(id string) (*structs.Posts, error) {
	s.UpdateCollectionByViewCount(id)

	var Posts structs.Posts
	postsDBByID, err := s.collectionRepository.GetById(id)
	postsDBTitle, _ := s.collectionRepository.GetByTitle(id)
	postsDBContent, _ := s.collectionRepository.GetByContent(id)
	// boolValue, _ := strconv.ParseBool(id)
	// postsDBPublished,_ := s.collectionRepository.GetByPublished(boolValue)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("ID not found")
		}
	}

	if postsDBByID != nil || postsDBTitle != nil || postsDBContent != nil {

		if postsDBByID != nil {
			Posts.Id = &postsDBByID.Id
			Posts.Title = postsDBByID.Title
			Posts.Content = postsDBByID.Content
			Posts.Published = postsDBByID.Published
			Posts.ViewCount = postsDBByID.ViewCount
			Posts.CreatedAt = &postsDBByID.CreatedAt
			Posts.UpdatedAt = &postsDBByID.UpdatedAt
		} else if postsDBTitle != nil {
			Posts.Id = &postsDBTitle.Id
			Posts.Title = postsDBTitle.Title
			Posts.Content = postsDBTitle.Content
			Posts.Published = postsDBTitle.Published
			Posts.ViewCount = postsDBTitle.ViewCount
			Posts.CreatedAt = &postsDBTitle.CreatedAt
			Posts.UpdatedAt = &postsDBTitle.UpdatedAt
		} else if postsDBContent != nil {
			Posts.Id = &postsDBContent.Id
			Posts.Title = postsDBContent.Title
			Posts.Content = postsDBContent.Content
			Posts.Published = postsDBContent.Published
			Posts.ViewCount = postsDBContent.ViewCount
			Posts.CreatedAt = &postsDBContent.CreatedAt
			Posts.UpdatedAt = &postsDBContent.UpdatedAt
		}
		return &Posts, nil
	} else {
		//กรณีใส่ ไอดีผิดส่ง nilออกไป
		return &Posts, nil
	}

}

// Update view_count
func (s collectionService) UpdateCollectionByViewCount(id string) (structs.Posts, error) {
	var dataPosts structs.Posts
	ViewCount := 0
	data, _ := s.collectionRepository.GetAll()
	// if len(data) > 0 {
	for _, obj := range data {
		if id == fmt.Sprint(obj.Id) {
			if obj.Published == true {
				ViewCount = obj.ViewCount + 1
				s.collectionRepository.ViewCountCollection(ViewCount, id)
				dataPosts.Id = &obj.Id
				dataPosts.Title = obj.Title
				dataPosts.Content = obj.Content
				dataPosts.Published = obj.Published
				dataPosts.ViewCount = obj.ViewCount
				dataPosts.CreatedAt = &obj.CreatedAt
				dataPosts.UpdatedAt = &obj.UpdatedAt
			}
		}
	}

	return dataPosts, nil
}

// GET BY ID LIST
func (s collectionService) GetCollectionServiceByListId(id string) ([]structs.Posts, error) {

	var listPosts []structs.Posts
	var PostsInfo structs.Posts

	if id == "true" || id == "false" {
		postsDBPublished, _ := s.collectionRepository.GetByPublished(id)
		if postsDBPublished != nil {

			for _, obj := range postsDBPublished {
				PostsInfo.Id = &obj.Id
				PostsInfo.Title = obj.Title
				PostsInfo.Content = obj.Content
				PostsInfo.Published = obj.Published
				PostsInfo.ViewCount = obj.ViewCount
				PostsInfo.CreatedAt = &obj.CreatedAt
				PostsInfo.UpdatedAt = &obj.UpdatedAt

				listPosts = append(listPosts, PostsInfo)
			}
		}
	} else {
		timeNow := time.Now().AddDate(0, 0, 1)
		CreatedAt := strings.Split(fmt.Sprint(timeNow), "T")
		Date := strings.Split(fmt.Sprint(CreatedAt[0]), " ")
		today := fmt.Sprint(Date[0])
		postsDBDate, _ := s.collectionRepository.GetByDate(id, today)
		if postsDBDate != nil {
			for _, obj := range postsDBDate {
				PostsInfo.Id = &obj.Id
				PostsInfo.Title = obj.Title
				PostsInfo.Content = obj.Content
				PostsInfo.Published = obj.Published
				PostsInfo.ViewCount = obj.ViewCount
				PostsInfo.CreatedAt = &obj.CreatedAt
				PostsInfo.UpdatedAt = &obj.UpdatedAt
				listPosts = append(listPosts, PostsInfo)
			}
		}
	}
	return listPosts, nil
}

// Create
func (s collectionService) CreateNewCollection(title, content string, published bool) (*structs.Posts, error) {

	var Posts structs.Posts
	s.collectionRepository.CreateNewCollection(title, content, published) //ใส่ข้อมูลเข้า
	database, _ := s.collectionRepository.GetAll()                        //นำมาแสดงออก
	Obj := database[len(database)-1]
	if &Obj != nil {
		Posts.Id = &Obj.Id
		Posts.Title = Obj.Title
		Posts.Content = Obj.Content
		Posts.Published = Obj.Published
		Posts.ViewCount = Obj.ViewCount
		Posts.CreatedAt = &Obj.CreatedAt
		Posts.UpdatedAt = &Obj.UpdatedAt
	}
	return &Posts, nil

}

// Update
func (s collectionService) UpdateCollection(id, title, content string, published bool) (*structs.Posts, error) {

	var Posts structs.Posts
	s.collectionRepository.UpdateCollection(id, title, content, published) //ใส่ข้อมูลเข้า
	// database, _ := s.GetCollectionServiceById(id)
	database, _ := s.collectionRepository.GetById(id)                      //นำมาแสดงออก
	// // fmt.Println("Create::", Create)
	if &database != nil {
		Posts.Id = &database.Id
		Posts.Title = database.Title
		Posts.Content = database.Content
		Posts.Published = database.Published
		Posts.ViewCount = database.ViewCount
		Posts.CreatedAt = &database.CreatedAt
		Posts.UpdatedAt = &database.UpdatedAt

	}
	return &Posts, nil

}

// Delete
func (s collectionService) DeleteCollection(id string) error {
	s.collectionRepository.DeleteCollection(id) //ใส่ข้อมูลเข้า
	
	return nil
}

// GET
func (s collectionService) GetCollectionServiceLimit(page, limit int) structs.ListPosts {
	var dataList structs.ListPosts
	rows, _ := s.GetServiceLimit(page, limit)

	// rows, _ := s.GetCustomerService()
	if len(rows) > 0 {

		dataList.Posts = rows
		dataList.Count = len(rows)
		dataList.Limit = len(rows)
		total := (dataList.Count / dataList.Limit)
		dataList.Page = total
		remainder := (dataList.Count % dataList.Limit)
		if remainder == 0 {
			dataList.TotalPage = total
		} else {
			dataList.TotalPage = total + 1
		}
		return dataList
	} else {
		dataList := structs.ListPosts{}
		return dataList
	}
}

// GET Limit
func (s collectionService) GetServiceLimit(page, limit int) ([]structs.Posts, error) {
	PostsRes := []structs.Posts{}
	postsDB, err := s.collectionRepository.LimitCollection(page, limit)
	if err != nil {
		log.Panicln(err)
		panic(err)
		// return nil, err
	}

	var Posts structs.Posts
	if len(postsDB) > 0 {
		for _, obj := range postsDB {
			id := uuid.MustParse(fmt.Sprint(&obj.Id)) // uuid ไม่สามารถยัดค่าไปตรงๆได้ ต้องแปลงเป็น สตริงแล้วแปลงกลับเป็น uuid
			Posts.Id = &id
			Posts.Title = obj.Title
			Posts.Content = obj.Content
			Posts.Published = obj.Published
			Posts.ViewCount = obj.ViewCount
			Posts.CreatedAt = &obj.CreatedAt
			Posts.UpdatedAt = &obj.UpdatedAt
			PostsRes = append(PostsRes, Posts)
		}
	}

	return PostsRes, nil
}

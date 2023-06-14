package service

// import (
// 	"workshop/repository"

// 	"github.com/stretchr/testify/mock"
// )

// type promotionServiceMock struct {
// 	mock.Mock
// }

// func NewPromotionServiceMock() *promotionServiceMock {
// 	return &promotionServiceMock{}
// }

// func (m *promotionServiceMock) GetPromotion() (repository.MockPosts, error) {
// 	args := m.Called()

// 	return args.Get(0).(repository.MockPosts), args.Error(1)
// }

// func (m *promotionServiceMock) GetCollection(id string) ([]repository.MockPosts, error) {
// 	args := m.Called(id)

// 	var r0 []repository.MockPosts	
// 	v0 := args.Get(0)

// 	if v0 != nil {
// 		r0 = v0.([]repository.MockPosts)
// 	}

// 	var r1 error
// 	v1 := args.Get(1)
// 	if v1 != nil {
// 		r1 = v1.(error)
// 	}

// 	return r0,r1
// }

package repository

import (
	"github.com/stretchr/testify/mock"
)

type promotionRepositoryMock struct {
	mock.Mock
}

func NewPromotionRepositoryMock() *promotionRepositoryMock {
	return &promotionRepositoryMock{}
}

func (m *promotionRepositoryMock) GetPromotion() (MockPosts, error) {
	args := m.Called()

	return args.Get(0).(MockPosts), args.Error(1)
}

// func (m *promotionRepositoryMock) GetCollection() ([]MockPosts, error) {
// 	fmt.Println("args := m.Called()")
// 	args := m.Called()

// 	var r0 []MockPosts
// 	v0 := args.Get(0)

// 	if v0 != nil {
// 		r0 = v0.([]MockPosts)
// 	}

// 	var r1 error
// 	v1 := args.Get(1)
// 	if v1 != nil {
// 		r1 = v1.(error)
// 	}

//		return r0,r1
//	}
func (m *promotionRepositoryMock) GetCollection() ([]MockPosts, error) {
	args := m.Called()

	return args.Get(0).([]MockPosts), args.Error(1)
}

func (m *promotionRepositoryMock) CreateCollection() (MockPosts, error) {
	args := m.Called()
	return args.Get(0).(MockPosts), args.Error(1)
}

func (m *promotionRepositoryMock) UpdateCollection() (MockPosts, error) {
	args := m.Called()
	return args.Get(0).(MockPosts), args.Error(1)
}
func (m *promotionRepositoryMock) DeleteCollection() (MockPosts, error) {
	args := m.Called()
	return args.Get(0).(MockPosts), args.Error(1)
}

//	func (m *promotionRepositoryMock) CreateCollection() (title,content string,published bool, err error) {
//		args := m.Called()
//		return args.String(0),args.String(1),args.Bool(2),args.Error(3)
//	}
// func (m *promotionRepositoryMock) UpdateCollection(id string) (title, content string, published bool, err error) {
// 	args := m.Called(id)

// 	return args.String(0), args.String(1), args.Bool(2), args.Error(3)
// }
// func (m *promotionRepositoryMock) DeleteCollection(id string) ([]MockPosts, error) {
// 	args := m.Called(id)

// 	return args.Get(0).([]MockPosts), args.Error(1)
// }

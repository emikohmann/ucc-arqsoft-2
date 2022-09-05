package services

import (
	"github.com/emikohmann/ucc-arqsoft-2/ej-books/dtos"
	e "github.com/emikohmann/ucc-arqsoft-2/ej-books/utils/errors"
)

type ServiceMock struct{}

func NewServiceMock() ServiceMock {
	return ServiceMock{}
}

func (ServiceMock) Get(id string) (dtos.BookDTO, e.ApiError) {
	//TODO implement me
	panic("implement me")
}

func (ServiceMock) Insert(book dtos.BookDTO) (dtos.BookDTO, e.ApiError) {
	//TODO implement me
	panic("implement me")
}

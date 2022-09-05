package services

import (
	"github.com/emikohmann/ucc-arqsoft-2/ej-books/dtos"
	e "github.com/emikohmann/ucc-arqsoft-2/ej-books/utils/errors"
)

type ServiceMock struct{}

func NewServiceMock() ServiceMock {
	return ServiceMock{}
}

func (ServiceMock) Get(id string) (dtos.BookDto, e.ApiError) {
	//TODO implement me
	panic("implement me")
}

func (ServiceMock) Insert(book dtos.BookDto) (dtos.BookDto, e.ApiError) {
	//TODO implement me
	panic("implement me")
}

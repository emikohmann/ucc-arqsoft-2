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
	return dtos.BookDTO{
		Id:   "12345",
		Name: "Mocked book",
	}, nil
}

func (ServiceMock) Insert(book dtos.BookDTO) (dtos.BookDTO, e.ApiError) {
	return dtos.BookDTO{
		Id:   "12345",
		Name: book.Name,
	}, nil
}

package services

import (
	"github.com/emikohmann/ucc-arqsoft-2/ej-books/dtos"
	e "github.com/emikohmann/ucc-arqsoft-2/ej-books/utils/errors"
)

type ServiceImpl struct{}

func NewServiceImpl() ServiceImpl {
	return ServiceImpl{}
}

func (ServiceImpl) Get(id string) (dtos.BookDto, e.ApiError) {
	//TODO implement me
	panic("implement me")
}

func (ServiceImpl) Insert(book dtos.BookDto) (dtos.BookDto, e.ApiError) {
	//TODO implement me
	panic("implement me")
}

package repositories

import (
	"github.com/emikohmann/ucc-arqsoft-2/ej-books/dtos"
	"github.com/emikohmann/ucc-arqsoft-2/ej-books/utils/errors"
)

type RepositoryMongo struct{}

func NewRepositoryMongo() RepositoryMongo {
	return RepositoryMongo{}
}

func (RepositoryMongo) Get(id string) (dtos.BookDto, errors.ApiError) {
	//TODO implement me
	panic("implement me")
}

func (RepositoryMongo) Insert(dto dtos.BookDto) errors.ApiError {
	//TODO implement me
	panic("implement me")
}

func (RepositoryMongo) Update(dto dtos.BookDto) errors.ApiError {
	//TODO implement me
	panic("implement me")
}

func (RepositoryMongo) Delete(id string) errors.ApiError {
	//TODO implement me
	panic("implement me")
}

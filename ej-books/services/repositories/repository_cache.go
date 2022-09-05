package repositories

import (
	"github.com/emikohmann/ucc-arqsoft-2/ej-books/dtos"
	"github.com/emikohmann/ucc-arqsoft-2/ej-books/utils/errors"
)

type RepositoryCache struct{}

func NewRepositoryCache() RepositoryCache {
	return RepositoryCache{}
}

func (RepositoryCache) Get(id string) (dtos.BookDto, errors.ApiError) {
	//TODO implement me
	panic("implement me")
}

func (RepositoryCache) Insert(dto dtos.BookDto) errors.ApiError {
	//TODO implement me
	panic("implement me")
}

func (RepositoryCache) Update(dto dtos.BookDto) errors.ApiError {
	//TODO implement me
	panic("implement me")
}

func (RepositoryCache) Delete(id string) errors.ApiError {
	//TODO implement me
	panic("implement me")
}

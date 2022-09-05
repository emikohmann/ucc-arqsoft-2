package repositories

import (
	"github.com/emikohmann/ucc-arqsoft-2/ej-books/dtos"
	"github.com/emikohmann/ucc-arqsoft-2/ej-books/utils/errors"
)

type RepositoryCache struct{}

func NewRepositoryCache() *RepositoryCache {
	return &RepositoryCache{}
}

func (repo *RepositoryCache) Get(id string) (dtos.BookDTO, errors.ApiError) {
	//TODO implement me
	panic("implement me")
}

func (repo *RepositoryCache) Insert(dto dtos.BookDTO) errors.ApiError {
	//TODO implement me
	panic("implement me")
}

func (repo *RepositoryCache) Update(dto dtos.BookDTO) errors.ApiError {
	//TODO implement me
	panic("implement me")
}

func (repo *RepositoryCache) Delete(id string) errors.ApiError {
	//TODO implement me
	panic("implement me")
}

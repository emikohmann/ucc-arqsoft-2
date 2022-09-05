package services

import (
	"github.com/emikohmann/ucc-arqsoft-2/ej-books/dtos"
	"github.com/emikohmann/ucc-arqsoft-2/ej-books/services/repositories"
	e "github.com/emikohmann/ucc-arqsoft-2/ej-books/utils/errors"
)

type ServiceImpl struct {
	cache     repositories.RepositoryCache
	memcached repositories.RepositoryMemcached
	mongo     repositories.RepositoryMongo
}

func NewServiceImpl(
	cache repositories.RepositoryCache,
	memcached repositories.RepositoryMemcached,
	mongo repositories.RepositoryMongo,
) ServiceImpl {
	return ServiceImpl{
		cache:     cache,
		memcached: memcached,
		mongo:     mongo,
	}
}

func (ServiceImpl) Get(id string) (dtos.BookDto, e.ApiError) {
	//TODO implement me
	panic("implement me")
}

func (ServiceImpl) Insert(book dtos.BookDto) (dtos.BookDto, e.ApiError) {
	//TODO implement me
	panic("implement me")
}

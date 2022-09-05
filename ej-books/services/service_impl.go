package services

import (
	"fmt"
	"github.com/emikohmann/ucc-arqsoft-2/ej-books/dtos"
	"github.com/emikohmann/ucc-arqsoft-2/ej-books/services/repositories"
	e "github.com/emikohmann/ucc-arqsoft-2/ej-books/utils/errors"
	"net/http"
)

type ServiceImpl struct {
	cache     *repositories.RepositoryCache
	memcached *repositories.RepositoryMemcached
	mongo     *repositories.RepositoryMongo
}

func NewServiceImpl(
	cache *repositories.RepositoryCache,
	memcached *repositories.RepositoryMemcached,
	mongo *repositories.RepositoryMongo,
) *ServiceImpl {
	return &ServiceImpl{
		cache:     cache,
		memcached: memcached,
		mongo:     mongo,
	}
}

func (serv *ServiceImpl) Get(id string) (dtos.BookDTO, e.ApiError) {
	var book dtos.BookDTO
	var apiErr e.ApiError
	var source string

	book, apiErr = serv.cache.Get(id)
	if apiErr != nil {
		if apiErr.Status() != http.StatusNotFound {
			return dtos.BookDTO{}, apiErr
		}
		book, apiErr = serv.memcached.Get(id)
		if apiErr != nil {
			if apiErr.Status() != http.StatusNotFound {
				return dtos.BookDTO{}, apiErr
			}
			book, apiErr = serv.mongo.Get(id)
			if apiErr != nil {
				return dtos.BookDTO{}, apiErr
			} else {
				source = "mongo"
			}
		} else {
			source = "memcached"
		}
	} else {
		source = "cache"
	}
	fmt.Println(fmt.Sprintf("Obtained book from %s!", source))
	return book, nil
}

func (serv *ServiceImpl) Insert(book dtos.BookDTO) (dtos.BookDTO, e.ApiError) {
	//TODO implement me
	panic("implement me")
}

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

	// try to find it in cache
	book, apiErr = serv.cache.Get(id)
	if apiErr != nil {
		if apiErr.Status() != http.StatusNotFound {
			return dtos.BookDTO{}, apiErr
		}
		defer func() {
			if apiErr != nil {
				if apiErr := serv.cache.Insert(book); apiErr != nil {
					fmt.Println(fmt.Sprintf("Error trying to save book in cache %v", apiErr))
				}
			}
		}()
		// try to find it in memcached
		book, apiErr = serv.memcached.Get(id)
		if apiErr != nil {
			if apiErr.Status() != http.StatusNotFound {
				return dtos.BookDTO{}, apiErr
			}
			defer func() {
				if apiErr != nil {
					if apiErr := serv.memcached.Insert(book); apiErr != nil {
						fmt.Println(fmt.Sprintf("Error trying to save book in memcached %v", apiErr))
					}
				}
			}()
			// try to find it in mongo
			book, apiErr = serv.mongo.Get(id)
			if apiErr != nil {
				if apiErr.Status() != http.StatusNotFound {
					return dtos.BookDTO{}, apiErr
				} else {
					fmt.Println(fmt.Sprintf("Not found book %s in any datasource", id))
					apiErr = e.NewNotFoundApiError(fmt.Sprintf("book %s not found", id))
					return dtos.BookDTO{}, apiErr
				}
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
	if apiErr := serv.mongo.Insert(book); apiErr != nil {
		fmt.Println(fmt.Sprintf("Error inserting book in mongo: %v", apiErr))
		return dtos.BookDTO{}, apiErr
	}

	if apiErr := serv.memcached.Insert(book); apiErr != nil {
		fmt.Println(fmt.Sprintf("Error inserting book in memcached: %v", apiErr))
		return dtos.BookDTO{}, nil
	}

	if apiErr := serv.cache.Insert(book); apiErr != nil {
		fmt.Println(fmt.Sprintf("Error inserting book in cache: %v", apiErr))
		return dtos.BookDTO{}, nil
	}

	return book, nil
}

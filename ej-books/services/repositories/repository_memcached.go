package repositories

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/emikohmann/ucc-arqsoft-2/ej-books/dtos"
	"github.com/emikohmann/ucc-arqsoft-2/ej-books/utils/errors"
	json "github.com/json-iterator/go"
)

type RepositoryMemcached struct {
	Client *memcache.Client
}

func NewRepositoryMemcached(host string, port int) *RepositoryMemcached {
	client := memcache.New(fmt.Sprintf("%s:%d", host, port))
	return &RepositoryMemcached{
		Client: client,
	}
}

func (repo *RepositoryMemcached) Get(id string) (dtos.BookDto, errors.ApiError) {
	item, err := repo.Client.Get(id)
	if err != nil {
		return dtos.BookDto{}, errors.NewInternalServerApiError(err.Error(), err)
	}

	var bookDTO dtos.BookDto
	if err := json.Unmarshal(item.Value, &bookDTO); err != nil {
		return dtos.BookDto{}, errors.NewInternalServerApiError(err.Error(), err)
	}

	return bookDTO, nil
}

func (repo *RepositoryMemcached) Insert(book dtos.BookDto) errors.ApiError {
	bytes, err := json.Marshal(book)
	if err != nil {
		return errors.NewBadRequestApiError(err.Error())
	}

	if err := repo.Client.Set(&memcache.Item{
		Key:   book.Id,
		Value: bytes,
	}); err != nil {
		return errors.NewInternalServerApiError(err.Error(), err)
	}

	return nil
}

func (repo *RepositoryMemcached) Update(book dtos.BookDto) errors.ApiError {
	//TODO implement me
	panic("implement me")
}

func (repo *RepositoryMemcached) Delete(id string) errors.ApiError {
	//TODO implement me
	panic("implement me")
}

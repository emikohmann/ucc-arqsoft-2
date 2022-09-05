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

func (repo *RepositoryMemcached) Get(id string) (dtos.BookDTO, errors.ApiError) {
	item, err := repo.Client.Get(id)
	if err != nil {
		return dtos.BookDTO{}, errors.NewInternalServerApiError(fmt.Sprintf("error getting book %s", id), err)
	}

	var bookDTO dtos.BookDTO
	if err := json.Unmarshal(item.Value, &bookDTO); err != nil {
		return dtos.BookDTO{}, errors.NewInternalServerApiError(fmt.Sprintf("error getting book %s", id), err)
	}

	return bookDTO, nil
}

func (repo *RepositoryMemcached) Insert(book dtos.BookDTO) errors.ApiError {
	bytes, err := json.Marshal(book)
	if err != nil {
		return errors.NewBadRequestApiError(err.Error())
	}

	if err := repo.Client.Set(&memcache.Item{
		Key:   book.Id,
		Value: bytes,
	}); err != nil {
		return errors.NewInternalServerApiError(fmt.Sprintf("error inserting book %s", book.Id), err)
	}

	return nil
}

func (repo *RepositoryMemcached) Update(book dtos.BookDTO) errors.ApiError {
	bytes, err := json.Marshal(book)
	if err != nil {
		return errors.NewBadRequestApiError(fmt.Sprintf("invalid book %s: %v", book.Id, err))
	}

	if err := repo.Client.Set(&memcache.Item{
		Key:   book.Id,
		Value: bytes,
	}); err != nil {
		return errors.NewInternalServerApiError(fmt.Sprintf("error updating book %s", book.Id), err)
	}

	return nil
}

func (repo *RepositoryMemcached) Delete(id string) errors.ApiError {
	err := repo.Client.Delete(id)
	if err != nil {
		return errors.NewInternalServerApiError(fmt.Sprintf("error deleting book %s", id), err)
	}
	return nil
}

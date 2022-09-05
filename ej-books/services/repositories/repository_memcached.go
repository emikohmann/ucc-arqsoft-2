package repositories

import (
	"github.com/emikohmann/ucc-arqsoft-2/ej-books/dtos"
	"github.com/emikohmann/ucc-arqsoft-2/ej-books/utils/errors"
)

type RepositoryMemcached struct{}

func NewRepositoryMemcached() RepositoryMemcached {
	return RepositoryMemcached{}
}

func (RepositoryMemcached) Get(id string) (dtos.BookDto, errors.ApiError) {
	//TODO implement me
	panic("implement me")
}

func (RepositoryMemcached) Insert(dto dtos.BookDto) errors.ApiError {
	//TODO implement me
	panic("implement me")
}

func (RepositoryMemcached) Update(dto dtos.BookDto) errors.ApiError {
	//TODO implement me
	panic("implement me")
}

func (RepositoryMemcached) Delete(id string) errors.ApiError {
	//TODO implement me
	panic("implement me")
}

package repositories

import (
	"github.com/emikohmann/ucc-arqsoft-2/ej-books/dtos"
	"github.com/emikohmann/ucc-arqsoft-2/ej-books/utils/errors"
)

type Repository interface {
	Get(id string) (dtos.BookDto, errors.ApiError)
	Insert(dto dtos.BookDto) errors.ApiError
	Update(dto dtos.BookDto) errors.ApiError
	Delete(id string) errors.ApiError
}

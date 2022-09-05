package repositories

import (
	"github.com/emikohmann/ucc-arqsoft-2/ej-books/dtos"
	"github.com/emikohmann/ucc-arqsoft-2/ej-books/utils/errors"
)

type Repository interface {
	Get(id string) (dtos.BookDTO, errors.ApiError)
	Insert(dto dtos.BookDTO) errors.ApiError
	Update(dto dtos.BookDTO) errors.ApiError
	Delete(id string) errors.ApiError
}

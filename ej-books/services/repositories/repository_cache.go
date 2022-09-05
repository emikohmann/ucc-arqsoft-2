package repositories

import (
	"fmt"
	"github.com/emikohmann/ucc-arqsoft-2/ej-books/dtos"
	e "github.com/emikohmann/ucc-arqsoft-2/ej-books/utils/errors"
	"github.com/karlseguin/ccache/v2"
	"time"
)

type RepositoryCache struct {
	Client     *ccache.Cache
	DefaultTTL time.Duration
}

func NewRepositoryCache(
	maxSize int64,
	itemsToPrune uint32,
	defaultTTL time.Duration) *RepositoryCache {
	client := ccache.New(ccache.Configure().MaxSize(maxSize).ItemsToPrune(itemsToPrune))
	return &RepositoryCache{
		Client:     client,
		DefaultTTL: defaultTTL,
	}
}

func (repo *RepositoryCache) Get(id string) (dtos.BookDTO, e.ApiError) {
	item := repo.Client.Get(id)
	if item == nil {
		return dtos.BookDTO{}, e.NewNotFoundApiError(fmt.Sprintf("book %s not found", id))
	}
	return item.Value().(dtos.BookDTO), nil
}

func (repo *RepositoryCache) Insert(dto dtos.BookDTO) e.ApiError {
	repo.Client.Set(dto.Id, dto, repo.DefaultTTL)
	return nil
}

func (repo *RepositoryCache) Update(dto dtos.BookDTO) e.ApiError {
	repo.Client.Set(dto.Id, dto, repo.DefaultTTL)
	return nil
}

func (repo *RepositoryCache) Delete(id string) e.ApiError {
	repo.Client.Delete(id)
	return nil
}

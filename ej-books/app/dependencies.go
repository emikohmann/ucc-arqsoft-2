package app

import (
	"github.com/emikohmann/ucc-arqsoft-2/ej-books/controllers"
	service "github.com/emikohmann/ucc-arqsoft-2/ej-books/services"
	"github.com/emikohmann/ucc-arqsoft-2/ej-books/services/repositories"
	"time"
)

type Dependencies struct {
	BookController *controllers.Controller
}

func BuildDependencies() *Dependencies {
	// Repositories
	ccache := repositories.NewCCache(1000, 100, 30*time.Second)
	memcached := repositories.NewMemcached("localhost", 11211)
	mongo := repositories.NewMongoDB("localhost", 27017, "books")

	// Services
	service := service.NewServiceImpl(ccache, memcached, mongo)

	// Controllers
	controller := controllers.NewController(service)

	return &Dependencies{
		BookController: controller,
	}
}

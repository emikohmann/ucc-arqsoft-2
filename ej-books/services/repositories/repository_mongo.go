package repositories

import (
	"context"
	"fmt"
	"github.com/emikohmann/ucc-arqsoft-2/ej-books/dtos"
	"github.com/emikohmann/ucc-arqsoft-2/ej-books/utils/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RepositoryMongo struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func NewRepositoryMongo(
	user string,
	password string,
	host string,
	port int,
	source string,
	mechanism string) *RepositoryMongo {

	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(
			fmt.Sprintf("mongodb://%s:%s@%s:%d/?authSource=%s&authMechanism=%s",
				user, password, host, port, source, mechanism)))
	if err != nil {
		panic(fmt.Sprintf("Error initializing MongoDB: %v", err))
	}

	names, err := client.ListDatabaseNames(context.TODO(), bson.M{})
	if err != nil {
		panic(fmt.Sprintf("Error initializing MongoDB: %v", err))
	}
	fmt.Println(fmt.Sprintf("Available databases: %s", names))

	return &RepositoryMongo{
		Client:   client,
		Database: client.Database("books"),
	}
}

func (repo *RepositoryMongo) Get(id string) (dtos.BookDTO, errors.ApiError) {
	//TODO implement me
	panic("implement me")

}

func (repo *RepositoryMongo) Insert(dto dtos.BookDTO) errors.ApiError {
	//TODO implement me
	panic("implement me")
}

func (repo *RepositoryMongo) Update(dto dtos.BookDTO) errors.ApiError {
	//TODO implement me
	panic("implement me")
}

func (repo *RepositoryMongo) Delete(id string) errors.ApiError {
	//TODO implement me
	panic("implement me")
}

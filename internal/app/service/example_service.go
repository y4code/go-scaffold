package service

import (
	"go-scaffold/internal/app"
	"go-scaffold/internal/pkg/storage"
)

type ExampleService interface {
	CreateExample() app.Example
	GetExampleById(id int64) app.Example
	GetAllExamples() []app.Example
	UpdateExample(id int64, example app.Example) app.Example
	DeleteExampleById(id int64) app.Example
}

func NewExampleService(storage storage.Storage) ExampleService {
	return &exampleService{Storage: storage}
}

type exampleService struct {
	Storage storage.Storage
}

func (s *exampleService) CreateExample() app.Example {
	panic("implement me")
}

func (s *exampleService) GetExampleById(id int64) app.Example {
	panic("implement me")
}

func (s *exampleService) GetAllExamples() []app.Example {
	return s.Storage.GetAllExamples()
}

func (s *exampleService) UpdateExample(id int64, example app.Example) app.Example {
	panic("implement me")
}

func (s *exampleService) DeleteExampleById(id int64) app.Example {
	panic("implement me")
}

package service

import "showcase/repository"

type Service struct {
	repo repository.RepoInterface //Repository repository.ProductRepository
}

type ServiceInterface interface {
	BookService
}

func NewService(repo repository.RepoInterface) ServiceInterface {
	return &Service{repo: repo}
}

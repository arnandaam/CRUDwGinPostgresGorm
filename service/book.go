package service

import (
	"showcase/model"
)

type BookService interface {
	CreateBook(in model.Book) (res model.Book, err error)
	GetBookbyID(id int64) (res model.Book, err error)
	UpdateBookid(in model.Book) (res model.Book, err error)
	DeleteBookid(id int64) (err error)
	GetAllBooks() (res []model.Book, err error)
}

func (s *Service) CreateBook(in model.Book) (res model.Book, err error) {
	//cuma manggil dari repo
	res, err = s.repo.CreateBook(in)
	if err != nil {
		return res, err
	}

	return res, nil

	// return s.repo.CreateEmployee(in)
}

func (s *Service) GetBookbyID(id int64) (res model.Book, err error) {
	return s.repo.GetBookbyID(id) // cuma manggil dari repo
}

func (s *Service) UpdateBookid(in model.Book) (res model.Book, err error) {
	return s.repo.UpdateBookid(in)
}

func (s *Service) DeleteBookid(id int64) (err error) {
	return s.repo.DeleteBookid(id)
}

func (s *Service) GetAllBooks() (res []model.Book, err error) {
	return s.repo.GetAllBooks()
}

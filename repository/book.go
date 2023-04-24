package repository

import (
	"showcase/model"
)

// interface book
type BookRepo interface {
	GetBookbyID(id int64) (res model.Book, err error)
	CreateBook(in model.Book) (res model.Book, err error)
	UpdateBookid(in model.Book) (res model.Book, err error)
	DeleteBookid(id int64) (err error)
	GetAllBooks() (res []model.Book, err error)
}

// func (r Repo) CreateBook(in model.Book) (res model.Book, err error) {
// 	r.db.QueryRow(
// 		query.AddBook,
// 		in.Name,
// 		in.Author,
// 	).Scan(
// 		&res.BookID,
// 		&res.Name,
// 		&res.Author,
// 	)
// 	if err != nil {
// 		return res, err
// 	}

// 	return res, nil
// }

func (r Repo) CreateBook(in model.Book) (res model.Book, err error) {
	err = r.gorm.Create(&in).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetBookbyID(id int64) (res model.Book, err error) {
	err = r.gorm.First(&res, id).Error //
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r Repo) UpdateBookid(in model.Book) (res model.Book, err error) {
	err = r.gorm.Model(&res).Where("id = ?", in.ID).Updates(model.Book{ // update hanya 1 kolom updates semua kolom
		Name:   in.Name,
		Author: in.Author,
	}).Scan(&res).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) DeleteBookid(id int64) (err error) {
	book := model.Book{}
	err = r.gorm.Model(&book).Where("id = ?", id).Delete(&book).Error //Update("deleted_at", time.Now()).Error
	if err != nil {
		return err
	}

	return nil
}

func (r Repo) GetAllBooks() (res []model.Book, err error) {
	err = r.gorm.Find(&res).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

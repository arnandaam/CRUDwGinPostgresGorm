package handler

import (
	"showcase/helper"
	"showcase/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h HttpServer) CreateBook(c *gin.Context) {
	in := model.Book{}

	err := c.BindJSON(&in)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	err = in.Validation() // custom validasi
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	// panggil servise
	res, err := h.app.CreateBook(in)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

func (h HttpServer) GetBookbyID(c *gin.Context) {
	id := c.Param("id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	// panggil servise
	res, err := h.app.GetBookbyID(int64(idint))
	if err != nil {
		helper.NotFound(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

// func (h HttpServer) GetBookbyID(c *gin.Context) {
// 	id := c.Param("id")
// 	idint, err := strconv.Atoi(id)
// 	if err != nil {
// 		helper.BadRequest(c, err.Error())
// 		return
// 	}

// 	// panggil servise
// 	res, err := h.app.GetBookbyID(int64(idint))
// 	if err != nil {
// 		helper.NotFound(c, err.Error())
// 		return
// 	}

// 	// membuat map untuk response
// 	resMap := map[string]interface{}{
// 		"id":         res.ID,
// 		"name_book":  res.Name,
// 		"author":     res.Author,
// 		"created_at": res.CreatedAt,
// 		"updated_at": res.UpdatedAt,
// 	}

// 	helper.Ok(c, resMap)
// }

func (h HttpServer) UpdateBookid(c *gin.Context) {
	id := c.Param("id")
	idint, err := strconv.Atoi(id)

	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	in := model.Book{}

	err = c.BindJSON(&in)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	// err = in.Validation() // custom validasi // gunakan kalau semua harus di update
	// if err != nil {
	// 	helper.BadRequest(c, err.Error())
	// 	return
	// }
	in.ID = idint
	// panggil servise
	res, err := h.app.UpdateBookid(in)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

func (h HttpServer) DeleteBookid(c *gin.Context) {
	id := c.Param("id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	// panggil servise
	err = h.app.DeleteBookid(int64(idint))
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.OkWithMessage(c, "Book Deletes Successfully")
}

func (h HttpServer) GetAllBooks(c *gin.Context) {
	// panggil service
	books, err := h.app.GetAllBooks()
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	// cek apakah ada data yang ditemukan
	if len(books) == 0 {
		helper.NotFound(c, "No books found")
		return
	}

	helper.Ok(c, books)
}

package route

import (
	"showcase/handler"
	"showcase/service"

	"github.com/gin-gonic/gin"
)

func RegisterApi(r *gin.Engine, app service.ServiceInterface) {
	server := handler.NewHttpServer(app)
	api := r.Group("/showbook")
	{
		api.POST("", server.CreateBook)
		api.GET("/:id", server.GetBookbyID)
		api.PUT("/:id", server.UpdateBookid)
		api.DELETE("/:id", server.DeleteBookid)
		api.GET("", server.GetAllBooks)
	}
}

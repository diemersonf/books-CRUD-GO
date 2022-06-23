package routes

import (
	"github.com/diemersonf/books-crud-go/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(rounter *gin.Engine) *gin.Engine {
	main := rounter.Group("api/v1")
	{
		books := main.Group("books")
		{
			books.GET("/", controllers.ShowBook)
		}
	}

	return rounter
}

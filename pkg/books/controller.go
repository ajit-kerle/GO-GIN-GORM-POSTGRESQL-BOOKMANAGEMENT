package books

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	// pointer to access data by reference
	DB *gorm.DB
}

func RegisterRoutes (r *gin.Engine, db *gorm.DB) {
	h:= &handler{
		DB:db,
	}

	routes:=r.Group("/books")
	// mini routes
	routes.POST("/",h.AddBook)
	routes.GET("/",h.GetBooks)
	routes.GET("/:id",h.GetBook)
	routes.PUT("/:id",h.UpdateBook)
	routes.DELETE("/:id",h.DeleteBook)
}
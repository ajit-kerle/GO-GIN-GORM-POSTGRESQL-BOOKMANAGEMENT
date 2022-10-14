package author

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


type handler struct {
	// pointer to access data by reference
	DB *gorm.DB
}


func RegisterRoutesForAuthor (r *gin.Engine, db *gorm.DB) {
	h:= &handler{
		DB:db,
	}

	routes:=r.Group("/author")
	// mini routes
	
	routes.POST("/",h.AddAuthor)
	routes.POST("/publisher",h.AddPublisher)
	
}
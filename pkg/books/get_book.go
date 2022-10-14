package books

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/ajit-kerle/go-gin-api-medium/pkg/common/models"

)

func (h handler) GetBook(c *gin.Context) {
	id:= c.Param("id")

	var book models.Book
    
	result:=h.DB.First(&book,id)

	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
        return
	}
    c.JSON(http.StatusOK, &book) 
}
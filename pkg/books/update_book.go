package books

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/ajit-kerle/go-gin-api-medium/pkg/common/models"

)

type UpdateBookRequestBody struct {
	Title string `json:"title"`
	AuthorID int64 `json:"author_id"`
	Description string `json:"description"`
	PublisherID int64 `json:"publisher_id"`
}


func (h handler) UpdateBook(c *gin.Context) {
	id:=c.Param("id")
	body:=UpdateBookRequestBody{}

	// getting request body
    err := c.BindJSON(&body)

	if err != nil {
        c.AbortWithError(http.StatusBadRequest, err)
        return
    }

	var book models.Book

	// result fields

	result := h.DB.First(&book, id) // First getby id

	if result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    book.Title = body.Title
    book.AuthorID = body.AuthorID
    book.Description = body.Description 
	book.PublisherID=body.PublisherID

	h.DB.Save(&book)

    c.JSON(http.StatusOK, &book)

}



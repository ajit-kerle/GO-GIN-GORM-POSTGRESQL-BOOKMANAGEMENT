package books

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/ajit-kerle/go-gin-api-medium/pkg/common/models"

)

// structure of Books add
type AddBookRequestBody struct {
	Title string `json:"title"`
	AuthorID int64 `json:"author_id"`
	Description string `json:"description"`
	PublisherID int64 `json:"publisher_id"`
}


func (h handler) AddBook(c *gin.Context) {
	body:= AddBookRequestBody{}

	// getting request's body

	err:=c.BindJSON(&body);

	if err!=nil {
		c.AbortWithError(http.StatusBadRequest,err)
		return 
	}

	var book models.Book

	book.Title=body.Title
	book.AuthorID=body.AuthorID
	book.Description=body.Description
	book.PublisherID=body.PublisherID

	result:=h.DB.Create(&book)

	if result.Error!=nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
        return
	}
	c.JSON(http.StatusCreated, &book)

}






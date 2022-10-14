package author

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/ajit-kerle/go-gin-api-medium/pkg/common/models"
)


type AddAuthorRequestBody struct {
	Name string `json:"name"`
	Address string `json:"address"`
}

func (h handler) AddAuthor (c *gin.Context){
    body:=AddAuthorRequestBody{}

	err:=c.BindJSON(&body) // BindJSON reads the body buffer to de-serialize it to a struct

	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
		return 
	}

	var author models.Author

	author.Name=body.Name
	author.Address=body.Address

	result:=h.DB.Create(&author)

	if result.Error!=nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
        return
	}

	c.JSON(http.StatusCreated, &author)
}

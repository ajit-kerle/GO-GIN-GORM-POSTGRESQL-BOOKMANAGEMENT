package author

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/ajit-kerle/go-gin-api-medium/pkg/common/models"
)


type AddPublisherRequestBody struct {
	
	PublisherName string `json:"pname"`
	PublisherAddress string `json:"paddress"`
}

func (h handler) AddPublisher (c *gin.Context){
    body:=AddPublisherRequestBody{}

	err:=c.BindJSON(&body)
    
	if err!=nil {
       c.AbortWithError(http.StatusBadRequest,err)
		return 
	}

	var publisher models.Publisher

	
	publisher.PublisherName=body.PublisherName
	publisher.PublisherAddress=body.PublisherAddress

	result:=h.DB.Create(&publisher)

	if result.Error!=nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
        return
	}
	c.JSON(http.StatusOK,&publisher)

	
} 
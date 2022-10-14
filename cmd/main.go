package main

import (
    "github.com/gin-gonic/gin"
    "github.com/ajit-kerle/go-gin-api-medium/pkg/books"
    "github.com/ajit-kerle/go-gin-api-medium/pkg/author"
    "github.com/ajit-kerle/go-gin-api-medium/pkg/common/db"
    "github.com/spf13/viper"
)


func main() {
    viper.SetConfigFile("./pkg/common/envs/.env")
    viper.ReadInConfig()

    port := viper.Get("PORT").(string)
    dbUrl := viper.Get("DB_URL").(string)

    // gin.SetMode(gin.ReleaseMode)
    // r:= gin.New()
    r:= gin.Default()
    h:=db.Init(dbUrl)


    books.RegisterRoutes(r,h)
    author.RegisterRoutesForAuthor(r,h)

    // r.GET("/", func(c *gin.Context) {
    //     c.JSON(200, gin.H{
    //         "port": port,
    //         "dbUrl": dbUrl,
    //     })
    // })

    r.Run(port)
}
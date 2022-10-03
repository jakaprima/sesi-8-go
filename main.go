package main

import (
	"fmt"
	"golang-swagger/controller"
	"golang-swagger/db"

	_ "golang-swagger/docs"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger Latihan 8
// @version         1.0
// @description     sample API docs
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

func main() {
	db := db.ConnectGorm()
	fmt.Println("db", db)

	r := gin.Default()

	c := controller.NewController()

	v1 := r.Group("/api/v1")
	{
		peoples := v1.Group("/peoples")
		{
			peoples.GET("ping", c.PingPeoples)
			peoples.GET(":id", c.ShowPeople)
			peoples.GET("", c.ListPeoples)
			peoples.POST("", c.AddPeople)
			peoples.DELETE(":id", c.DeletePeople)
			peoples.PATCH(":id", c.UpdatePeople)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}

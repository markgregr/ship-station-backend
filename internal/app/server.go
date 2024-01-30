package app

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/markgregr/RIP/docs"
	"github.com/markgregr/RIP/internal/pkg/middleware"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func (app *Application) Run() {
    r := gin.Default()  
    // Это нужно для автоматического создания папки "docs" в вашем проекте
	docs.SwaggerInfo.Title = "ShipStation RestAPI"
	docs.SwaggerInfo.Description = "API server for ShipStation application"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    // Группа запросов для багажа
    ShipGroup := r.Group("/baggage")
    {   
        ShipGroup.GET("/", middleware.Guest(app.Repository.GetRedisClient(), []byte("AccessSecretKey"), app.Repository), app.Handler.GetShips)
        ShipGroup.GET("/:shipID", middleware.Guest(app.Repository.GetRedisClient(), []byte("AccessSecretKey"), app.Repository), app.Handler.GetShipByID)
        ShipGroup.DELETE("/:shipID", middleware.Authenticate(app.Repository.GetRedisClient(), []byte("AccessSecretKey"), app.Repository), app.Handler.DeleteShip)
        ShipGroup.POST("/", middleware.Authenticate(app.Repository.GetRedisClient(), []byte("AccessSecretKey"), app.Repository), app.Handler.CreateShip)
        ShipGroup.PUT("/:shipID", middleware.Authenticate(app.Repository.GetRedisClient(), []byte("AccessSecretKey"), app.Repository), app.Handler.UpdateShip)
        ShipGroup.POST("/:shipID/request", middleware.Authenticate(app.Repository.GetRedisClient(), []byte("AccessSecretKey"), app.Repository), app.Handler.AddShipToRequest)
        ShipGroup.DELETE("/:shipID/request", middleware.Authenticate(app.Repository.GetRedisClient(), []byte("AccessSecretKey"), app.Repository), app.Handler.RemoveShipFromRequest)
        ShipGroup.POST("/:shipID/image", middleware.Authenticate(app.Repository.GetRedisClient(), []byte("AccessSecretKey"), app.Repository), app.Handler.AddShipImage)
    }
    

    // Группа запросов для доставки
    RequestGroup := r.Group("/request").Use(middleware.Authenticate(app.Repository.GetRedisClient(), []byte("AccessSecretKey"), app.Repository))
    {
        RequestGroup.GET("/", app.Handler.GetRequests)
        RequestGroup.GET("/:requestID", app.Handler.GetRequestByID)
        RequestGroup.DELETE("/:requestID", app.Handler.DeleteRequest)
        RequestGroup.PUT("/:requestID/status/user", app.Handler.UpdateRequestStatusUser)  // Новый маршрут для обновления статуса доставки пользователем
        RequestGroup.PUT("/:requestID/status/moderator", app.Handler.UpdateRequestStatusModerator)  // Новый маршрут для обновления статуса доставки модератором
    }

    UserGroup := r.Group("/user")
    {
        UserGroup.POST("/register", app.Handler.Register)
        UserGroup.POST("/login", app.Handler.Login)
        UserGroup.POST("/logout", middleware.Authenticate(app.Repository.GetRedisClient(), []byte("AccessSecretKey"), app.Repository), app.Handler.Logout)
    }
    addr := fmt.Sprintf("%s:%d", app.Config.ServiceHost, app.Config.ServicePort)
    r.Run(addr)
    log.Println("Server down")
}
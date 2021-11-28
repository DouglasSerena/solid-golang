package routes

import (
	"github.com/DouglasSerena/solid-go/src/app/presentation/controllers"
	"github.com/gin-gonic/gin"
)

func UserSetupRoutes(controlls controllers.IUserController, router gin.IRouter) {
	router.GET("/user", controlls.Find)
	router.POST("/user", controlls.Insert)
	router.PUT("/user/:id", controlls.Update)
	router.GET("/user/:id", controlls.FindBy)
	router.DELETE("/user/:id", controlls.Delete)
}

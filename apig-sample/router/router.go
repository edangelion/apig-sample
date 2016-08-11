package router

import (
	"github.com/wantedly/apig-sample/controllers"

	"github.com/gin-gonic/gin"
)

func Initialize(r *gin.Engine) {
	api := r.Group("")
	{

		api.GET("/emails", controllers.GetEmails)
		api.GET("/emails/:id", controllers.GetEmail)
		api.POST("/emails", controllers.CreateEmail)
		api.PUT("/emails/:id", controllers.UpdateEmail)
		api.DELETE("/emails/:id", controllers.DeleteEmail)

		api.GET("/users", controllers.GetUsers)
		api.GET("/users/:id", controllers.GetUser)
		api.POST("/users", controllers.CreateUser)
		api.PUT("/users/:id", controllers.UpdateUser)
		api.DELETE("/users/:id", controllers.DeleteUser)

	}
}

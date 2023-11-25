package server

import (
	"github.com/yashipro13/queryMaster/handlers"
	"github.com/yashipro13/queryMaster/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(userService users.Service) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "pong"})
	})

	r.GET("/get_project_by_user", handlers.GetProjectByUserIDHandler(userService))

	return r
}

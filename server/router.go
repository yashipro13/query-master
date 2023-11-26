package server

import (
	"github.com/yashipro13/queryMaster/elasticsearch"
	"github.com/yashipro13/queryMaster/handlers"
	"github.com/yashipro13/queryMaster/hashtags"
	"github.com/yashipro13/queryMaster/repository"
	"github.com/yashipro13/queryMaster/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(userService users.Service, hashtagService hashtags.Service, esService elasticsearch.Service, repo *repository.Repo) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "pong"})
	})

	r.GET("/get_project_by_user", handlers.GetProjectByUserIDHandler(userService))
	r.GET("/get_project_by_hashtags", handlers.GetProjectByHashtagsHandler(hashtagService))

	r.GET("/get_project_by_search", handlers.FuzzySearchHandler(esService))
	r.POST("/mock-data", handlers.MockInsertionHandler(repo))

	return r
}

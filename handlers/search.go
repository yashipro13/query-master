package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/yashipro13/queryMaster/elasticsearch"
	"log"
	"net/http"
	"time"
)

func FuzzySearchHandler(svc elasticsearch.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		searchTerm := c.Query("query")
		deadlineCtx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))
		defer cancel()
		projects, svcErr := svc.Search(deadlineCtx, searchTerm)
		if svcErr != nil {
			log.Printf("%s", svcErr.Message)
			c.JSON(svcErr.Code, gin.H{"status": false, "message": svcErr.Message})
			return

		}
		c.JSON(http.StatusOK, gin.H{"success": true, "data": projects})
	}
}

package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/yashipro13/queryMaster/hashtags"
	"log"
	"net/http"
	"time"
)

func GetProjectByHashtagsHandler(svc hashtags.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		inputHashtags := c.QueryArray("hashtags")
		if len(inputHashtags) == 0 {
			c.JSON(400, gin.H{"success": false, "error": "hashtags cannot be empty"})
		}
		deadlineCtx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))
		defer cancel()
		svcRes := svc.Run(deadlineCtx, inputHashtags)
		if svcRes.Error != nil {
			log.Printf("%s", svcRes.Error.Message)
			c.JSON(svcRes.Error.Code, gin.H{"status": false, "message": svcRes.Error.Message})
			return

		}
		c.JSON(http.StatusOK, gin.H{"success": true, "data": svcRes.Data})
	}
}

package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/yashipro13/queryMaster/users"
	"log"
	"net/http"
	"strconv"
	"time"
)

func GetProjectByUserIDHandler(svc users.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDStr := c.Query("user_id")
		userID, _ := strconv.ParseInt(userIDStr, 10, 32)
		if userID == 0 {
			c.JSON(400, gin.H{"success": false, "error": "user id cannot be empty"})
		}
		ctx := context.WithValue(context.Background(), "userID", userID)
		deadlineCtx, cancel := context.WithDeadline(ctx, time.Now().Add(1*time.Second))
		defer cancel()
		svcRes := svc.Run(deadlineCtx, int(userID))
		if svcRes.Error != nil {
			log.Printf("%s", svcRes.Error.Message)
			c.JSON(svcRes.Error.Code, gin.H{"status": false, "message": svcRes.Error.Message})
			return

		}
		c.JSON(http.StatusOK, gin.H{"success": true, "data": svcRes.Data})
	}
}

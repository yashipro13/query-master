package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/yashipro13/queryMaster/repository"
	"net/http"
)

// MockInsertionHandler is a mock handler, this has hardcoded data, that is just being inserted for the demo of
// elastic search live sync, to make this work, ensure you have `seed_insertion.sql` already inserted in your database
func MockInsertionHandler(repo *repository.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		repo.InsertMockDataForUser3(context.Background())
		c.JSON(http.StatusOK, gin.H{"success": true})
	}
}

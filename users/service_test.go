package users

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/yashipro13/queryMaster/models"
	"github.com/yashipro13/queryMaster/users/mocks"
	"testing"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mocks.NewMockDBManager(ctrl)
	svc := Service{DBManager: mockDB}
	t.Run("should return nil data when db call fails", func(t *testing.T) {
		mockDB.EXPECT().GetProjectsByUser(context.Background(), 1).Return([]models.Project{}, &models.Error{
			Message: "db call failed",
			Err:     fmt.Errorf("db call failed"),
		})
		res := svc.Run(context.Background(), 1)
		assert.Equal(t, models.Response{
			Success: false,
			Data:    nil,
			Error: &models.Error{
				Message: "db call failed",
				Err:     fmt.Errorf("db call failed"),
			},
		}, res)
	})

	t.Run("should return data and nil error when db call succeeds", func(t *testing.T) {
		mockDB.EXPECT().GetProjectsByUser(context.Background(), 1).Return([]models.Project{{ID: 1}}, nil)
		res := svc.Run(context.Background(), 1)
		assert.Equal(t, models.Response{
			Success: true,
			Data:    []models.Project{{ID: 1}},
			Error:   nil,
		}, res)
	})

}

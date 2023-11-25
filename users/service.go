package users

import (
	"context"
	"github.com/yashipro13/queryMaster/models"
)

type Service struct {
	DBManager DBManager
}

type DBManager interface {
	GetProjectsByUser(ctx context.Context, userID int) ([]models.Project, *models.Error)
}

func (svc Service) Run(ctx context.Context, userID int) models.Response {
	projects, svcErr := svc.DBManager.GetProjectsByUser(ctx, userID)
	if svcErr != nil {
		return models.Response{
			Success: false,
			Data:    nil,
			Error:   svcErr,
		}
	}
	return models.Response{
		Success: true,
		Data:    projects,
		Error:   nil,
	}
}

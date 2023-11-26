package hashtags

import (
	"context"
	"github.com/yashipro13/queryMaster/models"
)

type Service struct {
	DBManager DBManager
}

type DBManager interface {
	GetProjectByHashtags(ctx context.Context, hashtags []string) ([]models.Project, *models.Error)
}

func (svc Service) Run(ctx context.Context, hashtags []string) models.Response {
	projects, svcErr := svc.DBManager.GetProjectByHashtags(ctx, hashtags)
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

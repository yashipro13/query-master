package ingestor

import (
	"context"
	"fmt"
	"github.com/yashipro13/queryMaster/models"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"time"
)

type Service struct {
	DBManager     DBManager
	ElasticClient *elastic.Client
}

var lastSyncTime time.Time
var lastIternationsCount models.Count

type DBManager interface {
	GetAllProjects(ctx context.Context) ([]models.Project, *models.Error)
	IsCUDPerformed(ctx context.Context, lastSyncTime time.Time, lastCount models.Count) (bool, models.Count, *models.Error)
}

func (svc *Service) IngestData(ctx context.Context) *models.Error {
	lastSyncTime = time.Now().UTC()
	var isCUD bool
	var repoErr *models.Error
	isCUD, lastIternationsCount, repoErr = svc.DBManager.IsCUDPerformed(ctx, lastSyncTime, lastIternationsCount)
	lastSyncTime = time.Now()
	if repoErr != nil {
		return repoErr
	}
	if isCUD {
		log.Println("cud is performed, syncing again")
		return svc.syncData(ctx)
	}
	log.Printf("no change, no sync needed")
	return nil
}

func (svc *Service) syncData(ctx context.Context) *models.Error {
	projects, repoErr := svc.DBManager.GetAllProjects(ctx)
	if repoErr != nil {
		return repoErr
	}
	bulk := svc.ElasticClient.Bulk()
	for _, project := range projects {
		req := elastic.NewBulkIndexRequest().
			Index("description_and_slug_index").
			Type("description").
			Id(fmt.Sprintf("%s_%s", project.Description, project.Slug)). // Combine both values into a single ID
			Doc(project)
		bulk.Add(req)
	}

	_, err := bulk.Do(ctx)
	if err != nil {
		log.Println("Error indexing data into Elasticsearch:", err)
		return &models.Error{
			Code:    500,
			Message: fmt.Sprintf("indexing failed with err %s", err.Error()),
			Err:     err,
		}
	}
	return nil
}

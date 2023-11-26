package elasticsearch

import (
	"context"
	"encoding/json"
	"github.com/yashipro13/queryMaster/models"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

type Service struct {
	ElasticClient *elastic.Client
}

func (svc *Service) Search(ctx context.Context, searchTerm string) ([]models.Project, *models.Error) {

	searchResult, err := svc.ElasticClient.Search().
		Index("description_and_slug_index").
		Query(elastic.NewQueryStringQuery(searchTerm)).
		Do(context.Background())
	if err != nil {
		log.Printf("Error performing the search: %s", err)
	}
	log.Printf("search successful ")
	var projects []models.Project
	for _, hit := range searchResult.Hits.Hits {
		rawBytes, err := hit.Source.MarshalJSON()
		if err != nil {
			log.Printf("source is %v, but failed marshalling it %s", hit.Source, err.Error())
			return []models.Project{}, &models.Error{
				Code:    500,
				Message: "failed to get bytes",
				Err:     err,
			}
		}
		var project models.Project
		err = json.Unmarshal(rawBytes, &project)
		if err != nil {
			return []models.Project{}, &models.Error{
				Code:    500,
				Message: "failed to unmarshal",
				Err:     err,
			}
		}
		projects = append(projects, project)
	}
	return projects, nil
}

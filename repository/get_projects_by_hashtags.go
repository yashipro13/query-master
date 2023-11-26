package repository

import (
	"context"
	"fmt"
	"github.com/lib/pq"
	"github.com/yashipro13/queryMaster/models"
	"log"
)

func (r *Repo) GetProjectByHashtags(ctx context.Context, hashtags []string) ([]models.Project, *models.Error) {
	hashtagIDs, err := r.getHashtagIDsbyHastags(ctx, hashtags)
	if err != nil {
		return []models.Project{}, &models.Error{
			Code:    500,
			Message: "failed to get hashtag id for this hashtag name",
			Err:     err,
		}
	}
	projectIDs, err := r.findProjectIDsByHashtagsIDs(ctx, hashtagIDs)
	if err != nil {
		return []models.Project{}, &models.Error{
			Code:    500,
			Message: "failed to get project id for this hashtag name",
			Err:     err,
		}
	}
	projects, err := r.findProjectsByProjectID(ctx, projectIDs)
	if err != nil {
		return []models.Project{}, &models.Error{
			Code:    500,
			Message: fmt.Sprintf("could find projects err %s", err.Error()),
			Err:     err,
		}
	}
	var domainProjects []models.Project
	for _, project := range projects {
		user, _ := r.findUserNamesByProjectID(ctx, project.id)
		log.Printf("name for project id %d is %v", project.id, user)
		projectHashtags, _ := r.findHastagsIDsByProjectID(ctx, project.id)
		domainProjects = append(domainProjects, models.Project{
			ID:          project.id,
			Name:        project.name,
			Slug:        project.slug,
			Description: project.description,
			CreatedAt:   project.createdAt,
			Hashtags:    projectHashtags,
			CreatedBy:   user,
		})
	}
	return domainProjects, nil

}

func (r *Repo) getHashtagIDsbyHastags(ctx context.Context, hashtags []string) ([]int, error) {
	rows, err := r.db.Query(ctx, "SELECT id FROM hashtags WHERE name = ANY($1)", pq.Array(hashtags))
	if err != nil {
		return []int{}, err
	}
	var hashtagIDs []int
	for rows.Next() {
		var hashtagID int
		err = rows.Scan(&hashtagID)
		if err != nil {
			break
		}
		hashtagIDs = append(hashtagIDs, hashtagID)
	}
	return hashtagIDs, nil
}

func (r *Repo) findProjectIDsByHashtagsIDs(ctx context.Context, hashtagIDs []int) ([]int, error) {
	rows, err := r.db.Query(ctx, "SELECT project_id FROM project_hashtags WHERE hashtag_id = ANY($1)", pq.Array(hashtagIDs))
	if err != nil {
		return []int{}, err
	}
	var projectIDs []int
	for rows.Next() {
		var projectID int
		err = rows.Scan(&projectID)
		if err != nil {
			break
		}
		projectIDs = append(projectIDs, projectID)
	}
	return projectIDs, nil
}

func (r *Repo) findUserNamesByProjectID(ctx context.Context, projectID int) ([]string, error) {
	row, err := r.db.Query(ctx, "SELECT name FROM users WHERE id = ANY(select user_id from user_projects where project_id = $1)", projectID)
	if err != nil {
		log.Printf("failed to find user id for project %d with err %s", projectID, err.Error())
		return []string{}, err
	}
	var names []string
	for row.Next() {
		var name string
		err = row.Scan(&name)
		if err != nil {
			log.Printf("failed to find id with err %s", err.Error())
		}
		log.Printf("appended to names %s", name)
		names = append(names, name)
	}
	return names, nil
}

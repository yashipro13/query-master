package repository

import (
	"context"
	"fmt"
	"github.com/lib/pq"
	"github.com/yashipro13/queryMaster/models"
	"time"
)

func (r *Repo) GetProjectsByUser(ctx context.Context, userID int) ([]models.Project, *models.Error) {
	projectIDs, err := r.findProjectIDsForUser(ctx, userID)
	if err != nil && err.Error() == "no rows found" {
		return []models.Project{}, &models.Error{
			Code:    422,
			Message: "no project ids found for selected user",
			Err:     err,
		}
	}
	if err != nil {
		return []models.Project{}, &models.Error{
			Code:    500,
			Message: fmt.Sprintf("project id couldn't be found with err %s", err.Error()),
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
	userName := r.findUserNameByID(ctx, userID)
	var domainProjects []models.Project
	for _, project := range projects {
		hashtags, _ := r.findHastagsIDsByProjectID(ctx, project.id)
		domainProjects = append(domainProjects, models.Project{
			ID:          project.id,
			Name:        project.name,
			Slug:        project.slug,
			Description: project.description,
			CreatedAt:   project.createdAt,
			Hashtags:    hashtags,
			CreatedBy:   userName,
		})
	}
	return domainProjects, nil
}

func (r *Repo) findProjectIDsForUser(ctx context.Context, userID int) ([]int, error) {
	rows, err := r.db.Query(ctx, "select project_id from user_projects where user_id = $1", userID)
	if err != nil {
		return []int{}, err
	}
	var projectIDs []int
	for rows.Next() {
		var projectID int
		err := rows.Scan(&projectID)
		if err != nil {
			break
		}
		projectIDs = append(projectIDs, projectID)
	}
	if len(projectIDs) == 0 {
		return projectIDs, fmt.Errorf("no rows found")
	}
	return projectIDs, nil
}

func (r *Repo) findProjectsByProjectID(ctx context.Context, projectID []int) ([]project, error) {
	rows, err := r.db.Query(ctx, "SELECT id, name, slug, description, created_at FROM projects WHERE id = ANY($1)", pq.Array(projectID))
	if err != nil {
		return []project{}, err
	}
	var projects []project
	for rows.Next() {
		var projectID int
		var name, slug, description string
		var createdAt time.Time
		err := rows.Scan(&projectID, &name, &slug, &description, &createdAt)
		if err != nil {
			break
		}
		projects = append(projects, project{
			id:          projectID,
			name:        name,
			slug:        slug,
			description: description,
			createdAt:   createdAt,
		})
	}
	return projects, nil
}

func (r *Repo) findHastagsIDsByProjectID(ctx context.Context, projectID int) ([]string, error) {
	rows, err := r.db.Query(ctx, "SELECT hashtag_id FROM project_hashtags WHERE project_id = $1", projectID)
	if err != nil {
		return []string{}, err
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
	var hashtags []string
	rows, err = r.db.Query(ctx, "SELECT name FROM hashtags WHERE id = ANY($1)", pq.Array(hashtagIDs))
	if err != nil {
		return []string{}, err
	}
	for rows.Next() {
		var hashtag string
		err = rows.Scan(&hashtag)
		if err != nil {
			break
		}
		hashtags = append(hashtags, hashtag)
	}
	return hashtags, nil
}

func (r *Repo) findUserNameByID(ctx context.Context, userID int) string {
	var name string
	row, _ := r.db.Query(ctx, "SELECT name FROM users WHERE id = $1", userID)
	row.Next()
	_ = row.Scan(&name)
	return name
}

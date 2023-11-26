package repository

import (
	"context"
	"fmt"
	"github.com/yashipro13/queryMaster/models"
	"log"
	"time"
)

func (r *Repo) GetAllProjects(ctx context.Context) ([]models.Project, *models.Error) {
	rows, err := r.db.Query(ctx, "SELECT id, name, slug, description, created_at FROM projects")
	if err != nil {
		return []models.Project{}, &models.Error{
			Code:    500,
			Message: fmt.Sprintf("could not find projects with err %s", err.Error()),
			Err:     err,
		}
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

	var domainProjects []models.Project
	for _, project := range projects {
		userName, _ := r.findUserNamesByProjectID(ctx, project.id)
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

func (r *Repo) IsCUDPerformed(ctx context.Context, lastSyncTime time.Time, LastIterationCounts models.Count) (bool, models.Count, *models.Error) {
	log.Printf("checked cud applicability")
	var isCUDPerformed bool
	rows, err := r.db.Query(ctx, "SELECT updated_at FROM projects")
	if err != nil {
		log.Printf("projects not found")
		return false, models.Count{}, &models.Error{
			Code:    500,
			Message: fmt.Sprintf("could not find projects with err %s", err.Error()),
			Err:     err,
		}
	}
	var projectCount int
	for rows.Next() {
		var updatedAt time.Time
		err := rows.Scan(&updatedAt)
		if err != nil {
			break
		}
		if lastSyncTime.Before(updatedAt) {
			isCUDPerformed = true
		}
		projectCount += 1
	}
	if projectCount != LastIterationCounts.Projects {
		isCUDPerformed = true
	}
	rows, err = r.db.Query(ctx, "SELECT updated_at FROM users")
	if err != nil {
		log.Printf("users not found")
		return false, models.Count{}, &models.Error{
			Code:    500,
			Message: fmt.Sprintf("could not find users with err %s", err.Error()),
			Err:     err,
		}
	}
	var userCount int
	for rows.Next() {
		var updatedAt time.Time
		err := rows.Scan(&updatedAt)
		if err != nil {
			break
		}
		if lastSyncTime.Before(updatedAt) {
			isCUDPerformed = true
		}
		userCount += 1
	}
	if userCount != LastIterationCounts.Projects {
		isCUDPerformed = true
	}

	rows, err = r.db.Query(ctx, "SELECT updated_at FROM hashtags")
	if err != nil {
		log.Printf("hashtags not found")
		return false, models.Count{}, &models.Error{
			Code:    500,
			Message: fmt.Sprintf("could not find projects with err %s", err.Error()),
			Err:     err,
		}
	}
	var hashtagCount int
	for rows.Next() {
		var updatedAt time.Time
		err := rows.Scan(&updatedAt)
		if err != nil {
			break
		}
		if lastSyncTime.Before(updatedAt) {
			isCUDPerformed = true
		}
		hashtagCount += 1
	}
	if hashtagCount != LastIterationCounts.Projects {
		isCUDPerformed = true
	}
	return isCUDPerformed, models.Count{
		Users:    userCount,
		Projects: projectCount,
		Hashatgs: hashtagCount,
	}, nil
}

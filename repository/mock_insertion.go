package repository

import (
	"context"
	"time"
)

func (r *Repo) InsertMockDataForUser3(ctx context.Context) {
	r.db.Exec(ctx, "INSERT INTO projects (name, slug, description, created_at) values ($1, $2, $3, $4)",
		"Project 4", "slug for project4", "descriptive", time.Now())
	r.db.Exec(ctx, "INSERT INTO user_projects (user_id, project_1) values ($1, $2)",
		3, 4)
}

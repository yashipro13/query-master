ALTER TABLE "project_hashtags" ADD CONSTRAINT  "c2" FOREIGN KEY ("project_id") REFERENCES "projects" ("id");

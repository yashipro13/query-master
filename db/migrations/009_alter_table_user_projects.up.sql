ALTER TABLE "user_projects" add constraint  "c4" FOREIGN KEY ("user_id") REFERENCES "users" ("id");

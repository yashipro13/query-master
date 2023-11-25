INSERT INTO "users" ("name", "created_at") VALUES
                                               ('User1', CURRENT_TIMESTAMP),
                                               ('User2', CURRENT_TIMESTAMP),
                                               ('User3', CURRENT_TIMESTAMP);

INSERT INTO "hashtags" ("name", "created_at") VALUES
                                                  ('Tag1', CURRENT_TIMESTAMP),
                                                  ('Tag2', CURRENT_TIMESTAMP),
                                                  ('Tag3', CURRENT_TIMESTAMP);

INSERT INTO "projects" ("name", "slug", "description", "created_at") VALUES
                                                                         ('Project1', 'project1', 'Description for Project1', CURRENT_TIMESTAMP),
                                                                         ('Project2', 'project2', 'Description for Project2', CURRENT_TIMESTAMP),
                                                                         ('Project3', 'project3', 'Description for Project3', CURRENT_TIMESTAMP);

INSERT INTO "project_hashtags" ("hashtag_id", "project_id") VALUES
                                                                (1, 1),
                                                                (2, 1),
                                                                (2, 2),
                                                                (3, 3);

INSERT INTO "user_projects" ("project_id", "user_id") VALUES
                                                          (1, 1),
                                                          (2, 1),
                                                          (2, 2),
                                                          (3, 3);

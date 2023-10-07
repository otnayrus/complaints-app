CREATE TABLE IF NOT EXISTS "public"."users" (
    "id" SERIAL PRIMARY KEY,
    "name" TEXT NOT NULL,
    "password" TEXT NOT NULL,
    "email" TEXT NOT NULL,
    "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
    "updated_at" DATETIME
);

CREATE UNIQUE INDEX IF NOT EXISTS "idx_users_email" ON "users" ("email");

CREATE TABLE IF NOT EXISTS "public"."roles" (
    "id" SERIAL PRIMARY KEY,
    "name" TEXT NOT NULL,
    "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
    "updated_at" DATETIME
);

INSERT INTO "roles" ("id", "name") VALUES (1, "admin");
INSERT INTO "roles" ("id", "name") VALUES (2, "user");

CREATE TABLE IF NOT EXISTS "public"."users_roles" (
    "id" SERIAL PRIMARY KEY,
    "user_id" INTEGER NOT NULL,
    "role_id" INTEGER NOT NULL,
    "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
    "updated_at" DATETIME,

    FOREIGN KEY("user_id") REFERENCES "users"("id"),
    FOREIGN KEY("role_id") REFERENCES "roles"("id")
);

CREATE UNIQUE INDEX IF NOT EXISTS "idx_users_roles_user_id_role_id" ON "users_roles" ("user_id", "role_id");

CREATE TABLE IF NOT EXISTS "public"."categories" (
    "id" SERIAL PRIMARY KEY,
    "name" TEXT NOT NULL,
    "extra_fields" JSON NOT NULL,
    "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
    "updated_at" DATETIME
);

CREATE UNIQUE INDEX IF NOT EXISTS "idx_categories_name" ON "categories" ("name");

CREATE TABLE IF NOT EXISTS "public"."complaints" (
    "id" SERIAL PRIMARY KEY,
    "category_id" INTEGER NOT NULL,
    "description" TEXT NOT NULL,
    "status" INTEGER NOT NULL,
    "remarks" TEXT NOT NULL,
    "extra_fields" JSON NOT NULL,
    "created_by" INTEGER NOT NULL,
    "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
    "updated_at" DATETIME,

    FOREIGN KEY("category_id") REFERENCES "categories"("id") ON DELETE RESTRICT,
    FOREIGN KEY("created_by") REFERENCES "users"("id")
);


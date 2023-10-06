CREATE TABLE IF NOT EXISTS 'users' (
    'id' INTEGER PRIMARY KEY AUTOINCREMENT,
    'name' TEXT NOT NULL,
    'password' TEXT NOT NULL,
    'email' TEXT NOT NULL,
    'created_at' DATETIME DEFAULT CURRENT_TIMESTAMP
    'updated_at' DATETIME
);

CREATE UNIQUE INDEX IF NOT EXISTS 'idx_users_email' ON 'users' ('email');

CREATE TABLE IF NOT EXISTS 'roles' (
    'id' INTEGER PRIMARY KEY AUTOINCREMENT,
    'name' TEXT NOT NULL,
    'created_at' DATETIME DEFAULT CURRENT_TIMESTAMP
    'updated_at' DATETIME
);

CREATE TABLE IF NOT EXISTS 'users_roles' (
    'id' INTEGER PRIMARY KEY AUTOINCREMENT,
    'user_id' INTEGER NOT NULL,
    'role_id' INTEGER NOT NULL,
    'created_at' DATETIME DEFAULT CURRENT_TIMESTAMP
    'updated_at' DATETIME

    FOREIGN KEY('user_id') REFERENCES 'users'('id'),
    FOREIGN KEY('role_id') REFERENCES 'roles'('id')
);

CREATE UNIQUE INDEX IF NOT EXISTS 'idx_users_roles_user_id_role_id' ON 'users_roles' ('user_id', 'role_id');

CREATE TABLE IF NOT EXISTS 'categories' (
    'id' INTEGER PRIMARY KEY AUTOINCREMENT,
    'name' TEXT NOT NULL,
    'extra_fields' JSON NOT NULL,
    'created_at' DATETIME DEFAULT CURRENT_TIMESTAMP
    'updated_at' DATETIME
);

CREATE TABLE IF NOT EXISTS 'complaints' (
    'id' INTEGER PRIMARY KEY AUTOINCREMENT,
    'category_id' INTEGER NOT NULL,
    'description' TEXT NOT NULL,
    'status' INTEGER NOT NULL,
    'remarks' TEXT NOT NULL,
    'extra_fields' JSON NOT NULL,
    'created_by' INTEGER NOT NULL,
    'created_at' DATETIME DEFAULT CURRENT_TIMESTAMP
    'updated_at' DATETIME

    FOREIGN KEY('category_id') REFERENCES 'categories'('id'),
    FOREIGN KEY('created_by') REFERENCES 'users'('id')
);


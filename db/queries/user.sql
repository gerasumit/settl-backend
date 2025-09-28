-- name: CreateUser :one
INSERT INTO users (
    first_name,
    last_name,
    email_id,
    password_hash,
    password_salt,
    auth_provider,
    is_active,
    preferred_currency
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
) RETURNING *;


--   "id" bigserial PRIMARY KEY,
--   "first_name" varchar NOT NULL,
--   "last_name" varchar NOT NULL,
--   "email_id" varchar UNIQUE NOT NULL,
--   "phone" varchar,
--   "username" varchar,
--   "profile_picture_url" varchar,
--   "password_hash" varchar NOT NULL,
--   "password_salt" varchar NOT NULL,
--   "auth_provider" varchar NOT NULL,
--   "last_login_at" timestamptz NOT NULL,
--   "is_active" boolean NOT NULL,
--   "preferred_currency" varchar NOT NULL,
--   "created_at" timestamptz NOT NULL DEFAULT (now()),
--   "updated_at" timestamptz

CREATE TYPE "split_status" AS ENUM (
  'pending',
  'settled',
  'partial'
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "email_id" varchar UNIQUE NOT NULL,
  "phone" varchar,
  "username" varchar,
  "profile_picture_url" varchar,
  "password_hash" varchar NOT NULL,
  "password_salt" varchar NOT NULL,
  "auth_provider" varchar NOT NULL,
  "last_login_at" timestamptz NOT NULL,
  "is_active" boolean NOT NULL,
  "preferred_currency" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz
);

CREATE TABLE "groups" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz
);

CREATE TABLE "user_groups" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "group_id" bigint NOT NULL,
  "role" varchar NOT NULL,
  "joined_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "expenses" (
  "id" bigserial PRIMARY KEY,
  "group_id" bigint NOT NULL,
  "title" varchar NOT NULL,
  "amount" numeric NOT NULL,
  "currency" varchar NOT NULL,
  "notes" text,
  "date_incurred" timestamptz,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz
);

CREATE TABLE "expense_payments" (
  "id" bigserial PRIMARY KEY,
  "expense_id" bigint NOT NULL,
  "user_id" bigint NOT NULL,
  "amount" numeric NOT NULL
);

CREATE TABLE "expense_splits" (
  "id" bigserial PRIMARY KEY,
  "expense_id" bigint NOT NULL,
  "user_id" bigint NOT NULL,
  "share" numeric NOT NULL,
  "status" split_status NOT NULL DEFAULT 'pending'
);

CREATE UNIQUE INDEX ON "user_groups" ("user_id", "group_id");

ALTER TABLE "user_groups" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "user_groups" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("id") ON DELETE CASCADE;

ALTER TABLE "expenses" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("id") ON DELETE CASCADE;

ALTER TABLE "expense_payments" ADD FOREIGN KEY ("expense_id") REFERENCES "expenses" ("id") ON DELETE CASCADE;

ALTER TABLE "expense_payments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "expense_splits" ADD FOREIGN KEY ("expense_id") REFERENCES "expenses" ("id") ON DELETE CASCADE;

ALTER TABLE "expense_splits" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

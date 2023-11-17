CREATE TABLE "accounts" (
  "username" varchar PRIMARY KEY,
  "role" varchar NOT NULL DEFAULT('customer'),
  "email" varchar NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "users" (
  "user_id" bigserial PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "phone_number" varchar UNIQUE NOT NULL,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "image" varchar
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("email") REFERENCES "users" ("email");
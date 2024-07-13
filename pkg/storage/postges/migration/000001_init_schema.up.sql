CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "external_id" varchar NOT NULL,
  "username" varchar NOT NULL,
  "role" int NOT NULL,
  "created_at" timestamp NOT NULL
);

CREATE TABLE "workouts" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "parameters" jsonb NOT NULL,
  "user_id" bigint NOT NULL,
  "created_at" timestamp NOT NULL
);

CREATE TABLE "stats" (
  "id" bigserial PRIMARY KEY,
  "workout_id" bigint NOT NULL,
  "unit" varchar NOT NULL,
  "current" int NOT NULL,
  "reps" int NOT NULL,
  "sets" int NOT NULL,
  "created_at" timestamp NOT NULL
);

CREATE TABLE "replics" (
  "id" bigserial PRIMARY KEY,
  "label" varchar UNIQUE NOT NULL,
  "ru_value" text NOT NULL,
  "en_value" text NOT NULL
);

CREATE INDEX ON "users" ("external_id");

CREATE INDEX ON "workouts" ("user_id");

CREATE INDEX ON "stats" ("workout_id");

CREATE INDEX ON "replics" ("label");

COMMENT ON COLUMN "workouts"."parameters" IS 'sets, reps and others';

ALTER TABLE "workouts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "stats" ADD FOREIGN KEY ("workout_id") REFERENCES "workouts" ("id");
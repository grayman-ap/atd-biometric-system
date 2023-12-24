CREATE TABLE "student" (
  "student_id" varchar PRIMARY KEY NOT NULL,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "tutor" (
  "staff_id" varchar PRIMARY KEY NOT NULL,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "department" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "department_name" varchar NOT NULL,
  "student" varchar NOT NULL,
  "tutor" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "course" (
  "id" bigserial NOT NULL,
  "student" varchar NOT NULL,
  "tutor"  varchar NOT NULL,
  "course_code" varchar PRIMARY KEY NOT NULL,
  "location" varchar NOT NULL,
  "duration" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "attendance" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "student" varchar NOT NULL,
  "course_code" varchar NOT NULL,
  "mark_student" bool NOT NULL,
  "last_attendance" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "student" ("student_id");

CREATE INDEX ON "tutor" ("staff_id");

CREATE INDEX ON "department" ("department_name");

CREATE INDEX ON "course" ("course_code");

CREATE INDEX ON "attendance" ("id");

ALTER TABLE "department" ADD FOREIGN KEY ("student") REFERENCES "student" ("student_id");

ALTER TABLE "department" ADD FOREIGN KEY ("tutor") REFERENCES "tutor" ("staff_id");

ALTER TABLE "course" ADD FOREIGN KEY ("student") REFERENCES "student" ("student_id");

ALTER TABLE "course" ADD FOREIGN KEY ("tutor") REFERENCES "tutor" ("staff_id");

ALTER TABLE "attendance" ADD FOREIGN KEY ("student") REFERENCES "student" ("student_id");

ALTER TABLE "attendance" ADD FOREIGN KEY ("course_code") REFERENCES "course" ("course_code");

CREATE TABLE "student" (
  "student_id" varchar PRIMARY KEY NOT NULL,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "department" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "tutor" (
  "staff_id" varchar PRIMARY KEY NOT NULL,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "department" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "department" (
  "department_id" varchar PRIMARY KEY NOT NULL,
  "school" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "course" (
  "course_code" varchar PRIMARY KEY NOT NULL,
  "department" varchar NOT NULL,
  "number_of_student" bigint NOT NULL,
  "course_title" varchar NOT NULL,
  "course_unit" varchar NOT NULL,
  "venue" varchar NOT NULL,
  "start_time"timestamptz NOT NULL DEFAULT (now()),
  "end_time" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "total_duration" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
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

CREATE INDEX ON "department" ("department_id");

CREATE INDEX ON "course" ("course_code");

CREATE INDEX ON "attendance" ("id");

ALTER TABLE "student" ADD FOREIGN KEY ("department") REFERENCES "department" ("department_id");

ALTER TABLE "tutor" ADD FOREIGN KEY ("department") REFERENCES "department" ("department_id");

ALTER TABLE "course" ADD FOREIGN KEY ("department") REFERENCES "department" ("department_id");

ALTER TABLE "attendance" ADD FOREIGN KEY ("student") REFERENCES "student" ("student_id");

ALTER TABLE "attendance" ADD FOREIGN KEY ("course_code") REFERENCES "course" ("course_code");

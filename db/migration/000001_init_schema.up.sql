CREATE TYPE "WorkerType" AS ENUM (
  'JANITOR',
  'COLLECTOR'
);

CREATE TYPE "GenderType" AS ENUM (
  'MALE',
  'FEMALE'
);

CREATE TYPE "WorkerStatusType" AS ENUM (
  'AVAILABLE',
  'WORKING'
);

CREATE TYPE "VehicleStatusType" AS ENUM (
  'AVAILABLE',
  'USINGE'
);

CREATE TYPE "TasksStatus" AS ENUM (
  'OPENED',
  'DONE'
);

CREATE TABLE "BackOfficers" (
  "id" BIGSERIAL PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "ssn" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "name" varchar NOT NULL,
  "phone" varchar NOT NULL,
  "age" int NOT NULL,
  "gender" GenderType NOT NULL,
  "date_of_birth" timestamp NOT NULL,
  "place_of_birth" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "Workers" (
  "id" BIGSERIAL PRIMARY KEY,
  "ssn" varchar UNIQUE NOT NULL,
  "name" varchar NOT NULL,
  "phone" varchar NOT NULL,
  "age" int NOT NULL,
  "worker_type" WorkerType NOT NULL,
  "gender" GenderType NOT NULL,
  "date_of_birth" timestamp NOT NULL,
  "place_of_birth" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "WorkerStatus" (
  "id" BIGSERIAL PRIMARY KEY,
  "worker_id" varchar UNIQUE NOT NULL,
  "status" WorkerStatusType NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "Vehicles" (
  "id" BIGSERIAL PRIMARY KEY,
  "make_by" varchar NOT NULL,
  "model" varchar NOT NULL,
  "capacity" varchar NOT NULL,
  "fuel_consumption" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "VehicleStatus" (
  "id" BIGSERIAL PRIMARY KEY,
  "vehicle_id" varchar UNIQUE NOT NULL,
  "status" VehicleStatusType NOT NULL,
  "current_fuel" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "MCPs" (
  "id" BIGSERIAL PRIMARY KEY,
  "location" varchar NOT NULL,
  "capacity" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "MCPStatus" (
  "id" BIGSERIAL PRIMARY KEY,
  "mcp_id" varchar UNIQUE NOT NULL,
  "location" varchar NOT NULL,
  "capacity" varchar NOT NULL,
  "current_level_fill" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "Routes" (
  "id" BIGSERIAL PRIMARY KEY,
  "start_location" varchar NOT NULL,
  "end_location" varchar NOT NULL,
  "distance" varchar NOT NULL,
  "estimated_time" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "Tasks" (
  "id" BIGSERIAL PRIMARY KEY,
  "start_time" timestamp NOT NULL,
  "end_time" timestamp NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "worker_id" int NOT NULL,
  "vehicle_id" int NOT NULL,
  "mcp_id" int NOT NULL,
  "route_id" int NOT NULL,
  "status" TasksStatus NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "Workers" ADD FOREIGN KEY ("id") REFERENCES "WorkerStatus" ("worker_id");

ALTER TABLE "Vehicles" ADD FOREIGN KEY ("id") REFERENCES "VehicleStatus" ("vehicle_id");

ALTER TABLE "MCPs" ADD FOREIGN KEY ("id") REFERENCES "MCPStatus" ("mcp_id");

ALTER TABLE "Workers" ADD FOREIGN KEY ("id") REFERENCES "Tasks" ("worker_id");

ALTER TABLE "Vehicles" ADD FOREIGN KEY ("id") REFERENCES "Tasks" ("vehicle_id");

ALTER TABLE "MCPs" ADD FOREIGN KEY ("id") REFERENCES "Tasks" ("mcp_id");

ALTER TABLE "Routes" ADD FOREIGN KEY ("id") REFERENCES "Tasks" ("route_id");

CREATE TABLE "user_credential" (
    "id" uuid PRIMARY KEY,
    "username" varchar UNIQUE NOT NULL,
    "hashed_password" varchar NOT NULL,
    "email" varchar UNIQUE NOT NULL,
    "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
    "is_admin" bool NOT NULL DEFAULT False,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "user_info" (
    "id" uuid PRIMARY KEY,
    "user_id" uuid NOT NULL,
    "phone_number" varchar UNIQUE NOT NULL,
    "first_name" varchar NOT NULL,
    "last_name" varchar NOT NULL,
    "middle_name" varchar NOT NULL,
    "updated_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "user_address" (
    "id" uuid PRIMARY KEY,
    owner uuid NOT NULL,
    "address_name" varchar NOT NULL,
    "address" varchar NOT NULL,
    "city" varchar NOT NULL,
    "state" varchar NOT NULL,
    "zipcode" int NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "user_info" ("id");

CREATE INDEX ON "user_info" ("user_id");

CREATE INDEX ON "user_address" ("id");

CREATE INDEX ON "user_address" ("owner");

CREATE UNIQUE INDEX ON "user_address" ("owner", "address_name");

ALTER TABLE "user_info" ADD FOREIGN KEY ("user_id") REFERENCES "user_credential" ("id");

ALTER TABLE "user_address" ADD FOREIGN KEY ("owner") REFERENCES "user_info" ("id");

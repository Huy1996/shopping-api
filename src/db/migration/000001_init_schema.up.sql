CREATE TABLE "user_credential" (
                                   "id" uuid PRIMARY KEY,
                                   "username" varchar UNIQUE NOT NULL,
                                   "hashed_password" varchar NOT NULL,
                                   "email" varchar UNIQUE NOT NULL,
                                   "password_changed_at" timestamptz NOT NULL DEFAULT 'now()',
                                   "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "user_info" (
                             "id" uuid PRIMARY KEY,
                             "user_id" uuid NOT NULL,
                             "phone_number" varchar UNIQUE NOT NULL,
                             "first_name" varchar NOT NULL,
                             "last_name" varchar NOT NULL,
                             "middle_name" varchar NOT NULL,
                             "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "address_book" (
                                "id" uuid PRIMARY KEY,
                                "user" uuid NOT NULL,
                                "address_name" varchar NOT NULL,
                                "address" varchar NOT NULL,
                                "city" varchar NOT NULL,
                                "state" varchar NOT NULL,
                                "zipcode" int NOT NULL
);

CREATE INDEX ON "user_info" ("id");

CREATE INDEX ON "user_info" ("user_id");

CREATE INDEX ON "address_book" ("id");

CREATE INDEX ON "address_book" ("user");

CREATE UNIQUE INDEX ON "address_book" ("user", "address_name");

ALTER TABLE "user_info" ADD FOREIGN KEY ("user_id") REFERENCES "user_credential" ("id");

ALTER TABLE "address_book" ADD FOREIGN KEY ("user") REFERENCES "user_info" ("id");

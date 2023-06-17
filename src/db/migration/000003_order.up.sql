CREATE TYPE "order_status" AS ENUM (
  'Preparing',
  'Shipped',
  'Delivered',
  'Pending Cancel',
  'Canceled'
);

CREATE TYPE "payment_status" AS ENUM (
  'Pending',
  'Succeed',
  'Rejected'
);

CREATE TYPE "payment_type" AS ENUM (
  'VISA',
  'MASTER CARD',
  'AMERICAN EXPRESS'
);

CREATE TABLE "order_detail" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid NOT NULL,
  "total" float8 NOT NULL,
  "payment_id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE TABLE "order_item" (
  "id" uuid PRIMARY KEY,
  "order_id" uuid NOT NULL,
  "product_id" uuid NOT NULL,
  "quantity" int NOT NULL,
  "status" order_status,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE TABLE "payment_detail" (
  "id" uuid PRIMARY KEY,
  "amount" float8 NOT NULL,
  "type" payment_type NOT NULL,
  "status" payment_status NOT NULL,
  "card_number" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

ALTER TABLE "order_detail" ADD FOREIGN KEY ("user_id") REFERENCES "user_info" ("id");

ALTER TABLE "order_detail" ADD FOREIGN KEY ("payment_id") REFERENCES "payment_detail" ("id");

ALTER TABLE "order_item" ADD FOREIGN KEY ("order_id") REFERENCES "order_detail" ("id");

ALTER TABLE "order_item" ADD FOREIGN KEY ("product_id") REFERENCES "product" ("id");

CREATE TABLE "user_cart" (
  "id" uuid PRIMARY KEY,
  "owner" uuid NOT NULL
);

CREATE TABLE "cart_item" (
  "id" uuid PRIMARY KEY,
  "cart_id" uuid NOT NULL,
  "product_id" uuid NOT NULL,
  "quantity" int NOT NULL
);

CREATE TABLE "product" (
  "id" uuid PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" text NOT NULL,
  "SKU" varchar NOT NULL,
  "price" decimal NOT NULL,
  "category_id" uuid NOT NULL,
  "inventory_id" uuid NOT NULL,
  "discount_id" uuid
);

CREATE TABLE "product_inventory" (
  "id" uuid PRIMARY KEY,
  "quantity" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE TABLE "product_category" (
  "id" uuid PRIMARY KEY,
  "name" varchar NOT NULL UNIQUE ,
  "description" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE TABLE "product_discount" (
  "id" uuid PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" text NOT NULL,
  "discount_percent" decimal NOT NULL,
  "active" bool NOT NULL DEFAULT False,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);


COMMENT ON COLUMN "cart_item"."quantity" IS 'Cannot be less than 1';

COMMENT ON COLUMN "product"."price" IS 'Cannot be negative';

ALTER TABLE "user_cart" ADD FOREIGN KEY ("owner") REFERENCES "user_info" ("id");

ALTER TABLE "cart_item" ADD FOREIGN KEY ("cart_id") REFERENCES "user_cart" ("id");

ALTER TABLE "cart_item" ADD FOREIGN KEY ("product_id") REFERENCES "product" ("id");

ALTER TABLE "product" ADD FOREIGN KEY ("category_id") REFERENCES "product_category" ("id");

ALTER TABLE "product" ADD FOREIGN KEY ("inventory_id") REFERENCES "product_inventory" ("id");

ALTER TABLE "product" ADD FOREIGN KEY ("discount_id") REFERENCES "product_discount" ("id");

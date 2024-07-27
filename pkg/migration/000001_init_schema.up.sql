CREATE TABLE "accounts" (
  "username" varchar PRIMARY KEY,
  "role" varchar NOT NULL DEFAULT('customer'),
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "type" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now() at time zone 'utc')
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "phone_number" varchar UNIQUE NOT NULL,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "image" varchar
);

CREATE TABLE "payments" (
  "id" bigserial PRIMARY KEY,
  "type" varchar NOT NULL,
  "card_number" varchar NOT NULL,
  "cvc_code" varchar NOT NULL,
  "owner" varchar NOT NULL
);

CREATE TABLE "addresses" (
  "id" bigserial PRIMARY KEY,
  "city" varchar NOT NULL,
  "ward" varchar NOT NULL,
  "district" varchar NOT NULL,
  "street" varchar NOT NULL,
  "uuid" varchar UNIQUE NOT NULL
);

CREATE TABLE "comments" (
  "id" bigserial PRIMARY KEY,
  "level" bigint NOT NULL DEFAULT('0'),
  "content" varchar NOT NULL,
  "user_id" bigint NOT NULL,
  "product_id" bigint NOT NULL,
  "rating" int
);

CREATE TABLE "products" (
  "id" bigserial PRIMARY KEY,
  "image" varchar NOT NULL,
  "price" varchar NOT NULL,
  "description" varchar NOT NULL,
  "name" varchar NOT NULL,
  "supplier_id" bigint NOT NULL,
  "category_id" bigint NOT NULL,
  "created" timestamp default (now() at time zone 'utc'),
  "spec" jsonb,
  "status" varchar NOT NULL
);

CREATE TABLE "suppliers" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL
);

CREATE TABLE "categories" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" varchar NOT NULL
);

CREATE TABLE "carts" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "inventory_id" bigint NOT NULL,
  "quantity" bigint NOT NULL
);

CREATE TABLE "orders" (
  "id" bigserial PRIMARY KEY,
  "uuid" varchar NOT NULL,
  "time" timestamp default (now() at time zone 'utc'),
  "user_id" bigint NOT NULL,
  "total_amount" NUMERIC(19, 4) NOT NULL,
  "status" varchar NOT NULL
);

CREATE TABLE "product_in_order" (
  "id" bigserial PRIMARY KEY,
  "order_id" bigint NOT NULL,
  "inventory_id" bigint NOT NULL,
  "quantity" bigint NOT NULL,
  "total_amount" NUMERIC(19, 4) NOT NULL,
  "currency_code" varchar(3) NOT NULL
);

CREATE TABLE "deliveries" (
  "id" bigserial PRIMARY KEY,
  "sent_date" timestamptz NOT NULL,
  "received_date" timestamptz NOT NULL,
  "method" varchar NOT NULL,
  "note" varchar
);

CREATE TABLE "receipts" (
  "id" bigserial PRIMARY KEY,
  "time" timestamptz NOT NULL,
  "delivery_id" bigint NOT NULL,
  "payment_id" bigint NOT NULL
);

CREATE TABLE "user_address" (
  "user_id" bigint PRIMARY KEY,
  "address_uuid" varchar NOT NULL
);

CREATE TABLE "suppliers_address" (
  "suppliers_id" bigint PRIMARY KEY,
  "address_uuid" varchar NOT NULL
);

CREATE TABLE "favorite_product" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "inventory_id" bigint NOT NULL
);

CREATE TABLE "inventories" (
  "id" bigserial PRIMARY KEY,
  "product_id" int NOT NULL,
  "price" NUMERIC(19, 4) NOT NULL,
  "status" varchar NOT NULL,
  "currency_code" varchar(3) NOT NULL,
  "specs" jsonb,
  "available" int NOT NULL,
  "color" varchar NOT NULL,
  "color_img" varchar,
  "image" varchar
);

CREATE TABLE "collections" (
  "id" bigserial PRIMARY KEY,
  "created" timestamp default (timezone('utc', now())),
  "position" varchar NOT NULL ,
  "headline" varchar,
  "body" jsonb
);

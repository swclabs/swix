CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
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

CREATE TABLE "addresses" (
  "id" bigserial PRIMARY KEY,
  "city" varchar NOT NULL,
  "ward" varchar NOT NULL,
  "district" varchar NOT NULL,
  "street" varchar NOT NULL,
  "user_id" bigint NOT NULL
);

CREATE TABLE "comments" (
  "id" bigserial PRIMARY KEY,
  "level" bigint NOT NULL DEFAULT('0'),
  "content" varchar NOT NULL,
  "user_id" bigint NOT NULL,
  "product_id" bigint NOT NULL,
  "rating" int,
  "liked" int,
  "disliked" int,
  "parent_id" bigint
);

CREATE TABLE "products" (
  "id" bigserial PRIMARY KEY,
  "image" varchar NOT NULL,
  "shop_image" varchar,
  "price" varchar NOT NULL,
  "description" varchar NOT NULL,
  "name" varchar NOT NULL,
  "supplier_id" bigint NOT NULL,
  "category_id" bigint NOT NULL,
  "created" timestamp default (now() at time zone 'utc'),
  "specs" jsonb,
  "status" varchar NOT NULL,
  "rating" NUMERIC(3, 1) CHECK ("rating" = -1.0 OR ("rating" >= 0 AND "rating" <= 5)) DEFAULT -1.0
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
  "quantity" bigint NOT NULL,
  CONSTRAINT unique_inventory UNIQUE (inventory_id, user_id)
);

CREATE TABLE "orders" (
  "id" bigserial PRIMARY KEY,
  "uuid" varchar NOT NULL,
  "time" timestamp default (now() at time zone 'utc'),
  "user_id" bigint NOT NULL,
  "delivery_id" bigint NOT NULL,
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
  "user_id" bigint NOT NULL,
  "address_id" bigint NOT NULL,
  "sent_date" timestamptz,
  "status" varchar NOT NULL,
  "method" varchar NOT NULL,
  "note" varchar
);

CREATE TABLE "favorite" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "inventory_id" bigint NOT NULL,
  CONSTRAINT unique_favorite UNIQUE (user_id, inventory_id)
);

CREATE TABLE "inventories" (
  "id" bigserial PRIMARY KEY,
  "product_id" int NOT NULL,
  "price" NUMERIC(19, 4) NOT NULL,
  "status" varchar NOT NULL,
  "currency_code" varchar(3) NOT NULL,
  "available" int NOT NULL,
  "color" varchar NOT NULL,
  "color_img" varchar,
  "image" varchar,
  "specs" jsonb
);

CREATE TABLE "news" (
  "id" bigserial PRIMARY KEY,
  "created" timestamp default (timezone('utc', now())),
  "category" varchar NOT NULL,
  "header" varchar,
  "body" jsonb
);

CREATE TABLE "stars" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint UNIQUE NOT NULL,
  "product_id" bigint UNIQUE NOT NULL,
  CONSTRAINT unique_stars UNIQUE (user_id, product_id)
);

CREATE TABLE "coupons" (
  "id" bigserial PRIMARY KEY,
  "code" varchar UNIQUE NOT NULL,
  "discount" NUMERIC(19, 4) NOT NULL,
  "status" varchar NOT NULL,
  "used" int NOT NULL,
  "max_use" int NOT NULL,
  "description" varchar NOT NULL,
  "expired_at" timestamptz NOT NULL
);

CREATE TABLE "coupons_used" (
  "id" bigserial PRIMARY KEY,
  "coupon_code" varchar NOT NULL,
  "order_id" bigint NOT NULL,
  "user_id" bigint NOT NULL,
  "used_at" timestamptz default (timezone('utc', now())),
  CONSTRAINT unique_coupons_used UNIQUE (user_id, coupon_code)
);

CREATE TABLE "province" (
  "pid" bigserial PRIMARY KEY,
  "id" varchar UNIQUE,
  "name" varchar
);

CREATE TABLE "district" (
  "pid" bigserial PRIMARY KEY,
  "id" varchar UNIQUE,
  "province_id" varchar,
  "name" varchar
);

CREATE TABLE "commune" (
  "pid" bigserial PRIMARY KEY,
  "id" varchar,
  "district_id" varchar ,
  "name" varchar
);
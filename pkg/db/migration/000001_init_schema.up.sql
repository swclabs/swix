CREATE TABLE "accounts" (
  "username" varchar PRIMARY KEY,
  "role" varchar NOT NULL DEFAULT('customer'),
  "email" varchar NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
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
  "street" varchar NOT NULL
);

CREATE TABLE "comments" (
  "id" bigserial PRIMARY KEY,
  "level" varchar NOT NULL DEFAULT('0'),
  "content" varchar NOT NULL,
  "user_id" bigint NOT NULL,
  "product_id" bigint NOT NULL
);

CREATE TABLE "products" (
  "id" bigserial PRIMARY KEY,
  "image" varchar NOT NULL,
  "price" varchar NOT NULL,
  "description" varchar NOT NULL,
  "name" varchar NOT NULL,
  "supplier_id" bigint NOT NULL,
  "category_id" bigint NOT NULL
);

CREATE TABLE "suppliers" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "phone_number" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL
);

CREATE TABLE "categories" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" varchar NOT NULL
);

CREATE TABLE "carts" (
  "id" bigserial PRIMARY KEY,
  "quantity" bigint NOT NULL,
  "total_price" bigint NOT NULL,
  "user_id" bigint NOT NULL
);

CREATE TABLE "product_in_cart" (
  "id" bigserial PRIMARY KEY,
  "cart_id" bigint NOT NULL,
  "product_id" bigint NOT NULL,
  "amount" bigint NOT NULL
);

CREATE TABLE "orders" (
  "id" bigserial PRIMARY KEY,
  "time" timestamptz NOT NULL,
  "user_id" bigint NOT NULL,
  "total_price" bigint NOT NULL,
  "status" varchar NOT NULL,
  "quantity" bigint NOT NULL
);

CREATE TABLE "product_in_order" (
  "id" bigserial PRIMARY KEY,
  "order_id" bigint NOT NULL,
  "product_id" bigint NOT NULL,
  "amount" bigint NOT NULL
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
  "address_id" bigint NOT NULL
);

CREATE TABLE "supplier_address" (
  "supplier_id" bigint PRIMARY KEY,
  "address_id" bigint NOT NULL
);


ALTER TABLE "accounts" ADD FOREIGN KEY ("email") REFERENCES "users" ("email");

ALTER TABLE "payments" ADD FOREIGN KEY ("owner") REFERENCES "users" ("email");

ALTER TABLE "comments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "comments" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id")

ALTER TABLE "products" ADD FOREIGN KEY ("supplier_id") REFERENCES "suppliers" ("id");
ALTER TABLE "products" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "receipts" ADD FOREIGN KEY ("delivery_id") REFERENCES "deliveries" ("id");
ALTER TABLE "receipts" ADD FOREIGN KEY ("id") REFERENCES "orders" ("id");
ALTER TABLE "receipts" ADD FOREIGN KEY ("payment_id") REFERENCES "payments" ("id");

ALTER TABLE "carts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_address" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "user_address" ADD FOREIGN KEY ("address_id") REFERENCES "addresses" ("id");

ALTER TABLE "supplier_address" ADD FOREIGN KEY ("address_id") REFERENCES "addresses" ("id");
ALTER TABLE "supplier_address" ADD FOREIGN KEY ("supplier_id") REFERENCES "suppliers" ("id");

ALTER TABLE "product_in_cart" ADD FOREIGN KEY ("cart_id") REFERENCES "carts" ("id");
ALTER TABLE "product_in_cart" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "product_in_order" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");
ALTER TABLE "product_in_order" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

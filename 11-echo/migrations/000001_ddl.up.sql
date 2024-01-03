CREATE TABLE "users" (
  "user_id" SERIAL PRIMARY KEY,
  "username" varchar NOT NULL,
  "password" varchar NOT NULL,
  "deposit_amount" numeric(11,2)
);

CREATE TABLE "products" (
  "product_id" SERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "stock" int NOT NULL,
  "price" numeric(11,2)
);

CREATE TABLE "transactions" (
  "transaction_id" SERIAL PRIMARY KEY,
  "user_id" int NOT NULL,
  "product_id" int NOT NULL,
  "quantity" int NOT NULL,
  "total_amount" numeric(11,2)
);

ALTER TABLE "transactions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("product_id");
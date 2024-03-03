CREATE TYPE "transfer_type" AS ENUM (
  'crypto',
  'money'
);

CREATE TABLE "users" (
    "id" serial PRIMARY KEY,
    "name" varchar NOT NULL,
    "surname" varchar NOT NULL,
    "cpf" varchar(11) UNIQUE NOT NULL,
    "password" varchar NOT NULL,
    "email" varchar UNIQUE NOT NULL,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "account" (
    "id" serial PRIMARY KEY,
    "account_address" varchar UNIQUE NOT NULL,
    "balance" bigint NOT NULL DEFAULT 0,
    "balance_btc" bigint NOT NULL DEFAULT 0,
    "credit" bigint NOT NULL DEFAULT 0,
    "user_id" serial NOT NULL,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "deposit" (
    "id" bigserial PRIMARY KEY,
    "deposit_uuid" varchar(36) UNIQUE NOT NULL,
    "amount" bigint NOT NULL,
    "account_address" varchar NOT NULL,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "transfer" (
    "id" bigserial PRIMARY KEY,
    "transfer_uuid" varchar(36) UNIQUE NOT NULL,
    "origin_address" varchar NOT NULL,
    "destination_address" varchar NOT NULL,
    "amount" bigint NOT NULL,
    "type" transfer_type NOT NULL,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "address" (
    "id" serial PRIMARY KEY,
    "user_id" integer NOT NULL,
    "street" varchar NOT NULL,
    "neighbor" varchar NOT NULL,
    "number" varchar NOT NULL,
    "comp" varchar NOT NULL,
    "cep" varchar NOT NULL,
    "state" varchar(2) NOT NULL,
    "city" varchar NOT NULL,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "card" (
    "id" bigserial PRIMARY KEY,
    "account_id" serial NOT NULL,
    "name" varchar NOT NULL,
    "number" varchar UNIQUE NOT NULL,
    "cvv" varchar NOT NULL,
    "exp" varchar(7) NOT NULL,
    "limit" bigint NOT NULL,
    "used_limit" bigint NOT NULL,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX ON "users" ("id");

CREATE INDEX ON "users" ("cpf");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "users" ("created_at");

CREATE INDEX ON "account" ("id");

CREATE INDEX ON "account" ("user_id");

CREATE INDEX ON "account" ("created_at");

CREATE INDEX ON "deposit" ("id");

CREATE INDEX ON "deposit" ("deposit_uuid");

CREATE INDEX ON "deposit" ("amount");

CREATE INDEX ON "deposit" ("account_address");

CREATE INDEX ON "deposit" ("created_at");

CREATE INDEX ON "transfer" ("id");

CREATE INDEX ON "transfer" ("origin_address");

CREATE INDEX ON "transfer" ("destination_address");

CREATE INDEX ON "transfer" ("destination_address", "origin_address");

CREATE INDEX ON "transfer" ("type");

CREATE INDEX ON "transfer" ("amount");

CREATE INDEX ON "transfer" ("created_at");

CREATE INDEX ON "address" ("id");

CREATE INDEX ON "address" ("user_id");

CREATE INDEX ON "address" ("neighbor");

CREATE INDEX ON "address" ("city");

CREATE INDEX ON "address" ("state");

CREATE INDEX ON "address" ("street");

CREATE INDEX ON "address" ("created_at");

CREATE INDEX ON "card" ("id");

CREATE INDEX ON "card" ("account_id");

ALTER TABLE "account" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "deposit" ADD FOREIGN KEY ("account_address") REFERENCES "account" ("account_address");

ALTER TABLE "transfer" ADD FOREIGN KEY ("origin_address") REFERENCES "account" ("account_address");

ALTER TABLE "transfer" ADD FOREIGN KEY ("destination_address") REFERENCES "account" ("account_address");

ALTER TABLE "address" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "card" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");

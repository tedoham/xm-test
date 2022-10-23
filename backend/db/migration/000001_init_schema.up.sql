CREATE TABLE "company_type_name" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "companies" (
  "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "company_type_id" bigint NOT NULL,
  "name" varchar(15) NOT NULL,
  "description" text,
  "amount" integer NOT NULL,
  "registered" boolean NOT NULL
);

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "email" varchar(225) NOT NULL UNIQUE,
  "password" varchar(225) NOT NULL,
  "username" varchar(255) NOT NULL,
  "tokenhash" varchar(255) NOT NULL,
  "isverified" boolean NOT NULL,
  "createdat" TIMESTAMPTZ DEFAULT Now(),
  "updatedat" TIMESTAMPTZ DEFAULT Now() 
);

ALTER TABLE "companies" ADD FOREIGN KEY ("company_type_id") REFERENCES "company_type_name" ("id");
INSERT INTO company_type_name(name) VALUES ('Corporations'), ('NonProfit'), ('Cooperative'), ('Sole Proprietorship');


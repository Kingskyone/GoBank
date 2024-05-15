CREATE TABLE "users"
(
    "username"           varchar PRIMARY KEY,
    "hashed_password"    varchar        NOT NULL,
    "full_name"          varchar        NOT NULL,
    "email"              varchar UNIQUE NOT NULL,
    "password_change_at" timestamptz    NOT NULL DEFAULT '0001-01-01 00:00:00Z',
    "created_at"         timestamptz    NOT NULL DEFAULT (now())
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

-- CREATE INDEX ON "accounts" ("owner", "currency");    添加唯一约束，同时加上复合索引
ALTER TABLE "accounts" ADD CONSTRAINT "owner_currency_key" UNIQUE ("owner", "currency");
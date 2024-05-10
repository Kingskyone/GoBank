CREATE TABLE "accounts" (
                            "id" bigserial PRIMARY KEY,
                            "owner" varchar NOT NULL,
                            "balance" float NOT NULL,
                            "currency" varchar NOT NULL,
                            "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "ectries" (
                           "id" bigserial PRIMARY KEY,
                           "account_id" bigint NOT NULL,
                           "amount" float NOT NULL,
                           "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfer" (
                            "id" bigserial PRIMARY KEY,
                            "from_account_id" bigint NOT NULL,
                            "to_account_id" bigint NOT NULL,
                            "money" float NOT NULL,
                            "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "accounts" ("owner");

CREATE INDEX ON "ectries" ("account_id");

CREATE INDEX ON "transfer" ("from_account_id");

CREATE INDEX ON "transfer" ("to_account_id");

CREATE INDEX ON "transfer" ("from_account_id", "to_account_id");

COMMENT ON COLUMN "ectries"."amount" IS '可正可负';

COMMENT ON COLUMN "transfer"."money" IS '非负';

ALTER TABLE "ectries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfer" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfer" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");

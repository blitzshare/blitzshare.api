CREATE TABLE "api_keys"
(
    "id"           bigserial PRIMARY KEY,
    "created_at"   timestamp DEFAULT (now()),
    "api_key"      varchar NOT NULL UNIQUE,
    "enabled"      boolean NOT NULL
);

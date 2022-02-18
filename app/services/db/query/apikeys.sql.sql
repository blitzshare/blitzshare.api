
-- name: GetApiKeysByPartner :many
SELECT * FROM "api_keys" WHERE "id" = $1;

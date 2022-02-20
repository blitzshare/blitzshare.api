
-- name: GetApiKey :one
SELECT * FROM "api_keys" WHERE api_key = $1;

-- name: GetApiKeys :many
SELECT * FROM "api_keys";

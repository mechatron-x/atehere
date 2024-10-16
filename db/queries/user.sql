-- name: SaveUser :one
INSERT INTO users (
    id, email, full_name, created_at, updated_at
) VALUES (
    $1, $2, $3, $4, $5
) ON CONFLICT (id) DO UPDATE SET email = $2, full_name = $3, updated_at = NOW()
RETURNING *;
-- name: SaveUser :one
INSERT INTO users (
    id, full_name, birth_date, created_at, updated_at, deleted_at
) VALUES (
    $1, $2, $3, $4, $5, $6
) ON CONFLICT (id) DO UPDATE SET full_name = $2, birth_date = $3, updated_at = NOW()
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id=$1;
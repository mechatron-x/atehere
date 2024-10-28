-- name: SaveManager :one
INSERT INTO managers (
    id, full_name, phone_number, created_at, updated_at, deleted_at
) VALUES (
    $1, $2, $3, $4, $5, $6
) ON CONFLICT (id) DO UPDATE SET full_name = $2, phone_number = $3, updated_at = NOW()
RETURNING *;

-- name: GetManager :one
SELECT * FROM managers
WHERE id=$1;
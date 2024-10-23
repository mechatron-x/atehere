-- name: SaveCustomer :one
INSERT INTO customers (
    id, full_name, birth_date, created_at, updated_at, deleted_at
) VALUES (
    $1, $2, $3, $4, $5, $6
) ON CONFLICT (id) DO UPDATE SET full_name = $2, birth_date = $3, updated_at = NOW()
RETURNING *;

-- name: GetCustomer :one
SELECT * FROM customers
WHERE id=$1;
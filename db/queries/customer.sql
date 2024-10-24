-- name: SaveCustomer :one
INSERT INTO customers (
    id, full_name, gender, birth_date, created_at, updated_at, deleted_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) ON CONFLICT (id) DO UPDATE SET full_name = $2,gender = $3 ,birth_date = $4, updated_at = NOW()
RETURNING *;

-- name: GetCustomer :one
SELECT * FROM customers
WHERE id=$1;
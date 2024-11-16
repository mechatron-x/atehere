-- name: SaveMenu :exec
INSERT INTO menus (
    id,
    restaurant_id,
    category,
    created_at,
    updated_at,
    deleted_at
) VALUES (
    $1, $2, $3, $4, $5, $6 
) ON CONFLICT (id) DO
UPDATE SET
    restaurant_id = $2,
    category = $3,
    updated_at = NOW()
;

-- name: GetMenuByCategory :one
SELECT * FROM menus
WHERE restaurant_id=$1
AND category LIKE $2;

-- name: GetAllMenus :many
SELECT * FROM menus
WHERE restaurant_id=$1
ORDER BY category;
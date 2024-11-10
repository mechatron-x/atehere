-- name: SaveRestaurantTable :exec
INSERT INTO restaurant_tables (
    id,
    restaurant_id,
    name,
    created_at,
    updated_at
) VALUES (
    $1, $2, $3, $4, $5
) ON CONFLICT (id, restaurant_id) DO
UPDATE SET
    name=$3,
    updated_at=NOW()
;

-- name: GetRestaurantTables :many
SELECT * FROM restaurant_tables
WHERE restaurant_id=$1;

-- name: DeleteRestaurantTables :exec
DELETE FROM restaurant_tables
WHERE restaurant_id=$1;
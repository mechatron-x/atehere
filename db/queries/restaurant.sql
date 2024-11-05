-- name: SaveRestaurant :exec
INSERT INTO restaurants (
    id, 
    owner_id, 
    name, 
    foundation_year, 
    phone_number, 
    opening_time, 
    closing_time,
    working_days,
    created_at,
    updated_at,
    deleted_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
) ON CONFLICT (id) DO 
UPDATE SET
    owner_id = $2,
    name = $3,
    foundation_year = $4,
    phone_number = $5,
    opening_time = $6,
    closing_time = $7,
    working_days = $8,
    updated_at = NOW()
;

-- name: GetRestaurant :one
SELECT * FROM restaurants
WHERE id=$1;
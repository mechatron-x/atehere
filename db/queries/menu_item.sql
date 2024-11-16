-- name: SaveMenuItem :exec
INSERT INTO menu_items (
    id,
    menu_id,
    name,
    description,
    image_name,
    price_amount,
    price_currency,
    discount_percentage,
    ingredients,
    created_at,
    updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11 
) ON CONFLICT (id, menu_id) DO
UPDATE SET
    name=$3,
    description=$4,
    image_name=$5,
    price_amount=$6,
    price_currency=$7,
    discount_percentage=$8,
    ingredients=$9,
    updated_at=NOW()
;

-- name: GetMenuItems :many
SELECT * FROM menu_items
WHERE menu_id=$1;

-- name: GetMenuItemByID :one
SELECT * FROM menu_items
WHERE id=$1;
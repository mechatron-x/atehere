CREATE TABLE IF NOT EXISTS menu_items (
    id UUID,
    menu_id UUID, 
    name VARCHAR(100) NOT NULL,
    description VARCHAR(300) NOT NULL,
    image_name VARCHAR(50),
    price_amount DOUBLE PRECISION NOT NULL,
    price_currency VARCHAR(5) NOT NULL,
    discount_percentage SMALLINT NOT NULL,
    ingredients VARCHAR(255)[] NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id, menu_id),
    FOREIGN KEY (menu_id) REFERENCES menus(id)
);
CREATE TABLE IF NOT EXISTS restaurants (
    id UUID,
    owner_id UUID NOT NULL,
    name VARCHAR(100) NOT NULL,
    foundation_year VARCHAR(20),
    phone_number VARCHAR(20),
    opening_time CHAR(5) NOT NULL,
    closing_time CHAR(5) NOT NULL,
    working_days VARCHAR(10)[],
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (owner_id) REFERENCES managers(id)
);
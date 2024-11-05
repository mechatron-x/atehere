CREATE TABLE IF NOT EXISTS customers (
  id UUID,
  full_name VARCHAR(255) NOT NULL,
  gender VARCHAR(50) NOT NULL,
  birth_date TIMESTAMP,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP,
  PRIMARY KEY (id)
);
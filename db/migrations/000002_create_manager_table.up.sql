CREATE TABLE IF NOT EXISTS managers (
  id UUID,
  full_name VARCHAR(255) NOT NULL,
  phone_number VARCHAR(20) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP,
  PRIMARY KEY (id)
);
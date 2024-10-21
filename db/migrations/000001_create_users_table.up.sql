CREATE TABLE IF NOT EXISTS users (
  id CHAR(36),
  full_name VARCHAR(255) NOT NULL,
  birth_date TIMESTAMP,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  deleted_at TIMESTAMP,
  PRIMARY KEY (id)
);
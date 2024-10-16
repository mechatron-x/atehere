CREATE TABLE IF NOT EXISTS users (
  id CHAR(36),
  email VARCHAR(200),
  full_name VARCHAR(255),
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  PRIMARY KEY (id)
);
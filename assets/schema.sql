-- -----------------------------------------------------
-- Table users
-- -----------------------------------------------------
DROP TABLE IF EXISTS users ;

CREATE TABLE IF NOT EXISTS users (
  id SERIAL NOT NULL,
  username VARCHAR(16) NOT NULL,
  email VARCHAR(255) NOT NULL,
  password VARCHAR(32) NOT NULL,
  created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  name VARCHAR(255) NOT NULL,
  status BOOLEAN NULL DEFAULT TRUE,
  pro BOOLEAN NULL DEFAULT FALSE,
  PRIMARY KEY (id)
);

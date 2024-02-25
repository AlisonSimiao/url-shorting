-- -----------------------------------------------------
-- Table users
-- -----------------------------------------------------
DROP TABLE IF EXISTS users ;

CREATE TABLE IF NOT EXISTS users (
  username VARCHAR(16) NOT NULL,
  email VARCHAR(255) NOT NULL,
  password VARCHAR(32) NOT NULL,
  id SERIAL NOT NULL,
  created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  name VARCHAR(255) NOT NULL,
  status BOOLEAN NULL DEFAULT TRUE,
  pro BOOLEAN NULL DEFAULT FALSE,
  PRIMARY KEY (id));


-- -----------------------------------------------------
-- Table users
-- -----------------------------------------------------
DROP TABLE IF EXISTS users ;

CREATE TABLE IF NOT EXISTS users (
  username VARCHAR(16) NOT NULL,
  email VARCHAR(255) NOT NULL,
  password VARCHAR(32) NOT NULL,
  id SERIAL NOT NULL,
  created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  name VARCHAR(255) NOT NULL,
  status BOOLEAN NULL DEFAULT TRUE,
  pro BOOLEAN NULL DEFAULT FALSE,
  PRIMARY KEY (id));

CREATE UNIQUE INDEX email_UNIQUE ON users (email ASC);

CREATE UNIQUE INDEX username_UNIQUE ON users (username ASC);


-- -----------------------------------------------------
-- Table clicks
-- -----------------------------------------------------
DROP TABLE IF EXISTS clicks ;

CREATE TABLE IF NOT EXISTS clicks (
  created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  id SERIAL NOT NULL,
  value INT NULL DEFAULT 0,
  PRIMARY KEY (id));



-- -----------------------------------------------------
-- Table qr_codes
-- -----------------------------------------------------
DROP TABLE IF EXISTS qr_codes ;

CREATE TABLE IF NOT EXISTS qr_codes (
  create_time TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  update_time TIMESTAMP NULL,
  id SERIAL NOT NULL,
  link VARCHAR(255) NOT NULL,
  PRIMARY KEY (id));


-- -----------------------------------------------------
-- Table links
-- -----------------------------------------------------
DROP TABLE IF EXISTS links ;

CREATE TABLE IF NOT EXISTS links (
  created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  id SERIAL NOT NULL,
  id_user BIGINT NOT NULL,
  id_click BIGINT NOT NULL,
  id_qr_code BIGINT NOT NULL,
  original TEXT NOT NULL,
  hash VARCHAR(6) NOT NULL,
  ative BOOLEAN NOT NULL DEFAULT TRUE,
  PRIMARY KEY (id),
  CONSTRAINT fk_links_user
    FOREIGN KEY (id_user)
    REFERENCES users (id)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT fk_links_click
    FOREIGN KEY (id_click)
    REFERENCES clicks (id)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT fk_links_qr_code
    FOREIGN KEY (id_qr_code)
    REFERENCES qr_codes (id)
    ON DELETE CASCADE
    ON UPDATE CASCADE);

CREATE INDEX fk_links_user_idx ON links (id_user ASC);

CREATE INDEX fk_links_click_idx ON links (id_click ASC);

CREATE INDEX fk_links_qr_code_idx ON links (id_qr_code ASC);
-- MySQL Script generated by MySQL Workbench
-- Sun Feb 25 10:46:33 2024
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema mydb
-- -----------------------------------------------------
DROP SCHEMA IF EXISTS `mydb` ;

-- -----------------------------------------------------
-- Schema mydb
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `mydb` DEFAULT CHARACTER SET utf8 ;
USE `mydb` ;

-- -----------------------------------------------------
-- Table `mydb`.`users`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `mydb`.`users` ;

CREATE TABLE IF NOT EXISTS `mydb`.`users` (
  `username` VARCHAR(16) NOT NULL,
  `email` VARCHAR(255) NOT NULL,
  `password` VARCHAR(32) NOT NULL,
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
  `name` VARCHAR(255) NOT NULL,
  `status` TINYINT(1) NULL DEFAULT 1,
  `pro` TINYINT(1) NULL DEFAULT 0,
  PRIMARY KEY (`id`));


-- -----------------------------------------------------
-- Table `mydb`.`users`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `mydb`.`users` ;

CREATE TABLE IF NOT EXISTS `mydb`.`users` (
  `username` VARCHAR(16) NOT NULL,
  `email` VARCHAR(255) NOT NULL,
  `password` VARCHAR(32) NOT NULL,
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
  `name` VARCHAR(255) NOT NULL,
  `status` TINYINT(1) NULL DEFAULT 1,
  `pro` TINYINT(1) NULL DEFAULT 0,
  PRIMARY KEY (`id`));

CREATE UNIQUE INDEX `email_UNIQUE` ON `mydb`.`users` (`email` ASC) VISIBLE;

CREATE UNIQUE INDEX `username_UNIQUE` ON `mydb`.`users` (`username` ASC) VISIBLE;


-- -----------------------------------------------------
-- Table `mydb`.`clicks`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `mydb`.`clicks` ;

CREATE TABLE IF NOT EXISTS `mydb`.`clicks` (
  `created_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `value` INT NULL DEFAULT 0,
  PRIMARY KEY (`id`))
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `mydb`.`qr_codes`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `mydb`.`qr_codes` ;

CREATE TABLE IF NOT EXISTS `mydb`.`qr_codes` (
  `create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` TIMESTAMP NULL,
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `link` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`));


-- -----------------------------------------------------
-- Table `mydb`.`links`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `mydb`.`links` ;

CREATE TABLE IF NOT EXISTS `mydb`.`links` (
  `created_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `id_user` BIGINT NOT NULL,
  `id_click` BIGINT NOT NULL,
  `id_qr_code` BIGINT NOT NULL,
  `original` TEXT NOT NULL,
  `hash` VARCHAR(6) NOT NULL,
  `ative` TINYINT(1) NOT NULL DEFAULT 1,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_links_user`
    FOREIGN KEY (`id_user`)
    REFERENCES `mydb`.`users` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_links_click`
    FOREIGN KEY (`id_click`)
    REFERENCES `mydb`.`clicks` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_links_qr_code`
    FOREIGN KEY (`id_qr_code`)
    REFERENCES `mydb`.`qr_codes` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
DEFAULT CHARACTER SET = utf8;

CREATE INDEX `index1` ON `mydb`.`links` () VISIBLE;

CREATE INDEX `fk_links_user_idx` ON `mydb`.`links` (`id_user` ASC) VISIBLE;

CREATE INDEX `fk_links_click_idx` ON `mydb`.`links` (`id_click` ASC) VISIBLE;

CREATE INDEX `fk_links_qr_code_idx` ON `mydb`.`links` (`id_qr_code` ASC) VISIBLE;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema gradetracker
-- -----------------------------------------------------
DROP SCHEMA IF EXISTS `gradetracker` ;

-- -----------------------------------------------------
-- Schema gradetracker
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `gradetracker` DEFAULT CHARACTER SET utf8 ;
USE `gradetracker` ;

-- -----------------------------------------------------
-- Table `gradetracker`.`tabUser`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `gradetracker`.`tabUser` ;

CREATE TABLE IF NOT EXISTS `gradetracker`.`tabUser` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `firstname` VARCHAR(255) NULL,
  `lastname` VARCHAR(255) NULL,
  `username` VARCHAR(255) NOT NULL,
  `email` VARCHAR(255) NOT NULL,
  `password` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  UNIQUE INDEX `username_UNIQUE` (`username` ASC),
  UNIQUE INDEX `email_UNIQUE` (`email` ASC))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `gradetracker`.`tabSchool`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `gradetracker`.`tabSchool` ;

CREATE TABLE IF NOT EXISTS `gradetracker`.`tabSchool` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `schoolname` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `schoolname_UNIQUE` (`schoolname` ASC),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `gradetracker`.`tabClass`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `gradetracker`.`tabClass` ;

CREATE TABLE IF NOT EXISTS `gradetracker`.`tabClass` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `classname` VARCHAR(255) NOT NULL,
  `school_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_tabClass_tabSchool1_idx` (`school_id` ASC),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  CONSTRAINT `fk_tabClass_tabSchool1`
    FOREIGN KEY (`school_id`)
    REFERENCES `gradetracker`.`tabSchool` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `gradetracker`.`tabSubject`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `gradetracker`.`tabSubject` ;

CREATE TABLE IF NOT EXISTS `gradetracker`.`tabSubject` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `subject` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  UNIQUE INDEX `subject_UNIQUE` (`subject` ASC))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `gradetracker`.`tabGrade`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `gradetracker`.`tabGrade` ;

CREATE TABLE IF NOT EXISTS `gradetracker`.`tabGrade` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `grade` FLOAT NOT NULL,
  `date` DATE NOT NULL,
  `student_id` INT NOT NULL,
  `class_id` INT UNSIGNED NOT NULL,
  `subject_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  INDEX `fk_tabGrade_tabUser_idx` (`student_id` ASC),
  INDEX `fk_tabGrade_tabClass1_idx` (`class_id` ASC),
  INDEX `fk_tabGrade_tabSubject1_idx` (`subject_id` ASC),
  CONSTRAINT `fk_tabGrade_tabUser`
    FOREIGN KEY (`student_id`)
    REFERENCES `gradetracker`.`tabUser` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_tabGrade_tabClass1`
    FOREIGN KEY (`class_id`)
    REFERENCES `gradetracker`.`tabClass` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_tabGrade_tabSubject1`
    FOREIGN KEY (`subject_id`)
    REFERENCES `gradetracker`.`tabSubject` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;

CREATE DATABASE IF NOT EXISTS Proyecto;
USE Proyecto;
CREATE TABLE `Proyecto`.`cpuinfo` (
  `idcpuinfo` INT NOT NULL AUTO_INCREMENT,
  `freecpu` INT NULL,
  `boundcpu` INT NULL,
  `time` DATE NULL,
  PRIMARY KEY (`idcpuinfo`));

CREATE TABLE `Proyecto`.`raminfo` (
  `idraminfo` INT NOT NULL AUTO_INCREMENT,
  `freeram` INT NULL,
  `boundram` INT NULL,
  `time` DATE NULL,
  PRIMARY KEY (`idraminfo`));

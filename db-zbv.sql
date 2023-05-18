
/*
   Use this script only during setup somas to set nul values
   To Login as an admin, use the cli and create one then login via the web interface
   at /login
*/
CREATE DATABASE `zbv`;
USE `zbv`;

CREATE TABLE IF NOT EXISTS  `servers`(
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `serverid` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL
)ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS  `apikey`(
  `serverid` varchar(255) NOT NULL,
  `key` varchar(255) NOT NULL,
  `comment` varchar(255) NOT NULL,
  `active` BOOLEAN NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL
)ENGINE=InnoDB DEFAULT CHARSET=latin1;

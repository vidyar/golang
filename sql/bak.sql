create database if not exists gosense;
use gosense;
CREATE TABLE `top_article` (
  `aid` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8 DEFAULT '',
  `content` longtext CHARACTER SET utf8,
  `publish_time` datetime DEFAULT current_timestamp,
  `publish_status` tinyint(1) DEFAULT 1,
  PRIMARY KEY (`aid`),
  FULLTEXT KEY `content` (`title`,`content`)
) ENGINE=InnoDB AUTO_INCREMENT=3856 DEFAULT CHARSET=utf8mb4

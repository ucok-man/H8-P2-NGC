CREATE TABLE `heroes` (
  `hero_id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `universe` varchar(255) NOT NULL,
  `image_url` varchar(255) NOT NULL
);

CREATE TABLE `villains` (
  `villain_id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `universe` varchar(255) NOT NULL,
  `image_url` varchar(255) NOT NULL
);

CREATE TABLE `crime_cases` (
  `crime_case_id` int PRIMARY KEY AUTO_INCREMENT,
  `hero_id` int NOT NULL COMMENT 'hero who handles',
  `villain_id` int NOT NULL,
  `description` text NOT NULL,
  `incident_date` date NOT NULL
);

ALTER TABLE `crime_cases` ADD FOREIGN KEY (`hero_id`) REFERENCES `heroes` (`hero_id`) ON DELETE CASCADE;

ALTER TABLE `crime_cases` ADD FOREIGN KEY (`villain_id`) REFERENCES `villains` (`villain_id`) ON DELETE CASCADE;

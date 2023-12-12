CREATE TABLE `heroes` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `universe` varchar(255) NOT NULL,
  `skill` varchar(255) NOT NULL,
  `image_url` varchar(255) NOT NULL
);

CREATE TABLE `villains` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `universe` varchar(255) NOT NULL,
  `image_url` varchar(255) NOT NULL
);

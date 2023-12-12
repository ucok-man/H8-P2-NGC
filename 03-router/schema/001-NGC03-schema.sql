CREATE TABLE `inventories` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `item_code` varchar(255) NOT NULL,
  `stock` int NOT NULL,
  `description` varchar(255) NOT NULL,
  `status` enum('active', 'broken') NOT NULL
);
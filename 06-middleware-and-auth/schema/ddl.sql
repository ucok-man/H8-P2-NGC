CREATE TABLE `users` (
  `user_id` int PRIMARY KEY AUTO_INCREMENT,
  `email` varchar(255) NOT NULL UNIQUE,
  `password_hash` blob NOT NULL,
  `name` varchar(15) NOT NULL,
  `age` int NOT NULL,
  `occupation` varchar(255) NOT NULL,
  `role` enum('admin', 'superadmin') NOT NULL
);

CREATE TABLE `recipes` (
  `recipe_id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL UNIQUE,
  `description` text NOT NULL,
  `time_required` int NOT NULL,
  `rating` int NOT NULL
);
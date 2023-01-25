CREATE TABLE `users` (
  `email` varchar(255) UNIQUE PRIMARY KEY,
  `gender` varchar(255) NOT NULL,
  `age` datetime NOT NULL,
  `country` int NOT NULL,
  `diaray_id` integer NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE `diary` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `content` varchar(255),
  `user_email` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT (now())
);

CREATE INDEX `users_index_0` ON `users` (`email`);

CREATE INDEX `diary_index_1` ON `diary` (`user_email`);

ALTER TABLE `diary` ADD FOREIGN KEY (`user_email`) REFERENCES `users` (`email`);
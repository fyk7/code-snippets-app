DROP TABLE IF EXISTS `snippet`;
CREATE TABLE `snippet` (
  `snippet_id` INTEGER NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(100) NOT NULL,
  `description` VARCHAR(1000) NULL,
  `body` VARCHAR(5000) NOT NULL,
  `programing_language` VARCHAR NULL,
  `created_at` DATETIME NOT NULL,
  `created_by` INTEGER NOT NULL,
  `updated_at` DATETIME NOT NULL,
  `updated_by` INTEGER NOT NULL DEFAULT NULL,
  PRIMARY KEY (`snippet_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8_bin COMMENT='Code snippets.';

DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag` (
  `tag_id` INTEGER NOT NULL AUTO_INCREMENT,
  `tag_name` VARCHAR(100) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `created_by` INTEGER NOT NULL,
  `updated_at` DATETIME NOT NULL,
  `updated_by` INTEGER NOT NULL,
  PRIMARY KEY (`tag_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8_bin COMMENT='Tag for code snippets.';

DROP TABLE IF EXISTS `snippet_tag_relation`;
CREATE TABLE `snippet_tag_relation` (
  `snippet_tag_relation_id` INTEGER NOT NULL AUTO_INCREMENT,
  `snippet_id` INTEGER NOT NULL,
  `tag_id` INTEGER NOT NULL,
  `created_at` DATETIME NOT NULL,
  `created_by` INTEGER NOT NULL,
  `updated_at` DATETIME NOT NULL,
  `updated_by` INTEGER NOT NULL,
  PRIMARY KEY (`snippet_tag_relation_id`),
  FOREIGN KEY (snippet_id) REFERENCES `snippet` (`snippet_id`),
  FOREIGN KEY (tag_id) REFERENCES `tag` (`tag_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8_bin;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `user_id` INTEGER NOT NULL AUTO_INCREMENT,
  `user_name` VARCHAR(100) NOT NULL,
  `is_superuser` TINYINT(1) NOT NULL,
  `email` VARCHAR(100),
  `created_at` DATETIME NOT NULL,
  `created_by` INTEGER NOT NULL,
  `updated_at` DATETIME NOT NULL,
  `updated_by` INTEGER NOT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8_bin;
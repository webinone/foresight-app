CREATE TABLE `tb_commentboard` (
	`id` BIGINT(20) NOT NULL AUTO_INCREMENT,
	`title` VARCHAR(255) NULL DEFAULT NULL,
	`content` VARCHAR(255) NULL DEFAULT NULL,
	`user_id` BIGINT(20) NULL DEFAULT NULL,
	`created_at` TIMESTAMP NULL DEFAULT NULL,
	PRIMARY KEY (`id`)
)
COLLATE='utf8_general_ci'
ENGINE=InnoDB;


CREATE TABLE `tb_comment` (
	`id` BIGINT(20) NOT NULL AUTO_INCREMENT,
	`content` VARCHAR(255) NULL DEFAULT NULL,
	`comment_board_id` BIGINT(20) NULL DEFAULT NULL,
	`user_id` BIGINT(20) NULL DEFAULT NULL,
	`created_at` TIMESTAMP NULL DEFAULT NULL,
	PRIMARY KEY (`id`)
)
COLLATE='utf8_general_ci'
ENGINE=InnoDB;

CREATE TABLE `tb_user` (
	`id` BIGINT(20) NOT NULL AUTO_INCREMENT,
	`name` VARCHAR(255) NULL DEFAULT NULL,
	PRIMARY KEY (`id`)
)
COLLATE='utf8_general_ci'
ENGINE=InnoDB;

CREATE TABLE `tb_departure` (
	`id` BIGINT(20) NOT NULL AUTO_INCREMENT,
	`name` VARCHAR(255) NULL DEFAULT NULL,
	PRIMARY KEY (`id`)
)
COLLATE='utf8_general_ci'
ENGINE=InnoDB;


CREATE TABLE `tb_user_departures` (
	`user_id` BIGINT(20) NOT NULL DEFAULT '0',
	`departure_id` BIGINT(20) NOT NULL DEFAULT '0',
	PRIMARY KEY (`user_id`, `departure_id`)
)
COLLATE='utf8_general_ci'
ENGINE=InnoDB;

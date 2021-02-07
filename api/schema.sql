CREATE TABLE `users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `auth_id` varchar(255) NOT NULL,
  `rand_id` varchar(255) NOT NULL,
  `first_name` varchar(255) NOT NULL,
  `last_name` varchar(255) NOT NULL,
  `first_name_kana` varchar(255) NOT NULL,
  `last_name_kana` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `is_beautician` tinyint(1) NOT NULL,
  `phone_number` varchar(11) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  UNIQUE KEY `rand_id` (`rand_id`),
  UNIQUE KEY `auth_id` (`auth_id`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `beauticians` (
  `user_id` bigint NOT NULL,
  `line_id` varchar(255),
  `instagram_id` varchar(255),
  `comment` varchar(255),
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`user_id`),
  CONSTRAINT `user_beautician_fk` FOREIGN KEY (`user_id`) REFERENCES users (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `salons` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `rand_id` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `phone_number` varchar(11) NOT NULL,
  `opening_hours` time NOT NULL,
  `closing_hours` time NOT NULL,
  `postal_code` varchar(8) NOT NULL,
  `prefectures` varchar(255) NOT NULL,
  `city` varchar(255) NOT NULL,
  `town` varchar(255) NOT NULL,
  `address_code` varchar(255) NOT NULL,
  `address_other` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  UNIQUE KEY `rand_id` (`rand_id`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `menus` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `rand_id` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  UNIQUE KEY `rand_id` (`rand_id`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `spaces` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `salon_id` bigint NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `salon_spaces_fk` FOREIGN KEY (`salon_id`) REFERENCES salons (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `reservations` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `rand_id` varchar(255) NOT NULL,
  `date` datetime NOT NULL,
  `holiday` tinyint(1) NOT NULL,
  `space_id` bigint NOT NULL,
  `beautician_id` bigint NOT NULL,
  `user_id` bigint NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `rand_id` (`rand_id`),
  CONSTRAINT `space_reservations_fk` FOREIGN KEY (`space_id`) REFERENCES spaces (`id`),
  CONSTRAINT `beautician_reservations_fk` FOREIGN KEY (`beautician_id`) REFERENCES users (`id`),
  CONSTRAINT `user_reservations_fk` FOREIGN KEY (`user_id`) REFERENCES users (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `beautician_salons` (
  `beautician_id` bigint NOT NULL,
  `salon_id` bigint NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  CONSTRAINT `beautician_beautician_salons_fk` FOREIGN KEY (`beautician_id`) REFERENCES users (`id`),
  CONSTRAINT `salon_beautician_salons_fk` FOREIGN KEY (`salon_id`) REFERENCES salons (`id`),
  PRIMARY KEY (`beautician_id`, `salon_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `beautician_menus` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `rand_id` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `price` bigint NOT NULL,
  `beautician_id` bigint NOT NULL,
  `menu_id` bigint NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `beautician_beautician_menus_fk` FOREIGN KEY (`beautician_id`) REFERENCES users (`id`),
  CONSTRAINT `menu_beautician_menus_fk` FOREIGN KEY (`menu_id`) REFERENCES menus (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `reservation_menus` (
  `reservation_id` bigint NOT NULL,
  `beautician_menu_id` bigint NOT NULL, 
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`reservation_id`, `beautician_menu_id`),
  CONSTRAINT `reservation_reservation_menus_fk` FOREIGN KEY (`reservation_id`) REFERENCES reservations (`id`),
  CONSTRAINT `beautician_menu_reservation_menus_fk` FOREIGN KEY (`beautician_menu_id`) REFERENCES beautician_menus (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `user_salons`  (
  `salon_id` bigint NOT NULL,
  `user_id` bigint NOT NULL,
  `role` enum('admin') NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`salon_id`, `user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
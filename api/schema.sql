CREATE TABLE `beauticians` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `auth_id` varchar(255) NOT NULL,
  `rand_id` varchar(255) NOT NULL,
	`first_name` varchar(255) NOT NULL,
  `last_name` varchar(255) NOT NULL,
  `age` bigint NOT NULL,
  `phone_number` varchar(255) NOT NULL,
  `line_id` varchar(255) NOT NULL,
  `instagram_id` varchar(255) NOT NULL,
  `comment` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `guests` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `auth_id` varchar(255) NOT NULL,
  `rand_id` varchar(255) NOT NULL,
  `first_name` varchar(255) NOT NULL,
  `last_name` varchar(255) NOT NULL,
  `first_name_kana` varchar(255) NOT NULL,
  `last_name_kana` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `menus` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `rand_id` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
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
  `date` datetime NOT NULL,
  `holiday` tinyint(1) NOT NULL,
  `space_id` bigint NOT NULL,
  `beautician_id` bigint NOT NULL,
  `guest_id` bigint NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `space_reservations_fk` FOREIGN KEY (`space_id`) REFERENCES spaces (`id`),
  CONSTRAINT `beautician_reservations_fk` FOREIGN KEY (`beautician_id`) REFERENCES beauticians (`id`),
  CONSTRAINT `guest_reservations_fk` FOREIGN KEY (`guest_id`) REFERENCES guests (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `beautician_salons` (
  `beautician_id` bigint NOT NULL,
  `salon_id` bigint NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`beautician_id`, `salon_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `beautician_menus` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `price` bigint NOT NULL,
  `beautician_id` bigint NOT NULL,
  `menu_id` bigint NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `beautician_beautician_menus_fk` FOREIGN KEY (`beautician_id`) REFERENCES beauticians (`id`),
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
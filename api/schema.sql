CREATE TABLE `beauticians` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `auth_id` varchar(255) NOT NULL,
  `rand_id` varchar(255) NOT NULL,
  `first_name` varchar(255) NOT NULL,
  `last_name` varchar(255) NOT NULL,
  `age` bigint NOT NULL,
  `phone_number` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `salons` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `rand_id` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
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
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `menus` (
  `id` bigint NOT NULL AUTO_INCREMENT,
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

CREATE TABLE `reservation` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `date` date DEFAULT NULL,
  `time` time DEFAULT NULL,
  `space_id` bigint NOT NULL,
  `beautician_id` bigint NOT NULL,
  `guest_id` bigint NOT NULL,
  `menu_id` bigint NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `space_reservations_fk` FOREIGN KEY (`space_id`) REFERENCES spaces (`id`),
  CONSTRAINT `beautician_reservations_fk` FOREIGN KEY (`beautician_id`) REFERENCES beauticians (`id`),
  CONSTRAINT `guest_reservations_fk` FOREIGN KEY (`guest_id`) REFERENCES guests (`id`),
  CONSTRAINT `menu_reservations_fk` FOREIGN KEY (`menu_id`) REFERENCES menus (`id`)
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
  `price` bigint NOT NULL,
  `beautician_id` bigint NOT NULL,
  `menu_id` bigint NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`beautician_id`, `menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
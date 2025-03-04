CREATE TABLE `session` (
  `id` uuid NOT NULL DEFAULT uuid(),
  `summary` varchar(256) NOT NULL,
  `location` varchar(64) NOT NULL,
  `start_time` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

CREATE INDEX `idx_session_start_time` ON session(`start_time`);

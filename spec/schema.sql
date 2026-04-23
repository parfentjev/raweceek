create table `sessions` (
  `id` uuid not null default uuid(),
  `summary` varchar(256) not null,
  `location` varchar(64) not null,
  `start_time` datetime not null,
  primary key (`id`)
);

create index `idx_session_start_time` on session(`start_time`);

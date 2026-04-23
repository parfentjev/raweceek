create table sessions (
  id uuid not null default gen_random_uuid(),
  summary varchar(256) not null,
  location varchar(64) not null,
  start_time timestamp not null,
  primary key (id)
);

create index idx_session_start_time on sessions(start_time);

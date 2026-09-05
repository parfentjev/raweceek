-- name: CountThisWeek :one
select
	count(*)
from
	sessions
where
	start_time >= date_trunc('week', now())
	and start_time < date_trunc('week', now()) + interval '1 week';

-- name: FindNext :one
select
	summary,
	location,
	start_time,
	start_time >= date_trunc('week', now())
	and start_time < date_trunc('week', now()) + interval '1 week' as this_week
from
	sessions
where
	start_time > now()
order by
	start_time
limit 1;

-- name: FindUpcoming :many
with next_week as (
select
    date_trunc('week', start_time at time zone 'UTC') at time zone 'UTC' as week_start
from
    sessions
where
    start_time > now()
order by
    start_time asc
limit 1
)
select
    s.summary,
    s.location,
    s.start_time,
    n.week_start = date_trunc('week', now() at time zone 'UTC') at time zone 'UTC' as this_week
from
    sessions s
cross join next_week n
where
    s.start_time > now()
    and s.start_time >= n.week_start
    and s.start_time < n.week_start + interval '1 week'
order by
    s.start_time asc;
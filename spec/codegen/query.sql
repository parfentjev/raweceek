-- name: GetNextSession :one
select * from sessions
where start_time > now()
order by start_time
limit 1;

-- name: CountSessionsThisWeek :one
select count(*) from sessions
where date_trunc('week', start_time) = date_trunc('week', current_date);

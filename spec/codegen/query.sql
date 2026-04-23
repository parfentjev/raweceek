-- name: GetNextSession :one
select * from session
where start_time > now()
order by start_time
limit 1;

-- name: CountSessionsThisWeek :one
select count(*) from session
where year(start_time) = year(curdate())
and week(start_time, 1) = week(curdate(), 1);

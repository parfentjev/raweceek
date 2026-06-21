use anyhow::Result;
use serde::Serialize;
use sqlx::{PgPool, prelude::FromRow, types::time::OffsetDateTime};
use time::UtcOffset;

use crate::countdown::CountdownDto;

#[derive(FromRow)]
struct Session {
    summary: String,
    location: String,
    start_time: OffsetDateTime,
    this_week: bool,
}

#[derive(Serialize)]
#[serde(rename_all = "camelCase")]
pub struct SessionDto {
    pub summary: String,
    pub location: String,
    #[serde(with = "time::serde::rfc3339")]
    pub start_time: OffsetDateTime,
    pub this_week: bool,
    pub countdowns: Vec<CountdownDto>,
}

impl From<Session> for SessionDto {
    fn from(session: Session) -> Self {
        let current_time = OffsetDateTime::now_utc();
        let session_time = session.start_time.to_offset(UtcOffset::UTC);

        let remaining_time = session_time - current_time;
        let ceeks = CountdownDto::ceeks(&remaining_time);
        let time_until = CountdownDto::time_until(&remaining_time);

        let countdowns = match time_until {
            Ok(time_until) => vec![ceeks, time_until],
            Err(e) => {
                eprintln!("time_until failed: {e}");
                vec![ceeks]
            }
        };

        Self {
            summary: session.summary,
            location: session.location,
            start_time: session_time,
            this_week: session.this_week,
            countdowns,
        }
    }
}

pub async fn count_this_week(db: &PgPool) -> Result<i64> {
    let count = sqlx::query_scalar(
        "select count(*) from sessions where start_time >= date_trunc('week', now()) and start_time < date_trunc('week', now()) + interval '1 week'",
    ).fetch_one(db).await?;

    Ok(count)
}

pub async fn find_next(db: &PgPool) -> Result<SessionDto> {
    let session = sqlx::query_as::<_, Session>(
        "select summary, location, start_time, start_time >= date_trunc('week', now()) and start_time < date_trunc('week', now()) + interval '1 week' as this_week from sessions where start_time > now() order by start_time limit 1"
    )
    .fetch_one(db)
    .await?;

    Ok(session.into())
}

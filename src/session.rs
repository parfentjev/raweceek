use serde::Serialize;
use sqlx::{PgPool, prelude::FromRow, types::time::OffsetDateTime};
use time::UtcOffset;

use crate::countdown::{self, CountdownDto};

#[derive(Debug, thiserror::Error)]
pub enum Error {
    #[error("next session not found")]
    NotFound,
    #[error("database error: {0}")]
    Database(#[source] sqlx::Error),
    #[error("countdown error: {0}")]
    Countdown(#[from] countdown::Error),
}

impl From<sqlx::Error> for Error {
    fn from(value: sqlx::Error) -> Self {
        match value {
            sqlx::Error::RowNotFound => Self::NotFound,
            error => Self::Database(error),
        }
    }
}

#[derive(FromRow)]
struct Session {
    summary: String,
    location: String,
    start_time: OffsetDateTime,
    this_week: bool,
}

#[derive(FromRow)]
struct SessionV2 {
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

impl SessionDto {
    fn from_session(session: Session) -> Result<Self, Error> {
        let current_time = OffsetDateTime::now_utc();
        let session_time = session.start_time.to_offset(UtcOffset::UTC);
        let remaining_time = session_time - current_time;

        Ok(Self {
            summary: session.summary,
            location: session.location,
            start_time: session_time,
            this_week: session.this_week,
            countdowns: vec![
                CountdownDto::ceeks(&remaining_time),
                CountdownDto::time_until(&remaining_time)?,
            ],
        })
    }
}

#[derive(Serialize)]
#[serde(rename_all = "snake_case")]
pub struct SessionDtoV2 {
    pub summary: String,
    pub location: String,
    #[serde(with = "time::serde::rfc3339")]
    pub start_time: OffsetDateTime,
    pub this_week: bool,
    pub countdowns: Vec<CountdownDto>,
}

impl TryFrom<SessionV2> for SessionDtoV2 {
    type Error = Error;

    fn try_from(session: SessionV2) -> Result<Self, Self::Error> {
        let current_time = OffsetDateTime::now_utc();
        let session_time = session.start_time.to_offset(UtcOffset::UTC);
        let remaining_time = session_time - current_time;

        Ok(Self {
            summary: session.summary,
            location: session.location,
            start_time: session_time,
            this_week: session.this_week,
            countdowns: vec![
                CountdownDto::ceeks(&remaining_time),
                CountdownDto::time_until(&remaining_time)?,
            ],
        })
    }
}

pub async fn count_this_week(db: &PgPool) -> Result<i64, Error> {
    let count = sqlx::query_scalar(
        "select count(*) from sessions where start_time >= date_trunc('week', now()) and start_time < date_trunc('week', now()) + interval '1 week'",
    )
    .fetch_one(db)
    .await?;

    Ok(count)
}

pub async fn find_next(db: &PgPool) -> Result<SessionDto, Error> {
    let session = sqlx::query_as::<_, Session>(
        "select summary, location, start_time, start_time >= date_trunc('week', now()) and start_time < date_trunc('week', now()) + interval '1 week' as this_week from sessions where start_time > now() order by start_time limit 1",
    )
    .fetch_one(db)
    .await?;

    SessionDto::from_session(session)
}

/// Finds upcoming sessions in the UTC week that contains the next upcoming session. A successful result always contains at least one session.
///
/// # Errors
///
/// Returns [`Error::NotFound`] if there are no upcoming sessions. It also returns
/// an error if the database query or session conversion fails.
pub async fn find_upcoming(db: &PgPool) -> Result<Vec<SessionDtoV2>, Error> {
    let query = r#"
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
    "#;

    let sessions = sqlx::query_as::<_, SessionV2>(query)
        .fetch_all(db)
        .await?
        .into_iter()
        .map(SessionV2::try_into)
        .collect::<Result<Vec<_>, _>>()?;

    if sessions.is_empty() {
        return Err(Error::NotFound);
    }

    Ok(sessions)
}

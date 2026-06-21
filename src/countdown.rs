use anyhow::Result;
use serde::Serialize;
use std::fmt::Write;
use time::Duration;

const SECONDS_PER_WEEK: f64 = 604_800.0;

#[derive(Serialize)]
#[serde(rename_all = "SCREAMING_SNAKE_CASE")]
pub enum CountdownType {
    Ceeks,
    TimeUntil,
}

#[derive(Serialize)]
pub struct CountdownDto {
    #[serde(rename(serialize = "type"))]
    pub kind: CountdownType,
    pub value: String,
}

impl CountdownDto {
    pub fn ceeks(remaining_time: &Duration) -> Self {
        let ceeks = (remaining_time.as_seconds_f64() / SECONDS_PER_WEEK).max(0.001);

        Self {
            kind: CountdownType::Ceeks,
            value: format!("{ceeks:.3} ceeks"),
        }
    }

    pub fn time_until(remaining_time: &Duration) -> Result<Self> {
        let remaining_time = remaining_time.whole_seconds();

        let seconds = (remaining_time) % 60;
        let minutes = (remaining_time / 60) % 60;
        let hours = (remaining_time / (60 * 60)) % 24;

        // This part isn't very precise because not all months are exactly 30 days long,
        // But this level of precision is acceptable.
        let total_days = remaining_time / (60 * 60 * 24);
        let months = total_days / 30;

        let remaining_days = total_days % 30;
        let weeks = remaining_days / 7;
        let days = remaining_days % 7;

        let mut time_until = String::new();
        let value_unit_map = [
            (months, "month"),
            (weeks, "week"),
            (days, "day"),
            (hours, "hour"),
            (minutes, "minute"),
            (seconds, "second"),
        ];

        for (value, unit) in value_unit_map.iter() {
            // Don't add zero-values like '0 seconds' to the output.
            if *value < 1 {
                continue;
            }

            // If there is already some text in the output, add a separator.
            if !time_until.is_empty() {
                write!(&mut time_until, ", ")?;
            }

            // Finally, append this value-unit pair.
            Self::append_unit(&mut time_until, value, unit)?;
        }

        if time_until.is_empty() {
            write!(&mut time_until, "0s")?;
        }

        Ok(Self {
            kind: CountdownType::TimeUntil,
            value: time_until,
        })
    }

    fn append_unit(out: &mut String, value: &i64, unit: &str) -> Result<()> {
        write!(out, "{value} {unit}")?;

        // Append 's' to values greater than 1, e.g.:
        // - 1 second
        // - 2 seconds
        if *value > 1 {
            write!(out, "s")?;
        }

        Ok(())
    }
}

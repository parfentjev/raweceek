use std::fmt;

use serde::Serialize;
use std::fmt::Write;
use time::Duration;

const SECONDS_PER_WEEK: f64 = 604_800.0;

#[derive(thiserror::Error, Debug)]
pub enum Error {
    #[error("time_until: failed to write to String: {0}")]
    WriteError(#[from] fmt::Error),
}

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

struct TimeUntilUnit {
    name: &'static str,
    value: i64,
}

impl CountdownDto {
    pub fn ceeks(remaining_time: &Duration) -> Self {
        let ceeks = (remaining_time.as_seconds_f64() / SECONDS_PER_WEEK).max(0.001);

        Self {
            kind: CountdownType::Ceeks,
            value: format!("{ceeks:.3} ceeks"),
        }
    }

    pub fn time_until(remaining_time: &Duration) -> Result<Self, Error> {
        let time_units = extract_time_units(remaining_time);
        let mut non_zero_values = time_units.iter().filter(|v| v.value > 0).peekable();

        let mut result = String::new();
        while let Some(time_unit) = non_zero_values.next() {
            let TimeUntilUnit { name, value } = time_unit;

            add_separator(&mut result, non_zero_values.peek().is_none())?;
            append_unit(&mut result, value, name)?;
        }

        if result.is_empty() {
            write!(&mut result, "0s")?;
        }

        Ok(Self {
            kind: CountdownType::TimeUntil,
            value: result,
        })
    }
}

fn extract_time_units(remaining_time: &Duration) -> [TimeUntilUnit; 6] {
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

    [
        TimeUntilUnit {
            name: "month",
            value: months,
        },
        TimeUntilUnit {
            name: "week",
            value: weeks,
        },
        TimeUntilUnit {
            name: "day",
            value: days,
        },
        TimeUntilUnit {
            name: "hour",
            value: hours,
        },
        TimeUntilUnit {
            name: "minute",
            value: minutes,
        },
        TimeUntilUnit {
            name: "second",
            value: seconds,
        },
    ]
}

fn add_separator(result: &mut String, is_last: bool) -> Result<(), Error> {
    if result.is_empty() {
        return Ok(());
    }

    if is_last {
        write!(result, " and ")?;
    } else {
        write!(result, ", ")?;
    }

    Ok(())
}

fn append_unit(out: &mut String, value: &i64, unit: &str) -> Result<(), Error> {
    write!(out, "{value} {unit}")?;

    // Append 's' to values greater than 1, e.g.:
    // - 1 second
    // - 2 seconds
    if *value > 1 {
        write!(out, "s")?;
    }

    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn ceeks_formats_exact_weeks() {
        let countdown = CountdownDto::ceeks(&Duration::days(14));

        assert!(matches!(countdown.kind, CountdownType::Ceeks));
        assert_eq!(countdown.value, "2.000 ceeks");
    }

    #[test]
    fn ceeks_rounds_to_three_decimal_places() {
        let countdown = CountdownDto::ceeks(&Duration::days(10));

        assert_eq!(countdown.value, "1.429 ceeks");
    }

    #[test]
    fn ceeks_enforces_minimum_value() {
        for duration in [Duration::ZERO, Duration::seconds(1), Duration::seconds(-1)] {
            let countdown = CountdownDto::ceeks(&duration);

            assert_eq!(countdown.value, "0.001 ceeks");
        }
    }

    #[test]
    fn time_until_formats_all_units() {
        let duration =
            Duration::days(77) + Duration::hours(4) + Duration::minutes(5) + Duration::seconds(6);

        let countdown = CountdownDto::time_until(&duration).unwrap();

        assert!(matches!(countdown.kind, CountdownType::TimeUntil));
        assert_eq!(
            countdown.value,
            "2 months, 2 weeks, 3 days, 4 hours, 5 minutes and 6 seconds"
        );
    }

    #[test]
    fn time_until_uses_singular_units() {
        let duration =
            Duration::days(38) + Duration::hours(1) + Duration::minutes(1) + Duration::seconds(1);

        let countdown = CountdownDto::time_until(&duration).unwrap();

        assert_eq!(
            countdown.value,
            "1 month, 1 week, 1 day, 1 hour, 1 minute and 1 second"
        );
    }

    #[test]
    fn time_until_omits_zero_units() {
        let duration = Duration::days(30) + Duration::seconds(2);

        let countdown = CountdownDto::time_until(&duration).unwrap();

        assert_eq!(countdown.value, "1 month and 2 seconds");
    }

    #[test]
    fn time_until_uses_and_when_last_nonzero_unit_is_not_seconds() {
        let duration = Duration::hours(1) + Duration::minutes(1);

        let countdown = CountdownDto::time_until(&duration).unwrap();

        assert_eq!(countdown.value, "1 hour and 1 minute");
    }

    #[test]
    fn time_until_formats_non_positive_and_subsecond_durations_as_zero() {
        for duration in [
            Duration::ZERO,
            Duration::milliseconds(999),
            Duration::seconds(-1),
        ] {
            let countdown = CountdownDto::time_until(&duration).unwrap();

            assert_eq!(countdown.value, "0s");
        }
    }
}

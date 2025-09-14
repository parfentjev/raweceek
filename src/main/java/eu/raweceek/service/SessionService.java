package eu.raweceek.service;

import eu.raweceek.codegen.model.CountdownDto;
import eu.raweceek.codegen.model.SessionDto;
import eu.raweceek.repository.SessionRepository;
import org.springframework.data.util.Pair;
import org.springframework.stereotype.Service;

import java.time.Instant;
import java.time.OffsetDateTime;
import java.time.temporal.WeekFields;
import java.util.Set;
import java.util.stream.Stream;

@Service
public class SessionService {
  private static final double SECONDS_IN_WEEK = 604_800;

  private final SessionRepository repository;

  public SessionService(SessionRepository repository) {
    this.repository = repository;
  }

  public eu.raweceek.codegen.model.SessionDto getNextSession() {
    var session = repository.findNext();
    if (session == null) return null;

    var targetWeek = session.getStartTime().get(WeekFields.ISO.weekOfYear());
    var currentWeek = OffsetDateTime.now().get(WeekFields.ISO.weekOfYear());
    var thisWeek = targetWeek == currentWeek;

    var remainingTime = calculateRemainingTime(session.getStartTime());
    var timeUntil = calculateTimeUntil(remainingTime);
    var ceeks = calculateCeeks(remainingTime);

    return new SessionDto()
      .summary(session.getSummary())
      .location(session.getLocation())
      .startTime(session.getStartTime())
      .thisWeek(thisWeek)
      .countdowns(Set.of(timeUntil, ceeks));
  }

  public CountdownDto calculateTimeUntil(long remainingTime) {
    long seconds = (remainingTime) % 60;
    long minutes = (remainingTime / 60) % 60;
    long hours = (remainingTime / (60 * 60)) % 24;
    long days = (remainingTime / (60 * 60 * 24)) % 30;

    long totalDays = remainingTime / (60 * 60 * 24);
    long months = totalDays / 30;
    long weeks = (totalDays % 30) / 7;

    var result = new StringBuilder();
    Stream.of(
        Pair.of(months, "month"),
        Pair.of(weeks, "week"),
        Pair.of(days, "day"),
        Pair.of(hours, "hour"),
        Pair.of(minutes, "minute")
      )
      .filter(entry -> entry.getFirst() > 0)
      .forEach(entry -> appendTimeUnit(result, entry.getFirst(), entry.getSecond(), " "));

    if (!result.isEmpty()) {
      result.append("and ");
    }

    appendTimeUnit(result, seconds, "second", "");

    return new CountdownDto().type(CountdownDto.TypeEnum.TIME_UNTIL).value(result.toString());
  }

  private long calculateRemainingTime(OffsetDateTime startTime) {
    var targetTime = startTime.toEpochSecond();
    var currentTime = Instant.now().getEpochSecond();

    return targetTime - currentTime;
  }

  private void appendTimeUnit(StringBuilder builder, long value, String name, String end) {
    builder.append(value).append(" ").append(name);

    if (value > 1) {
      builder.append("s");
    }

    builder.append(end);
  }

  private CountdownDto calculateCeeks(long remainingTime) {
    var result = String.format("%.2f ceeks", remainingTime / SECONDS_IN_WEEK);

    return new CountdownDto().type(CountdownDto.TypeEnum.CEEKS).value(result);
  }
}

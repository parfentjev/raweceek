package eu.raweceek.service.session.service;

import eu.raweceek.codegen.models.CountdownDto;
import eu.raweceek.codegen.models.SessionDto;
import org.springframework.stereotype.Service;

import java.time.Instant;
import java.time.format.DateTimeFormatter;
import java.util.List;
import java.util.Map;

@Service
public class CountdownService {
  private static final Map<String, Double> UNIT_SECONDS_MAP = Map.of(
    "ceeks", 604_800.0,
    "supermaxes", 225.0,
    "dog years", 4_505_142.0,
    "eye blinks", 0.3
  );

  public List<CountdownDto> calculateRemainingTime(SessionDto sessionDto) {
    var targetTime = DateTimeFormatter.ISO_DATE_TIME.parse(sessionDto.getStartTime(), Instant::from).getEpochSecond();
    var currentTime = Instant.now().getEpochSecond();
    var remainingTime = targetTime - currentTime;

    return UNIT_SECONDS_MAP.entrySet()
      .stream()
      .map(e -> new CountdownDto().value(remainingTime / e.getValue()).unit(e.getKey()))
      .toList();
  }
}

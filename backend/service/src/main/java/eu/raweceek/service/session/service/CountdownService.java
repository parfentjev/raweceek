package eu.raweceek.service.session.service;

import java.time.Instant;
import java.time.OffsetDateTime;
import java.util.List;
import java.util.Map;

import org.springframework.data.util.Pair;
import org.springframework.stereotype.Service;

import eu.raweceek.codegen.models.CountdownDto;

@Service
public class CountdownService {

    private static final Map<CountdownDto.UnitEnum, Double> UNIT_SECONDS_MAP = Map.of(
            CountdownDto.UnitEnum.CEEKS, 604_800.0,
            CountdownDto.UnitEnum.SUPER_MAXES, 225.0,
            CountdownDto.UnitEnum.DOG_YEARS, 4_505_142.0,
            CountdownDto.UnitEnum.EYE_BLINKS, 3.0
    );

    public List<CountdownDto> getCountdowns(OffsetDateTime startTime) {
        var remainingTime = getRemainingTime(startTime);

        return UNIT_SECONDS_MAP.entrySet()
                .stream()
                .map(e -> new CountdownDto().value(remainingTime / e.getValue()).unit(e.getKey()))
                .toList();
    }

    public String getTimeUntil(OffsetDateTime startTime) {
        var remainingTime = getRemainingTime(startTime);

        long seconds = (remainingTime) % 60;
        long minutes = (remainingTime / 60) % 60;
        long hours = (remainingTime / (60 * 60)) % 24;
        long days = (remainingTime / (60 * 60 * 24)) % 30;

        long totalDays = remainingTime / (60 * 60 * 24);
        long months = totalDays / 30;
        long weeks = (totalDays % 30) / 7;

        var result = new StringBuilder();
        List.of(
                Pair.of("month", months),
                Pair.of("week", weeks),
                Pair.of("day", days),
                Pair.of("hour", hours),
                Pair.of("minute", minutes)
        )
                .stream()
                .filter(entry -> entry.getSecond() > 0)
                .forEach(entry -> appendTimeUnit(result, entry.getSecond(), entry.getFirst(), ", "));

        if (result.length() > 0) {
            result.append("and ");
        }

        appendTimeUnit(result, seconds, "second", "");

        return result.toString();
    }

    private long getRemainingTime(OffsetDateTime startTime) {
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
}

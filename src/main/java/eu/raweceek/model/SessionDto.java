package eu.raweceek.model;

import java.time.OffsetDateTime;
import java.util.Set;

public record SessionDto(String summary,
                         String location,
                         OffsetDateTime startTime,
                         boolean thisWeek,
                         Set<CountdownDto> countdowns) {
}

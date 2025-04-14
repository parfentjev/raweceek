package eu.raweceek.service.session.service;

import java.time.Instant;
import java.time.format.DateTimeFormatter;
import java.util.List;
import java.util.Map;

import org.springframework.stereotype.Service;

import eu.raweceek.codegen.models.CountdownDto;
import eu.raweceek.codegen.models.SessionDto;

@Service
public class CountdownService {

    private static final Map<CountdownDto.UnitEnum, Double> UNIT_SECONDS_MAP = Map.of(
            CountdownDto.UnitEnum.CEEKS, 604_800.0,
            CountdownDto.UnitEnum.SUPER_MAXES, 225.0,
            CountdownDto.UnitEnum.DOG_YEARS, 4_505_142.0,
            CountdownDto.UnitEnum.EYE_BLINKS, 3.0
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

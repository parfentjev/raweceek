package eu.raweceek.service.session.contoller;

import static java.lang.String.format;
import java.net.HttpURLConnection;
import static java.time.OffsetDateTime.now;
import static java.time.temporal.ChronoUnit.SECONDS;
import java.util.List;
import java.util.Map;
import java.util.stream.IntStream;

import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.Assertions.within;
import org.junit.jupiter.api.Test;

import eu.raweceek.codegen.models.CountdownDto;
import eu.raweceek.codegen.models.SessionDto;
import eu.raweceek.codegen.models.SessionsCountdownGet200Response;
import eu.raweceek.service.AbstractIntegrationTest;
import eu.raweceek.service.session.model.Session;

public class SessionControllerTest extends AbstractIntegrationTest {

    @Test
    public void getSessions() {
        final var sessionsCount = 5;

        var sessions = IntStream.range(0, sessionsCount)
                .mapToObj(i -> sessionRepository.save(generateSession(i)))
                .toList();

        var response = client.getForEntity("/sessions", SessionDto[].class);
        assertThat(response.getStatusCode().value()).isEqualTo(HttpURLConnection.HTTP_OK);
        assertThat(response.getBody()).isNotNull();
        assertThat(response.getBody().length).isEqualTo(sessions.size());

        sessions.forEach(expected
                -> assertThat(response.getBody()).filteredOnAssertions(actual -> {
                    assertThat(actual.getSummary()).isEqualTo(expected.getSummary());
                    assertThat(actual.getLocation()).isEqualTo(expected.getLocation());
                    assertThat(actual.getStartTime()).isCloseTo(expected.getStartTime(), within(1, SECONDS));
                })
                        .hasSize(1));
    }

    @Test
    public void getSessionsShouldOnlyReturnFutureSessions() {
        var pastSession = generateSession(0);
        pastSession.setStartTime(now());

        var futureSession = generateSession(1);

        sessionRepository.save(pastSession);
        sessionRepository.save(futureSession);

        var response = client.getForEntity("/sessions", SessionDto[].class);
        assertThat(response.getStatusCode().value()).isEqualTo(HttpURLConnection.HTTP_OK);
        assertThat(response.getBody()).isNotNull();
        assertThat(response.getBody().length).isEqualTo(1);
        assertThat(response.getBody()[0].getSummary()).isEqualTo(futureSession.getSummary());
    }

    @Test
    public void getSessionsNextShouldOnlyReturnOneSession() {
        var pastSession = generateSession(0);
        pastSession.setStartTime(now());

        var upcomingSession = generateSession(1);
        upcomingSession.setStartTime(now().plusDays(7));

        var subsequentSession = generateSession(2);
        subsequentSession.setStartTime(now().plusDays(7).plusMinutes(1));

        sessionRepository.saveAll(List.of(pastSession, upcomingSession, subsequentSession));

        var response = client.getForEntity("/sessions/next", SessionDto.class);
        assertThat(response.getStatusCode().value()).isEqualTo(HttpURLConnection.HTTP_OK);
        assertThat(response.getBody()).isNotNull();
        assertThat(response.getBody().getSummary()).isEqualTo(upcomingSession.getSummary());
    }

    @Test
    public void getSessionsCountdownShouldOnlyReturnOneSession() {
        var pastSession = generateSession(0);
        pastSession.setStartTime(now());

        var upcomingSession = generateSession(1);
        upcomingSession.setStartTime(now().plusDays(7));

        var subsequentSession = generateSession(2);
        subsequentSession.setStartTime(now().plusDays(7).plusMinutes(1));

        sessionRepository.saveAll(List.of(pastSession, upcomingSession, subsequentSession));

        var expectedCountdowns = Map.of(
                0.13d, CountdownDto.UnitEnum.DOG_YEARS,
                2016000d, CountdownDto.UnitEnum.EYE_BLINKS,
                1d, CountdownDto.UnitEnum.CEEKS,
                2688d, CountdownDto.UnitEnum.SUPER_MAXES
        );

        var response = client.getForEntity("/sessions/countdown", SessionsCountdownGet200Response.class);
        assertThat(response.getStatusCode().value()).isEqualTo(HttpURLConnection.HTTP_OK);
        assertThat(response.getBody()).isNotNull();
        assertThat(response.getBody().getSession()).isNotNull();
        assertThat(response.getBody().getSession().getSummary()).isEqualTo(upcomingSession.getSummary());
        assertThat(response.getBody().getSession().getLocation()).isEqualTo(upcomingSession.getLocation());
        assertThat(response.getBody().getSession().getStartTime()).isCloseTo(upcomingSession.getStartTime(), within(1, SECONDS));
        assertThat(response.getBody().getCountdowns()).hasSize(expectedCountdowns.size());
        expectedCountdowns.forEach((value, unit)
                -> assertThat(response.getBody().getCountdowns()).filteredOnAssertions(actual -> {
                    assertThat(actual.getValue()).isCloseTo(value, within(0.01));
                    assertThat(actual.getUnit()).isEqualTo(unit);
                })
                        .hasSize(1));
        assertThat(response.getBody().getIsRaceWeek()).isFalse();
    }

    @Test
    public void getSessionsCountdownShouldRaceWeekTrueIfSessionThisWeek() {
        sessionRepository.save(generateSession(0));

        var response = client.getForEntity("/sessions/countdown", SessionsCountdownGet200Response.class);
        assertThat(response.getStatusCode().value()).isEqualTo(HttpURLConnection.HTTP_OK);
        assertThat(response.getBody()).isNotNull();
        assertThat(response.getBody().getIsRaceWeek()).isTrue();
    }

    private Session generateSession(Integer id) {
        return Session.builder()
                .summary(format("THANKS FOR DEPRECATING RANDOM STRING GENERATOR APACHE COMMONS 2025 - Practice %d", id))
                .location("Moon Base City")
                .startTime(now().plusMinutes(1))
                .build();
    }
}

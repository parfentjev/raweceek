package eu.raweceek.service.session.controller;

import static java.time.OffsetDateTime.now;
import java.time.temporal.WeekFields;
import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.RestController;

import eu.raweceek.codegen.api.SessionsApi;
import eu.raweceek.codegen.models.SessionDto;
import eu.raweceek.codegen.models.SessionsCountdownGet200Response;
import eu.raweceek.service.session.model.util.SessionsCountdownGet200ResponseFactory;
import eu.raweceek.service.session.service.CountdownService;
import eu.raweceek.service.session.service.SessionService;

@RestController
public class SessionController implements SessionsApi {

    @Autowired
    private SessionService sessionService;

    @Autowired
    private CountdownService countdownService;

    @Override
    public ResponseEntity<List<SessionDto>> sessionsGet(String series) {
        return ResponseEntity.ok(sessionService.getSessions(series));
    }

    @Override
    public ResponseEntity<SessionDto> sessionsNextGet(String series) {
        return ResponseEntity.ok(sessionService.nextSession(series));
    }

    @Override
    public ResponseEntity<SessionsCountdownGet200Response> sessionsCountdownGet(String series) {
        var sessionDto = sessionService.nextSession(series);
        var countdowns = countdownService.getCountdowns(sessionDto.getStartTime());

        var targetWeek = sessionDto.getStartTime().get(WeekFields.ISO.weekOfYear());
        var currentWeek = now().get(WeekFields.ISO.weekOfYear());

        return ResponseEntity.ok(SessionsCountdownGet200ResponseFactory.withMandatoryData(sessionDto, countdowns, targetWeek == currentWeek));
    }
}

package eu.raweceek.service.session.controller;

import eu.raweceek.codegen.api.SessionsApi;
import eu.raweceek.codegen.models.SessionDto;
import eu.raweceek.codegen.models.SessionsCountdownGet200Response;
import eu.raweceek.service.session.model.util.SessionsCountdownGet200ResponseFactory;
import eu.raweceek.service.session.service.CountdownService;
import eu.raweceek.service.session.service.SessionService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

import static java.time.OffsetDateTime.now;
import static java.time.OffsetDateTime.parse;
import static java.time.temporal.ChronoField.ALIGNED_WEEK_OF_YEAR;

@RestController
public class SessionController implements SessionsApi {
  @Autowired
  private SessionService sessionService;

  @Autowired
  private CountdownService countdownService;

  @Override
  public ResponseEntity<List<SessionDto>> sessionsGet() {
    return ResponseEntity.ok(sessionService.getSessions());
  }

  @Override
  public ResponseEntity<SessionDto> sessionsNextGet() {
    return ResponseEntity.ok(sessionService.nextSession());
  }

  @Override
  public ResponseEntity<SessionsCountdownGet200Response> sessionsCountdownGet() {
    var sessionDto = sessionService.nextSession();
    var countdowns = countdownService.calculateRemainingTime(sessionDto);
    var isRaceWeek = parse(sessionDto.getStartTime()).get(ALIGNED_WEEK_OF_YEAR) == now().get(ALIGNED_WEEK_OF_YEAR);

    return ResponseEntity.ok(SessionsCountdownGet200ResponseFactory.withMandatoryData(sessionDto, countdowns, isRaceWeek));
  }
}

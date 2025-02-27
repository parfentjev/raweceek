package eu.raweceek.service.session.controller;

import eu.raweceek.codegen.api.SessionsApi;
import eu.raweceek.codegen.models.CountdownDto;
import eu.raweceek.codegen.models.SessionDto;
import eu.raweceek.codegen.models.SessionsCountdownGet200Response;
import eu.raweceek.service.session.service.CountdownService;
import eu.raweceek.service.session.service.SessionService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

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

    return ResponseEntity.ok(new SessionsCountdownGet200Response()
      .session(sessionDto)
      .countdowns(countdowns)
      .isRaceWeek(countdowns.stream()
        .filter(c -> c.getUnit() == CountdownDto.UnitEnum.CEEKS)
        .findFirst()
        .orElse(new CountdownDto().value(1.0))
        .getValue() < 1)
    );
  }
}

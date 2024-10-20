package eu.raweceek.service.session.controller;

import eu.raweceek.codegen.api.SessionsApi;
import eu.raweceek.codegen.models.SessionDto;
import eu.raweceek.service.session.service.SessionService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
public class SessionController implements SessionsApi {
  @Autowired
  private SessionService sessionService;

  @Override
  public ResponseEntity<List<SessionDto>> sessionsGet() {
    return ResponseEntity.ok(sessionService.getEvents());
  }

  @Override
  public ResponseEntity<SessionDto> sessionsNextGet() {
    return ResponseEntity.ok(sessionService.nextEvent());
  }
}

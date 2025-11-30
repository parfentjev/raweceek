package eu.raweceek.controller;

import eu.raweceek.codegen.api.SessionApi;
import eu.raweceek.codegen.model.SessionDto;
import eu.raweceek.service.SessionService;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class ApiController implements SessionApi {
  private final SessionService sessionService;

  public ApiController(SessionService sessionService) {
    this.sessionService = sessionService;
  }

  @Override
  public ResponseEntity<SessionDto> getNextSession() {
    var nextSession = sessionService.getNextSession();
    var status = nextSession == null ? 404 : 200;

    return ResponseEntity.status(status).body(nextSession);
  }
}

package eu.raweceek.controller;

import eu.raweceek.model.SessionDto;
import eu.raweceek.service.SessionService;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api")
public class ApiController {
  private final SessionService sessionService;

  public ApiController(SessionService sessionService) {
    this.sessionService = sessionService;
  }

  @GetMapping("/next-session")
  public ResponseEntity<SessionDto> getNextSession() {
    var nextSession = sessionService.getNextSession();
    var status = nextSession == null ? 404 : 200;

    return ResponseEntity.status(status).body(nextSession);
  }
}

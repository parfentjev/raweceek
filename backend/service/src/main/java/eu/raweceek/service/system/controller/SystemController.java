package eu.raweceek.service.system.controller;

import eu.raweceek.codegen.api.HealthApi;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class SystemController implements HealthApi {
  @Override
  public ResponseEntity<Void> healthGet() {
    return ResponseEntity.ok().build();
  }
}

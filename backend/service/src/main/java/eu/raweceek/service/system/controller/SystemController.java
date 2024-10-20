package eu.raweceek.service.system.controller;

import eu.raweceek.codegen.api.SystemApi;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class SystemController implements SystemApi {
  @Override
  public ResponseEntity<Void> systemHealthGet() {
    return ResponseEntity.ok().build();
  }
}

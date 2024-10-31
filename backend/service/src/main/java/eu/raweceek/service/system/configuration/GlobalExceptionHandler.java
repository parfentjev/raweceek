package eu.raweceek.service.system.configuration;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;

import java.util.NoSuchElementException;

@RestControllerAdvice
public class GlobalExceptionHandler {
  @ExceptionHandler
  public ResponseEntity<Void> noSuchElementException(NoSuchElementException e) {
    return ResponseEntity.notFound().build();
  }

  @ExceptionHandler
  public ResponseEntity<Void> exception(Throwable e) {
    return ResponseEntity.internalServerError().build();
  }
}

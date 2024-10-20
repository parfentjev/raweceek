package eu.raweceek.service.event.controller;

import eu.raweceek.codegen.api.EventsApi;
import eu.raweceek.codegen.models.EventDto;
import eu.raweceek.service.event.service.EventService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
public class EventController implements EventsApi {
  @Autowired
  private EventService eventService;

  @Override
  public ResponseEntity<List<EventDto>> eventsGet() {
    return ResponseEntity.ok(eventService.getEvents());
  }

  @Override
  public ResponseEntity<EventDto> eventsNextGet() {
    return ResponseEntity.ok(eventService.nextEvent());
  }
}

package eu.raweceek.service.event.service;

import eu.raweceek.codegen.models.EventDto;
import eu.raweceek.service.event.model.Event;
import eu.raweceek.service.event.model.EventRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.NoSuchElementException;

@Service
public class EventService {
  @Autowired
  private EventRepository eventRepository;

  public List<EventDto> getEvents() {
    return eventRepository.findAll()
      .stream()
      .map(this::mapToDto)
      .toList();
  }

  public EventDto nextEvent() {
    return mapToDto(eventRepository.findNext().orElseThrow(NoSuchElementException::new));
  }

  private EventDto mapToDto(Event event) {
    return new EventDto()
      .summary(event.getSummary())
      .location(event.getLocation())
      .startTime(event.getStartTime().toString());
  }
}

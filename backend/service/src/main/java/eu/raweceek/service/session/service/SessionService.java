package eu.raweceek.service.session.service;

import eu.raweceek.codegen.models.SessionDto;
import eu.raweceek.service.session.model.util.SessionFactory;
import eu.raweceek.service.session.model.SessionRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.NoSuchElementException;

@Service
public class SessionService {
  @Autowired
  private SessionRepository sessionRepository;

  public List<SessionDto> getSessions() {
    return sessionRepository.findUpcoming()
      .stream()
      .map(SessionFactory::mapToDto)
      .toList();
  }

  public SessionDto nextSession() {
    return SessionFactory.mapToDto(sessionRepository.findNext().orElseThrow(NoSuchElementException::new));
  }
}

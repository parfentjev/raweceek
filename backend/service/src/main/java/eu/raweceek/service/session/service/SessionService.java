package eu.raweceek.service.session.service;

import java.util.List;
import java.util.NoSuchElementException;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import eu.raweceek.codegen.models.SessionDto;
import eu.raweceek.service.session.model.Session;
import eu.raweceek.service.session.model.SessionRepository;
import eu.raweceek.service.session.model.util.SessionFactory;

@Service
public class SessionService {

    @Autowired
    private SessionRepository sessionRepository;

    @Autowired
    private CountdownService countdownService;

    public List<SessionDto> getSessions() {
        return sessionRepository.findUpcoming()
                .stream()
                .map(this::mapToDto)
                .toList();
    }

    public SessionDto nextSession() {
        return mapToDto(sessionRepository.findNext().orElseThrow(NoSuchElementException::new));
    }

    private SessionDto mapToDto(Session session) {
        var sessionDto = SessionFactory.mapToDto(session);
        sessionDto.setTimeUntil(countdownService.getTimeUntil(session.getStartTime()));

        return sessionDto;
    }
}

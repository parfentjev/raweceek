package eu.raweceek.service.session.model.util;

import eu.raweceek.codegen.models.SessionDto;
import eu.raweceek.service.session.model.Session;
import lombok.experimental.UtilityClass;

@UtilityClass
public class SessionFactory {

    public SessionDto mapToDto(Session session) {
        return new SessionDto()
                .summary(session.getSummary())
                .location(session.getLocation())
                .startTime(session.getStartTime())
                .series(session.getSeries());
    }
}

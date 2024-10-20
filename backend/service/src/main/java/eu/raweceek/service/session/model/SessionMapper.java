package eu.raweceek.service.session.model;

import eu.raweceek.codegen.models.SessionDto;
import lombok.experimental.UtilityClass;

@UtilityClass
public class SessionMapper {
  public SessionDto mapToDto(Session session) {
    return new SessionDto()
      .summary(session.getSummary())
      .location(session.getLocation())
      .startTime(session.getStartTime().toString());
  }
}

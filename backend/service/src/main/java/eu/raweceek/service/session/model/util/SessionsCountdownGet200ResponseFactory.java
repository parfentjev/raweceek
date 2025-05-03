package eu.raweceek.service.session.model.util;

import java.util.List;

import eu.raweceek.codegen.models.CountdownDto;
import eu.raweceek.codegen.models.SessionDto;
import eu.raweceek.codegen.models.SessionsCountdownGet200Response;
import lombok.experimental.UtilityClass;

@UtilityClass
public class SessionsCountdownGet200ResponseFactory {

    public static SessionsCountdownGet200Response withMandatoryData(SessionDto sessionDto, List<CountdownDto> countdowns, Boolean isRaceWeek) {
        return new SessionsCountdownGet200Response()
                .session(sessionDto)
                .countdowns(countdowns)
                .isRaceWeek(isRaceWeek);
    }
}

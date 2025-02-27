package eu.raweceek.codegen.models;

import java.net.URI;
import java.util.Objects;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonTypeName;
import eu.raweceek.codegen.models.CountdownDto;
import eu.raweceek.codegen.models.SessionDto;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import org.springframework.lang.Nullable;
import java.time.OffsetDateTime;
import jakarta.validation.Valid;
import jakarta.validation.constraints.*;
import io.swagger.v3.oas.annotations.media.Schema;


import java.util.*;
import jakarta.annotation.Generated;

/**
 * SessionsCountdownGet200Response
 */

@JsonTypeName("_sessions_countdown_get_200_response")
@Generated(value = "org.openapitools.codegen.languages.SpringCodegen", comments = "Generator version: 7.12.0-SNAPSHOT")
public class SessionsCountdownGet200Response {

  private @Nullable SessionDto session;

  @Valid
  private List<@Valid CountdownDto> countdowns = new ArrayList<>();

  public SessionsCountdownGet200Response session(SessionDto session) {
    this.session = session;
    return this;
  }

  /**
   * Get session
   * @return session
   */
  @Valid 
  @Schema(name = "session", requiredMode = Schema.RequiredMode.NOT_REQUIRED)
  @JsonProperty("session")
  public SessionDto getSession() {
    return session;
  }

  public void setSession(SessionDto session) {
    this.session = session;
  }

  public SessionsCountdownGet200Response countdowns(List<@Valid CountdownDto> countdowns) {
    this.countdowns = countdowns;
    return this;
  }

  public SessionsCountdownGet200Response addCountdownsItem(CountdownDto countdownsItem) {
    if (this.countdowns == null) {
      this.countdowns = new ArrayList<>();
    }
    this.countdowns.add(countdownsItem);
    return this;
  }

  /**
   * Get countdowns
   * @return countdowns
   */
  @Valid 
  @Schema(name = "countdowns", requiredMode = Schema.RequiredMode.NOT_REQUIRED)
  @JsonProperty("countdowns")
  public List<@Valid CountdownDto> getCountdowns() {
    return countdowns;
  }

  public void setCountdowns(List<@Valid CountdownDto> countdowns) {
    this.countdowns = countdowns;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    SessionsCountdownGet200Response sessionsCountdownGet200Response = (SessionsCountdownGet200Response) o;
    return Objects.equals(this.session, sessionsCountdownGet200Response.session) &&
        Objects.equals(this.countdowns, sessionsCountdownGet200Response.countdowns);
  }

  @Override
  public int hashCode() {
    return Objects.hash(session, countdowns);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SessionsCountdownGet200Response {\n");
    sb.append("    session: ").append(toIndentedString(session)).append("\n");
    sb.append("    countdowns: ").append(toIndentedString(countdowns)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }
}


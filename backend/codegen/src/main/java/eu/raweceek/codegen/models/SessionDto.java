package eu.raweceek.codegen.models;

import java.net.URI;
import java.util.Objects;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonCreator;
import java.time.OffsetDateTime;
import org.springframework.format.annotation.DateTimeFormat;
import org.springframework.lang.Nullable;
import java.time.OffsetDateTime;
import jakarta.validation.Valid;
import jakarta.validation.constraints.*;
import io.swagger.v3.oas.annotations.media.Schema;


import java.util.*;
import jakarta.annotation.Generated;

/**
 * SessionDto
 */

@Generated(value = "org.openapitools.codegen.languages.SpringCodegen", comments = "Generator version: 7.14.0-SNAPSHOT")
public class SessionDto {

  private String summary;

  private String location;

  @DateTimeFormat(iso = DateTimeFormat.ISO.DATE_TIME)
  private OffsetDateTime startTime;

  private @Nullable String timeUntil;

  public SessionDto() {
    super();
  }

  /**
   * Constructor with only required parameters
   */
  public SessionDto(String summary, String location, OffsetDateTime startTime) {
    this.summary = summary;
    this.location = location;
    this.startTime = startTime;
  }

  public SessionDto summary(String summary) {
    this.summary = summary;
    return this;
  }

  /**
   * Get summary
   * @return summary
   */
  @NotNull 
  @Schema(name = "summary", example = "🏁 GRAND PRIX NAME 2024 - Race", requiredMode = Schema.RequiredMode.REQUIRED)
  @JsonProperty("summary")
  public String getSummary() {
    return summary;
  }

  public void setSummary(String summary) {
    this.summary = summary;
  }

  public SessionDto location(String location) {
    this.location = location;
    return this;
  }

  /**
   * Get location
   * @return location
   */
  @NotNull 
  @Schema(name = "location", example = "United Arab Emirates", requiredMode = Schema.RequiredMode.REQUIRED)
  @JsonProperty("location")
  public String getLocation() {
    return location;
  }

  public void setLocation(String location) {
    this.location = location;
  }

  public SessionDto startTime(OffsetDateTime startTime) {
    this.startTime = startTime;
    return this;
  }

  /**
   * Get startTime
   * @return startTime
   */
  @NotNull @Valid 
  @Schema(name = "startTime", example = "2024-12-08T13:00Z", requiredMode = Schema.RequiredMode.REQUIRED)
  @JsonProperty("startTime")
  public OffsetDateTime getStartTime() {
    return startTime;
  }

  public void setStartTime(OffsetDateTime startTime) {
    this.startTime = startTime;
  }

  public SessionDto timeUntil(String timeUntil) {
    this.timeUntil = timeUntil;
    return this;
  }

  /**
   * Get timeUntil
   * @return timeUntil
   */
  
  @Schema(name = "timeUntil", example = "8 hour(s) 48 minute(s) 56 second(s)", requiredMode = Schema.RequiredMode.NOT_REQUIRED)
  @JsonProperty("timeUntil")
  public String getTimeUntil() {
    return timeUntil;
  }

  public void setTimeUntil(String timeUntil) {
    this.timeUntil = timeUntil;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    SessionDto sessionDto = (SessionDto) o;
    return Objects.equals(this.summary, sessionDto.summary) &&
        Objects.equals(this.location, sessionDto.location) &&
        Objects.equals(this.startTime, sessionDto.startTime) &&
        Objects.equals(this.timeUntil, sessionDto.timeUntil);
  }

  @Override
  public int hashCode() {
    return Objects.hash(summary, location, startTime, timeUntil);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SessionDto {\n");
    sb.append("    summary: ").append(toIndentedString(summary)).append("\n");
    sb.append("    location: ").append(toIndentedString(location)).append("\n");
    sb.append("    startTime: ").append(toIndentedString(startTime)).append("\n");
    sb.append("    timeUntil: ").append(toIndentedString(timeUntil)).append("\n");
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


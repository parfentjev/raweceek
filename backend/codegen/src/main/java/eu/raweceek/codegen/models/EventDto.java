package eu.raweceek.codegen.models;

import java.net.URI;
import java.util.Objects;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonCreator;
import java.time.OffsetDateTime;
import jakarta.validation.Valid;
import jakarta.validation.constraints.*;
import io.swagger.v3.oas.annotations.media.Schema;


import java.util.*;
import jakarta.annotation.Generated;

/**
 * EventDto
 */

@Generated(value = "org.openapitools.codegen.languages.SpringCodegen", comments = "Generator version: 7.10.0-SNAPSHOT")
public class EventDto {

  private String summary;

  private String location;

  private String startTime;

  public EventDto summary(String summary) {
    this.summary = summary;
    return this;
  }

  /**
   * Get summary
   * @return summary
   */
  
  @Schema(name = "summary", example = "🏁 FORMULA 1 ETIHAD AIRWAYS ABU DHABI GRAND PRIX 2024 - Race", requiredMode = Schema.RequiredMode.NOT_REQUIRED)
  @JsonProperty("summary")
  public String getSummary() {
    return summary;
  }

  public void setSummary(String summary) {
    this.summary = summary;
  }

  public EventDto location(String location) {
    this.location = location;
    return this;
  }

  /**
   * Get location
   * @return location
   */
  
  @Schema(name = "location", example = "United Arab Emirates", requiredMode = Schema.RequiredMode.NOT_REQUIRED)
  @JsonProperty("location")
  public String getLocation() {
    return location;
  }

  public void setLocation(String location) {
    this.location = location;
  }

  public EventDto startTime(String startTime) {
    this.startTime = startTime;
    return this;
  }

  /**
   * Get startTime
   * @return startTime
   */
  
  @Schema(name = "startTime", example = "2024-12-08T13:00Z", requiredMode = Schema.RequiredMode.NOT_REQUIRED)
  @JsonProperty("startTime")
  public String getStartTime() {
    return startTime;
  }

  public void setStartTime(String startTime) {
    this.startTime = startTime;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    EventDto eventDto = (EventDto) o;
    return Objects.equals(this.summary, eventDto.summary) &&
        Objects.equals(this.location, eventDto.location) &&
        Objects.equals(this.startTime, eventDto.startTime);
  }

  @Override
  public int hashCode() {
    return Objects.hash(summary, location, startTime);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class EventDto {\n");
    sb.append("    summary: ").append(toIndentedString(summary)).append("\n");
    sb.append("    location: ").append(toIndentedString(location)).append("\n");
    sb.append("    startTime: ").append(toIndentedString(startTime)).append("\n");
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


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
 * UserTokenDto
 */

@Generated(value = "org.openapitools.codegen.languages.SpringCodegen", date = "2024-10-19T12:41:34.257186005Z[Etc/UTC]", comments = "Generator version: 7.10.0-SNAPSHOT")
public class UserTokenDto {

  private String token;

  private Integer expires;

  public UserTokenDto() {
    super();
  }

  /**
   * Constructor with only required parameters
   */
  public UserTokenDto(String token, Integer expires) {
    this.token = token;
    this.expires = expires;
  }

  public UserTokenDto token(String token) {
    this.token = token;
    return this;
  }

  /**
   * Get token
   * @return token
   */
  @NotNull 
  @Schema(name = "token", requiredMode = Schema.RequiredMode.REQUIRED)
  @JsonProperty("token")
  public String getToken() {
    return token;
  }

  public void setToken(String token) {
    this.token = token;
  }

  public UserTokenDto expires(Integer expires) {
    this.expires = expires;
    return this;
  }

  /**
   * Get expires
   * @return expires
   */
  @NotNull 
  @Schema(name = "expires", requiredMode = Schema.RequiredMode.REQUIRED)
  @JsonProperty("expires")
  public Integer getExpires() {
    return expires;
  }

  public void setExpires(Integer expires) {
    this.expires = expires;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    UserTokenDto userTokenDto = (UserTokenDto) o;
    return Objects.equals(this.token, userTokenDto.token) &&
        Objects.equals(this.expires, userTokenDto.expires);
  }

  @Override
  public int hashCode() {
    return Objects.hash(token, expires);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class UserTokenDto {\n");
    sb.append("    token: ").append(toIndentedString(token)).append("\n");
    sb.append("    expires: ").append(toIndentedString(expires)).append("\n");
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


package eu.raweceek.controller;

import com.fasterxml.jackson.databind.json.JsonMapper;
import com.fasterxml.jackson.databind.util.StdDateFormat;
import com.fasterxml.jackson.datatype.jsr310.JavaTimeModule;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class Utils {
  @Bean
  public JsonMapper jsonMapper() {
    return JsonMapper.builder()
        .addModule(new JavaTimeModule())
        .defaultDateFormat(new StdDateFormat().withColonInTimeZone(true))
        .build();
  }
}

package eu.raweceek.controller;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.json.JsonMapper;
import eu.raweceek.service.SessionService;
import java.util.Map;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;

@Controller
public class WebController {
  private static final Logger log = LoggerFactory.getLogger(WebController.class);

  private final SessionService sessionService;
  private final JsonMapper jsonMapper;

  public WebController(SessionService sessionService, JsonMapper jsonMapper) {
    this.sessionService = sessionService;
    this.jsonMapper = jsonMapper;
  }

  @GetMapping("/")
  public String index(Model model) {
    var nextSession = sessionService.getNextSession();
    var isRaceWeek = sessionService.isRaceWeek();

    String nextSessionJson;
    try {
      nextSessionJson = jsonMapper.writeValueAsString(nextSession);
    } catch (JsonProcessingException e) {
      log.error("failed to convert nextSession to json", e);
      nextSessionJson = "";
    }

    model.addAllAttributes(
        Map.of("nextSession", nextSession, "isRaceWeek", isRaceWeek, "json", nextSessionJson));

    return "index";
  }
}

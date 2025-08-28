package eu.raweceek.controller;

import eu.raweceek.service.SessionService;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;

@Controller
public class WebController {
  private final SessionService sessionService;

  public WebController(SessionService sessionService) {
    this.sessionService = sessionService;
  }

  @GetMapping("/")
  public String index(Model model) {
    var nextSession = sessionService.getNextSession();
    model.addAttribute("nextSession", nextSession);

    return "index";
  }
}

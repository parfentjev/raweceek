package eu.raweceek.service;

import eu.raweceek.service.session.model.SessionRepository;
import org.junit.jupiter.api.AfterEach;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.web.client.TestRestTemplate;
import org.springframework.test.context.TestPropertySource;

@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
@TestPropertySource(locations = "classpath:integration.properties")
public abstract class AbstractIntegrationTest {
  @Autowired
  protected TestRestTemplate client;

  @Autowired
  protected SessionRepository sessionRepository;

  @AfterEach
  public final void tearDown() {
    sessionRepository.deleteAll();
  }
}

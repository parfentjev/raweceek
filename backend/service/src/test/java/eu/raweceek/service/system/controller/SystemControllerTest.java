package eu.raweceek.service.system.controller;

import eu.raweceek.service.AbstractIntegrationTest;
import org.junit.jupiter.api.Test;

import java.net.HttpURLConnection;

import static org.junit.jupiter.api.Assertions.assertEquals;

public class SystemControllerTest extends AbstractIntegrationTest {
  @Test
  public void healthCheck() {
    var res = client.getForEntity("/system/health", Void.class);
    assertEquals(HttpURLConnection.HTTP_OK, res.getStatusCode().value());
  }
}

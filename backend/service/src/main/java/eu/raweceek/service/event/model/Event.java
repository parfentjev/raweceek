package eu.raweceek.service.event.model;

import jakarta.persistence.*;
import lombok.AccessLevel;
import lombok.Data;
import lombok.experimental.FieldDefaults;

import java.time.OffsetDateTime;

@Entity
@Data
@FieldDefaults(level = AccessLevel.PRIVATE)
public class Event {
  @Id
  @GeneratedValue(strategy = GenerationType.UUID)
  @Column(unique = true)
  String summary;

  @Column(nullable = false)
  String location;

  @Column(nullable = false)
  OffsetDateTime startTime;
}

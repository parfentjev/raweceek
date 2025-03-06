package eu.raweceek.service.session.model;

import jakarta.persistence.*;
import lombok.*;
import lombok.experimental.FieldDefaults;

import java.time.OffsetDateTime;

@Entity
@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
@FieldDefaults(level = AccessLevel.PRIVATE)
public class Session {
  @Id
  @GeneratedValue(strategy = GenerationType.UUID)
  String id;

  @Column(unique = true)
  String summary;

  @Column(nullable = false)
  String location;

  @Column(nullable = false)
  OffsetDateTime startTime;
}

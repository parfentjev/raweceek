package eu.raweceek.service.event.model;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

import java.util.Optional;

@Repository
public interface EventRepository extends JpaRepository<Event, String> {
  @Query(value = "select * from event where start_time > now() order by start_time limit 1", nativeQuery = true)
  Optional<Event> findNext();
}

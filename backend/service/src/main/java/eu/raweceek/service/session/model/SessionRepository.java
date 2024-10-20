package eu.raweceek.service.session.model;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;

@Repository
public interface SessionRepository extends JpaRepository<Session, String> {
  @Query(value = "select * from session where start_time > now() order by start_time", nativeQuery = true)
  List<Session> findUpcoming();

  @Query(value = "select * from session where start_time > now() order by start_time limit 1", nativeQuery = true)
  Optional<Session> findNext();
}

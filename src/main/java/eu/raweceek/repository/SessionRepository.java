package eu.raweceek.repository;

import eu.raweceek.model.Session;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface SessionRepository extends CrudRepository<Session, String> {
  @Query(value = """
    select * from session
    where start_time > now()
    order by start_time
    limit 1
    """, nativeQuery = true)
  Session findNext();
}

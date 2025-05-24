package eu.raweceek.service.session.model;

import java.util.List;
import java.util.Optional;

import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface SessionRepository extends CrudRepository<Session, String> {

    @Query(value = """
    select * from session
    where start_time > now()
    and series = :series
    order by start_time
    """, nativeQuery = true)
    List<Session> findUpcoming(String series);

    @Query(value = """
    select * from session
    where start_time > now()
    and series = :series
    order by start_time
    limit 1
    """, nativeQuery = true)
    Optional<Session> findNext(String series);
}

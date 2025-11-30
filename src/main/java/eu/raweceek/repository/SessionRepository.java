package eu.raweceek.repository;

import eu.raweceek.model.Session;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface SessionRepository extends CrudRepository<Session, String> {
  @Query(
      value =
          """
    select * from session
    where start_time > now()
    order by start_time
    limit 1
    """,
      nativeQuery = true)
  Session findNext();

  @Query(
      value =
          """
        select count(*) from session
        where year(start_time) = year(curdate())
        and week(start_time, 1) = week(curdate(), 1)
        """,
      nativeQuery = true)
  Integer countThisWeek();
}

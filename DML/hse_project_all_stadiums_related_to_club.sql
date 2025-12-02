SELECT s.stadium_id,
       s.stadium_location,
       s.capacity,
       c.club_name AS club_name
FROM stadium s
JOIN stadium_club sc ON s.stadium_id = sc.stadium_id
JOIN club c ON sc.club_id = c.club_id
WHERE c.club_id = 1;

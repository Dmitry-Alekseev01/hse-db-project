SELECT p.player_id,
       p.player_name,
       p.player_surname,
       p.player_number,
       ps.status_type
FROM player p
LEFT JOIN playerstatus ps ON p.status_id = ps.status_id
WHERE p.team_id = 1
ORDER BY p.player_number;

SELECT t.team_id,
       t.team_name,
       COALESCE(SUM(p.salary), 0) AS total_players_salary,
       t.budget
FROM team t
LEFT JOIN player p ON t.team_id = p.team_id
GROUP BY t.team_id, t.team_name, t.budget;

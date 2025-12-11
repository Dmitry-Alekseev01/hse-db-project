DO $$
DECLARE
  rec RECORD;
  deficit NUMERIC;
BEGIN
  FOR rec IN
    SELECT t.team_id, t.team_name, t.budget, COALESCE(SUM(p.salary),0) AS total_salaries
    FROM team t
    LEFT JOIN player p ON p.team_id = t.team_id
    GROUP BY t.team_id, t.team_name, t.budget
    HAVING COALESCE(SUM(p.salary),0) > t.budget
  LOOP
    deficit := rec.total_salaries - rec.budget;
    RAISE NOTICE 'Команда % (id=%) имеет дефицит % (зарплаты=%, бюджет=%).', rec.team_name, rec.team_id, deficit, rec.total_salaries, rec.budget;

    UPDATE team SET budget = budget + deficit WHERE team_id = rec.team_id;
  END LOOP;

  RAISE NOTICE 'Проверка бюджета завершена.';
END $$ LANGUAGE plpgsql;

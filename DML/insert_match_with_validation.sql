BEGIN;

DO $$
DECLARE
  p_stadium_id INT := 1;
  p_home_team INT := 1;
  p_away_team INT := 2;
  p_match_date DATE := '2026-03-15';
  p_expected_attendance INT := 50000;
  v_capacity INT;
  v_conflict_count INT;
BEGIN
  IF p_home_team = p_away_team THEN
    RAISE EXCEPTION 'Домашняя и гостевая команды совпадают (id=%).', p_home_team;
  END IF;

  SELECT capacity
  INTO v_capacity
  FROM stadium
  WHERE stadium_id = p_stadium_id
  FOR UPDATE;

  IF NOT FOUND THEN
    RAISE EXCEPTION 'Стадион с id=% не найден.', p_stadium_id;
  END IF;

  IF v_capacity < p_expected_attendance THEN
    RAISE EXCEPTION 'Ожидаемая посещаемость (%) больше вместимости стадиона (%).', p_expected_attendance, v_capacity;
  END IF;

  SELECT COUNT(*) INTO v_conflict_count
  FROM game
  WHERE stadium_id = p_stadium_id
    AND match_date = p_match_date;

  IF v_conflict_count > 0 THEN
    RAISE EXCEPTION 'На стадионе % уже запланирован матч на дату %.', p_stadium_id, p_match_date;
  END IF;

  INSERT INTO game (stadium_id, team_1_id, team_2_id, match_date)
  VALUES (p_stadium_id, p_home_team, p_away_team, p_match_date);

  RAISE NOTICE 'Матч успешно запланирован: stadium=% date=% home=% away=%',
               p_stadium_id, p_match_date, p_home_team, p_away_team;
END $$ LANGUAGE plpgsql;

COMMIT;

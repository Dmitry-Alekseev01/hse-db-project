DO $$
DECLARE
  p_player_id INT := 2; -- player_id
  p_loan_team_id INT := 2; -- команда, принимающая в аренду
  p_start_date DATE := '2025-12-08';
  p_end_date DATE := '2026-06-30';
  v_old_team INT;
  v_status_id INT;
BEGIN
  SELECT team_id INTO v_old_team FROM player WHERE player_id = p_player_id;
  IF NOT FOUND THEN RAISE EXCEPTION 'Игрок % не найден', p_player_id; END IF;

  IF v_old_team = p_loan_team_id THEN
    RAISE EXCEPTION 'Игрок уже в этой команде';
  END IF;

  -- статус "в аренде"
  SELECT status_id INTO v_status_id FROM playerstatus WHERE status_type = 'в аренде' LIMIT 1;
  IF NOT FOUND THEN
    INSERT INTO playerstatus (status_type, start_date, expiration_date)
      VALUES ('в аренде', p_start_date, p_end_date)
    RETURNING status_id INTO v_status_id;
  ELSE
    INSERT INTO playerstatus (status_type, start_date, expiration_date)
      VALUES ('в аренде', p_start_date, p_end_date)
    RETURNING status_id INTO v_status_id;
  END IF;

  -- Перевод игрока в команду арендатора и новый статус
  UPDATE player SET team_id = p_loan_team_id, status_id = v_status_id WHERE player_id = p_player_id;

  RAISE NOTICE 'Игрок % отдан в аренду команде % с % по %', p_player_id, p_loan_team_id, p_start_date, p_end_date;
END $$ LANGUAGE plpgsql;

DO $$
DECLARE
  p_player_id INT := 1; -- player_id
  p_buyer_team_id INT := 2;  -- id покупающей команды
  p_transfer_fee NUMERIC := 100000.00; -- сумма трансфера
  seller_team_id INT;
  buyer_budget NUMERIC;
  seller_budget NUMERIC;
  v_status_id INT;
BEGIN
  -- текущая команда игрока
  SELECT team_id INTO seller_team_id FROM player WHERE player_id = p_player_id;
  IF seller_team_id IS NULL THEN
    RAISE EXCEPTION 'Игрок с id=% не привязан к команде', p_player_id;
  END IF;

  IF seller_team_id = p_buyer_team_id THEN
    RAISE EXCEPTION 'Покупающая команда = текущая команда';
  END IF;

  SELECT budget INTO buyer_budget FROM team WHERE team_id = p_buyer_team_id FOR UPDATE;
  IF NOT FOUND THEN RAISE EXCEPTION 'Покупающая команда % не найдена', p_buyer_team_id; END IF;

  SELECT budget INTO seller_budget FROM team WHERE team_id = seller_team_id FOR UPDATE;
  IF NOT FOUND THEN RAISE EXCEPTION 'Продающая команда % не найдена', seller_team_id; END IF;

  IF buyer_budget < p_transfer_fee THEN
    RAISE EXCEPTION 'Недостаточно бюджета у покупающей команды: % < %', buyer_budget, p_transfer_fee;
  END IF;

  -- Уменьшается бюджет покупающей
  UPDATE team SET budget = budget - p_transfer_fee WHERE team_id = p_buyer_team_id;

  -- увеличивается бюджет продающей
  UPDATE team SET budget = budget + p_transfer_fee WHERE team_id = seller_team_id;

  -- проверка, что есть статус "на контракте"
  SELECT status_id INTO v_status_id FROM playerstatus WHERE status_type = 'на контракте' LIMIT 1;
  IF NOT FOUND THEN
    INSERT INTO playerstatus (status_type, start_date, expiration_date)
    VALUES ('на контракте', current_date, current_date + INTERVAL '3 years')
    RETURNING status_id INTO v_status_id;
  END IF;

  -- статус "на контракте"
  UPDATE player SET team_id = p_buyer_team_id, status_id = v_status_id WHERE player_id = p_player_id;

  RAISE NOTICE 'Трансфер игрока % завершён: продавец team=% , покупатель team=%, сумма=%',
               p_player_id, seller_team_id, p_buyer_team_id, p_transfer_fee;
END $$ LANGUAGE plpgsql;

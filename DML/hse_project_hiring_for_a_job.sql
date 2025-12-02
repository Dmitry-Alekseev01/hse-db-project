DO $$
DECLARE
  p_club_id INT := 1;
  p_name TEXT := 'Ольга';
  p_surname TEXT := 'Иванова';
  p_salary NUMERIC := 35000.00;
  p_spec_type TEXT := 'врач';
  v_spec_id INT;
BEGIN
  SELECT specification_id INTO v_spec_id FROM staffspecification WHERE spec_type = p_spec_type LIMIT 1;
  IF NOT FOUND THEN
    INSERT INTO staffspecification (spec_type) VALUES (p_spec_type) RETURNING specification_id INTO v_spec_id;
  END IF;

  INSERT INTO staff (staff_name, staff_surname, salary, specification_id, club_id)
  VALUES (p_name, p_surname, p_salary, v_spec_id, p_club_id);

  RAISE NOTICE 'Сотрудник % % принят в клуб % как % (spec_id=%)', p_name, p_surname, p_club_id, p_spec_type, v_spec_id;
END $$ LANGUAGE plpgsql;

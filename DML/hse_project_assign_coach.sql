DO $$
DECLARE
  p_team_id INT := 1;
  p_coach_name TEXT := 'Юрген';
  p_coach_surname TEXT := 'Клопп';
  p_salary NUMERIC := 500000.00;
  p_phone TEXT := '+3757583847692';
  p_specs TEXT[] := ARRAY['главный'];
  v_coach_id INT;
  v_spec_id INT;
  v_spec TEXT;
BEGIN
  INSERT INTO coach (coach_name, coach_surname, salary, phone, team_id)
  VALUES (p_coach_name, p_coach_surname, p_salary, p_phone, p_team_id)
  RETURNING coach_id INTO v_coach_id;

  FOREACH v_spec IN ARRAY p_specs
  LOOP
    SELECT specification_id INTO v_spec_id
    FROM coachspecification
    WHERE specification_name = v_spec
    LIMIT 1;

    IF NOT FOUND THEN
      RAISE EXCEPTION 'Спецификация "%" не найдена в coachspecification', v_spec;
    END IF;

    INSERT INTO coach_specification (specification_id, coach_id)
      VALUES (v_spec_id, v_coach_id);
  END LOOP;

  RAISE NOTICE 'Добавлен тренер % % с id=% и специализациями %',
               p_coach_name, p_coach_surname, v_coach_id, array_to_string(p_specs, ', ');
END $$ LANGUAGE plpgsql;

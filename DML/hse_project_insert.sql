INSERT INTO club (club_name, creation_date, website) VALUES 
('Реал Мадрид', '1902-03-06', 'https://www.realmadrid.com/en-US'),
('Манчестер Сити', '1880-11-23', 'https://www.mancity.com/');

INSERT INTO stadium (capacity, stadium_location, build_date) VALUES 
(81044, 'Мадрид, Испания', '1947-12-14'),
(55097, 'Манчестер, Англия', '2003-08-10');

INSERT INTO stadium_club (club_id, stadium_id) VALUES (1,1), (2,2);

INSERT INTO team (team_name, budget, club_id) VALUES 
('Реал Кастилья', 2000000.00, 1),
('Манчестер Молодежь', 1500000.00, 2);

INSERT INTO playerstatus (status_type, start_date, expiration_date) VALUES 
('на контракте', '2023-01-01', '2026-12-31'),
('в аренде', '2024-07-01', '2025-06-30');

INSERT INTO player (player_name, player_surname, player_number, salary, phone, birth_date, team_id, status_id) VALUES 
('Иван', 'Компани', 9, 50000.00, '+37124949691', '1995-02-10', 1, 1),
('Петр', 'Стерлинг', 10, 60000.00, '+37128684642', '1993-06-20', 1, 1),
('Кевин', 'ДеБургер', 7, 45000.00, '+37123853783', '1998-11-05', 2, 1);

INSERT INTO coachspecification (specification_name) VALUES 
('главный'), ('тренер вратарей'), ('тренер по игре в обороне'), 
('тренер по игре в атаке'), ('тренер стандартных положений'), 
('тренер по физ. подготовке'), ('ассистент');

INSERT INTO coach (coach_name, coach_surname, salary, phone, team_id) VALUES
('Хосеп', 'Гвардиола', 70000.00, '+37130000001', 1),
('Энцо', 'Мареска', 65000.00, '+37130000002', 2);

INSERT INTO coach_specification (specification_id, coach_id) VALUES
(1,1), (7,1), (3,2);

INSERT INTO staffspecification (spec_type) VALUES 
('уборщик'), ('охранник'), ('газонокосильщик'), 
('врач'), ('директор'), ('президент');

INSERT INTO staff (staff_name, staff_surname, salary, specification_id, club_id) VALUES 
('Анна', 'Смит', 31000.00, 4, 1),
('Боб', 'Браун', 40000.00, 6, 2);

INSERT INTO game (stadium_id, team_1_id, team_2_id, match_date) VALUES 
(1, 1, 2, '2025-12-20'),
(2, 2, 1, '2026-01-10');

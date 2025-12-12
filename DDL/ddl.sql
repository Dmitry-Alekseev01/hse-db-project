create table CoachSpecification(
  specification_id serial primary key,
  specification_name varchar(50) not null check(specification_name in ('главный', 'тренер вратарей', 'тренер по игре в обороне', 
  'тренер по игре в атаке', 'тренер стандартных положений', 'тренер по физ. подготовке', 'ассистент')) 
);

create table StaffSpecification(
  specification_id serial primary key,
  spec_type varchar(40) not null check(spec_type in ('уборщик', 'охранник', 'газонокосильщик', 
  'врач', 'директор', 'президент')) 
);

create table Club(
  club_id serial primary key,
  club_name varchar(50) not null,
  creation_date date not null default current_date,
  website varchar(100)
);

create table Team(
  team_id serial primary key,
  team_name varchar(50) not null,
  budget decimal not null check(budget > 100000),
  club_id int not null references Club(club_id)
);

create table Staff(
  staff_id serial primary key,
  staff_name varchar(30) not null,
  staff_surname varchar(30) not null,
  salary decimal check(salary > 30000),
  specification_id int not null references StaffSpecification(specification_id),
  club_id int not null references Club(club_id)
);

create table Stadium(
  stadium_id serial primary key,
  capacity int not null check(capacity > 100),
  stadium_location varchar(150) not null,
  build_date date not null
);

create table Stadium_Club(
  club_id int not null references Club(club_id),
  stadium_id int not null references Stadium(stadium_id)
);

create table Game(
  match_id serial primary key,
  stadium_id int not null references Stadium(stadium_id),
  team_1_id int not null references Team(team_id),
  team_2_id int not null references Team(team_id),
  match_date date not null default current_date
);

create table PlayerStatus(
  status_id serial primary key,
  status_type varchar(20) not null check(status_type in ('в аренде', 'на контракте')),
  start_date date not null,
  expiration_date date not null
);

create table Player(
  player_id serial primary key,
  player_name varchar(30) not null,
  player_surname varchar(30) not null,
  player_number int not null,
  salary decimal not null check(salary > 30000),
  phone varchar(20) not null,
  birth_date date not null,
  team_id int not null references Team(team_id),
  status_id int not null references PlayerStatus(status_id)
);

create table Coach(
  coach_id serial primary key,
  coach_name varchar(30) not null,
  coach_surname varchar(30) not null,
  salary decimal not null,
  phone varchar(20) not null,
  team_id int not null references Team(team_id)
);

create table Coach_Specification(
  specification_id int not null references CoachSpecification(specification_id),
  coach_id int not null references Coach(coach_id)
);

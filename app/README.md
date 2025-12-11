# Программа, взимодействующая с БД

## Описание

Проект представляет собой микросервисное приложение для управления футбольными клубами, командами и тренерами, включающее REST-API, работу с PostgreSQL и автоматические миграции. Сервис обеспечивает создание, обновление и получение данных, соблюдая целостность и структуру доменной модели.

## Быстрый старт
```bash
make up
```

## Запуск

Клонирование репозитория:
```bash
git clone https://github.com/Dmitry-Alekseev01/hse-db-project
cd app
```

Генерация Swagger-документации:
```bash
swag init -g cmd/server/main.go -o docs
```

Создание .env-файла:
```bash
cp .env.example .env
```

Запуск:
```bash
docker compose up -d
```

Завершение:
```bash
docker compose down
```

## Структура проекта
```bash
├── cmd/
│   └── server/
│       └── main.go
├── config/
│   └── config.go
├── docs/
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── internal/
│   ├── delivery/
│   │   ├── club_handler.go
│   │   ├── coach_handler.go
│   │   ├── game_handler.go
│   │   ├── player_handler.go
│   │   ├── router.go
│   │   ├── stadium_handler.go
│   │   ├── staff_handler.go
│   │   └── team_handler.go
│   ├── domain/
│   │   ├── interfaces.go
│   │   └── models.go
│   ├── repository/
│   │   ├── club_repo.go
│   │   ├── coach_repo.go
│   │   ├── game_repo.go
│   │   ├── player_repo.go
│   │   ├── stadium_repo.go
│   │   ├── staff_repo.go
│   │   └── team_repo.go
│   └── usecase/
│       ├── club_usecase.go
│       ├── coach_handler.go
│       ├── game_usecase.go
│       ├── player_usecase.go
│       ├── stadium_usecase.go
│       ├── staff_usecase.go
│       └── team_repo.go
├── migrations/
│   ├── 000_create_db.sql
│   ├── 001_init_db.sql
│   ├── 002_fiil_reference_data.sql
│   ├── 003_fill_clubs_teams.sql
│   ├── 004_fill_people.sql
│   └── 005_fill_games.sql
├── .env.example
├── .girignore
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── Makefile
└── README.md
```
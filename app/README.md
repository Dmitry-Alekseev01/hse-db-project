# Программа, взимодействующая с БД

## Быстрый старт
```bash
make up
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
│   │   └── router.go
│   ├── domain/
│   │   ├── interfaces.go
│   │   └── models.go
│   ├── repository/
│   │   ├── coach_repo.go
│   │   └── club_repo.go
│   └── usecase/
│       ├── coach_handler.go
│       └── club_usecase.go
├── migrations/
│   ├── 000_create_db.sql
│   └── 001_init_db.sql
├── .env.example
├── .girignore
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── Makefile
└── README.md
```
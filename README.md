# Effective Mobile Subscription Service

Сервис для управления пользовательскими подписками: добавление, удаление, фильтрация, подсчёт общей стоимости и просмотр всех подписок.

---

## ⚙️ Про рефакторинг

Этот проект является тестовым задание для компании Effective Mobile. После отправки первой версии были получены замечания, которые были исправлены в этой версии. Я не смог снова связаться с компанией, так как они не допускают правку кода после первой попытки, но все-таки решил закончить проект. 🐸

---



## 🚀 Быстрый старт

### 🔧 Запуск через Docker Compose

```bash
cd docker
docker compose up -d (Обязательно сначала создайте все .env файлы!!!)
```

Приложение будет доступно по адресу:
`http://localhost:8082`

Swagger-документация:
`http://localhost:8082/swagger/index.html`

---

## ⚙️ Переменные окружения

Пример `.env` файла для самого приложение:

```
#<-------------------------->
#PostgreSql Conn

DB_URL="postgres://username:password@database:db_port/db_name"

#<-------------------------->


#<-------------------------->
#App Conn
IP=ip
PORT=port
#<-------------------------->


```

Пример `.env` файла для docker-compose (postgres):

```
#<-------------------------->
#PostgreSql

POSTGRES_USER=username
POSTGRES_PASSWORD=password
POSTGRES_DB=db_name
POSTGRES_PORT=db_port

#<-------------------------->

#<-------------------------->
#App

APP_PORT=app_port
TZ=tz (example: Europe/Moscow)
LOG_LEVEL=lvl
#<-------------------------->


#<-------------------------->
#Migrate

DB_URL="postgres://username:password@database:db_port/db_name?sslmode=disable"

#<-------------------------->

```

---

## 📂 Структура проекта

```
.
├── cmd/           # main.go — точка входа
├── docker/        # Dockerfile и docker-compose
│	└── .env (Вы должны его создать!!) # Зависимости для postgres
├── handlers/      # HTTP endpoints
├── logger/        # Логирование
├── logs/          # Автоматически сохраняемые логи
├── models/        # Структуры данных
├── repository/    # SQL-запросы, работа с базой
├── router/        # Маршруты Gin
├── service/       # Бизнес-логика
├── .env (Вы должны его создать!!) # Зависимости для самой программы
├── .gitignore
├── go.mod
├── go.sum
├── LICENSE
├── README.md
```

---

## 🧩 Используемые технологии

- Go
- Gin
- PostgreSQL
- Swagger (`swaggo`)
- Docker + docker-compose
- zerolog

---

## 📊 Описание эндпоинтов

| Метод | Путь                     | Описание                                                                               | Пример                                                                                                                                                                                                                                                                     |
| ---------- | ---------------------------- | ---------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| POST       | /api/v1/create-subs          | Создать подписку                                                                | curl -X POST http://localhost:8082/api/v1/create-subs   -H "Content-Type: application/json"   -d '{<br />"service_name": "Netflix",<br />"price": 499,<br />"user_id": "60601fee-2bf1-4721-ae6f-7636e79a0cba",<br />"start_date": "09-2022",<br />"end_date": "12-2022"<br />}'  |
| PUT        | /api/v1/update-subs/:user_id | Обновить подписку                                                              | curl -X POST http://localhost:8082/api/v1/create-subs   -H "Content-Type: application/json"   -d '{<br />"service_name": "Spotify",<br />"price": 1000,<br />"user_id": "60601fee-2bf1-4721-ae6f-7636e79a0cba",<br />"start_date": "09-2022",<br />"end_date": "12-2022"<br />}' |
| GET        | /api/v1/get-subs/:user_id    | Получить подписку по user_id                                                 | curl "http://localhost:8082/api/v1/get-subs/60601fee-2bf1-4721-ae6f-7636e79a0cba"                                                                                                                                                                                                |
| DELETE     | /api/v1/delete-subs/:user_id | Удалить подписку по user_id                                                   | curl -X DELETE "http://localhost:8082/api/v1/delete-subs/60601fee-2bf1-4721-ae6f-7636e79a0cba"                                                                                                                                                                                   |
| GET        | /api/v1/get-subs             | Получить все подписки                                                       | curl "http://localhost:8082/api/v1/get-subs"                                                                                                                                                                                                                                     |
| GET        | /api/v1/get-sum-subs         | Получить сумму с подписок<br />в указанном диапазоне. | curl "http://localhost:8082/api/v1/get-sum-subs?user_id=user_id": "60601fee-2bf1-4721-ae6f-7636e79a0cba"                                                                                                                                                                         |

---

## 🐛 Логирование

- Логи хранятся в каталоге `logs/`
- Один файл — на каждый день

## ℹ️ Уровни логирования

```
const (
    TraceLevel = iota  // 0
    DebugLevel         // 1
    InfoLevel          // 2
    WarnLevel          // 3
    ErrorLevel         // 4
    FatalLevel         // 5
    PanicLevel         // 6
    Disabled           // 7 (отключает всё)
) 
```

---

## 📄 Лицензия

MIT License. См. файл `LICENSE`.

---

## 📄 Автор

Ummuys / Егоров Евгений

# Task Service

HTTP API сервис на Go без сторонних сервисов. Задачи создаются, выполняются асинхронно, и хранятся в памяти.

### Что делает сервис

* Позволяет создать задачу (`POST /tasks`)
* Выполняет задачу в фоне (эмуляция 3-5 минут ожидания)
* Позволяет узнать статус и результат (`GET /tasks/{id}`)
* Позволяет удалить задачу (`DELETE /tasks/{id}`)


## Установка и запуск

1. Должен быть установлен Go >= 1.22
2. Необходимо клонировать или скачать репозиторий
github.com/iktkhor/task-service
3. Перейдите в папку и запустите сервер:

```bash
# Путь указан от корня репозитория
cd cmd/server

go run .
```

---

## Ручное тестирование

Сервер работает на `http://localhost:8080`.

### Создать задачу

```bash
curl -X POST http://localhost:8080/tasks
```

### Получить статус задачи

**id** - это уникальный идентификатор, получаемый при создании задачи

```bash
curl http://localhost:8080/tasks/{id}
```

### Удалить задачу

```bash
curl -X DELETE http://localhost:8080/tasks/{id}
```
# Currency Updater

Сервис на Go раз в заданный интервал получает курс доллара к рублю
из HTTP-источника и сохраняет результат в PostgreSQL.

## Возможности

- периодический опрос источника курса;
- пул воркеров для обработки источников;
- отдельная конфигурация для источников с авторизацией;
- Graceful shutdown при `SIGINT` или `SIGTERM`;
- CLI с настройкой числа воркеров.
- простое добавление новых источников курса через общий интерфейс.

## Добавление источника

Новый источник реализуется в пакете `internal/source` и должен соответствовать
интерфейсу `source.Source`:

```go
type Source interface {
    ParseCurrency() *Message
}
```

Для этого создайте структуру источника, реализуйте метод `ParseCurrency()` и
верните из него `*source.Message` с курсом, валютой, датой, названием источника
и возможной ошибкой запроса.

Конфигурация источников задается через `.env`. Например, API-ключ
`SOURCE_WITH_CREDS` загружается в `config.SourceConfig` и передается в метод
`init()` вместе с остальными настройками:

```go
source.NewWithCredSource(cnf.SourceWithCredsApiKey)
```

После реализации подключите источник в методе `init` файла
`internal/service/source_channel_filler.go`, добавив его в возвращаемый срез:

```go
return []source.Source{
    source.NewFrankfurterSource(sc.ctx, sc.client),
    source.NewMySource(sc.ctx, sc.client),
}
```

## Требования

- Go 1.27 или новее;
- PostgreSQL.

## Настройка

1. Скопируйте файл с примером переменных окружения:

   ```bash
   cp .env.example .env
   ```

2. Укажите параметры PostgreSQL и источников в `.env`.

3. Создайте таблицу для сохранения курсов:

   ```sql
   CREATE TABLE IF NOT EXISTS currency (
       data DOUBLE PRECISION NOT NULL,
       currency TEXT NOT NULL,
       source TEXT NOT NULL,
       created_at TIMESTAMP NOT NULL
   );
   ```

## Переменные окружения

| Переменная | Назначение | Значение по умолчанию |
|---|---|---|
| `DB_HOST` | хост PostgreSQL | `localhost` |
| `DB_PORT` | порт PostgreSQL | `5432` |
| `DB_PASSWORD` | пароль пользователя базы данных | обязательно |
| `DB_NAME` | имя базы данных | `main` |
| `UPDATE_PERIOD_SEC` | интервал опроса в секундах | `60` |
| `SOURCE_WITH_CREDS` | API-ключ для источника с авторизацией | обязательно |
| `APP_WORKERS` | число воркеров пула | `4` |

## Запуск

```bash
go run ./cmd
```

Остановка: `Ctrl+C`. Приложение завершит работу корректно после остановки воркеров.

Показать параметры CLI:

```bash
go run ./cmd --help
```
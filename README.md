# CarsCatalog

Решение тестового задания от EffectiveMobile: написание API по [ТЗ](docs%2FTASK.md)

## How to run

### Requirements

* **_Local running_**: Go v1.22 and started PostgreSQL server with "cars" db
* **_Running in Docker_**: Docker with Docker Compose

### Docker run

```
cp .docker.env.example .docker.env
docker-compose up -d
```
По умолчанию
```
docker-compose down
```
### Local run
1. Создать **.env** файл по примеру
    ```
    cp .env.example .env
    ```
2. Убедиться, что PostgreSQL запущен. Внести данные для подключения в **.env** файл
3. Установить зависимости
    ```
   go mod tidy
    ```
4. Запустить приложение
    ```
    make run
    ```
    или собрать в бинарный файл и запустить
    ```
    make build
    chmod +x ./build/app
    ./build/app serve
    ```

## How to use

Все REST методы описаны в [Swagger.json](docs%2Fapi%2Fweb%2Fswagger.json) и [Swagger.yaml](docs%2Fapi%2Fweb%2Fswagger.yaml) файлах.
Также, при запуске приложения **Swagger** доступен как endpoint **/api/v1/swagger**

При создании автомобилей, информация по ним запрашивается из внешнего [API (Swagger)](docs%2Fexternal_api_swagger.yml). В случае, если API вернет информацию о пользователе, не зарегистрированном в БД, то приложение сохраняет его.


## Make commands
```
# runs application by 'go run'
make run

# runs test external API
make run-external-API

# install dev tools(wire, swag, ginkgo)
make install-tools

# build application
make build

# run all unit tests
make test-unit

# run go generate
make gen

# generate OpenAPI docs with swag
make swagger

# generate dependencies with wire
make deps
```

## Project structure

```
├── cmd
│   └── app - точка входа приложения
├── docs - документация
│   └── api
│       └── web - сгенерированные swagger файлы
├── internal
│   ├── app - сборка приложения
│   │   ├── build
│   │   ├── cli
│   │   ├── dependencies - зависимости
│   │   └── initializers - инициализаторы компонентов
│   ├── client_errors - кастомные ошибки
│   ├── configs 
│   ├── external_api
│   ├── gateways
│   │   └── web - http роутер и контроллеры
│   ├── interfaces - интерфейсы для соединения компонентов
│   ├── models - модели сущностей
│   ├── repositories - слой работы с хранилищем
│   └── services - слой бизнес-логики
├── pkg
│   └── logger - вспомогательные структуры и функции для логов
└── test - тесты и тестовое внешнее API, сгенерированное по ТЗ

```


## Tools and packages
* gin-gonic 
* ginkgo with gomega 
* spf13/viper
* spf13/cobra
* envy
* zerolog
* wire
* swag
* docker with docker-compose
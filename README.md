# CarsCatalog

Решение тестового задания от EffectiveMobile: написание API по [ТЗ](docs%2FTASK.md)

## How to run

## How to use

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

# run all tests
make test-all

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
├── db
│   └── migrations - миграции
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
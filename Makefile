## Simple projects tooling for every day

## Project name and source directory path
export DIR  := $(strip $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST)))))

## Creating .env file from template, if file not exists
ifeq ("$(wildcard $(DIR)/.env)","")
  RSP1      := $(shell cp -v $(DIR)/.example_env $(DIR)/.env)
endif

include $(DIR)/.env

## Сценарий по умолчанию - отображение доступных команд
default: help

# Свагер
# swag i --parseVendor --parseDependency -o template/swagger;
swag:
	@swag i -o template/swagger;
	@rm template/swagger/docs.go;
	@rm template/swagger/swagger.yaml;
.PHONY: swag

# FMT & GOIMPORT
fmt:
	@cd $(DIR)/src && go fmt ./... && goimports -w .
.PHONY: fmt

# Linters
lint:
	@cd $(DIR)/src && golangci-lint run
.PHONY: lint

# Test
test:
	export CONF="$(DIR)/conf/config.yaml" && go test ./...
.PHONY: test

# Сборка
com:
	@cd $(DIR) && go mod vendor
	@cd $(DIR) && go build -i -mod vendor -o $(DIR)/bin/app;
.PHONY: com

# Запуск в режиме разработки
run: com
	$(DIR)/bin/app -c conf/config.yaml;
.PHONY: run

# Запуск в режиме отладки
dev: fmt lint test com
	$(DIR)/bin/app -c conf/config.yaml;
.PHONY: dev

# database dump
dbdump:
	pg_dump -F p -h $(PG_HOST) -p $(PG_PORT) -U $(PG_USER) -w -f bin/dump.sql -d $(PG_NAME)
.PHONY: dbdump

# database restore
dbinit:
	psql -h "$(PG_HOST)" -p "$(PG_PORT)" -U $(PG_USER) -w -f bin/dump.sql -d $(PG_NAME)
.PHONY: dbinit

# генерация типов по БД
dbtype:
	cd $(DIR) && go run cmd/gendb/main.go;
	cd $(DIR)/src && go fmt ./typ && goimports -w typ;
.PHONY: dbtype

# Создание шаблона миграции
mig:
	@gsmigrate --dir=${PG_DIR} --drv="postgres" --dsn=${PG_DSN} create tpl;
.PHONY: mig

# Статус миграции
mig-st:
	gsmigrate --dir=${PG_DIR} --drv="postgres" --dsn=${PG_DSN} status;
.PHONY: mig-st

# Миграция на одну позицию вниз
mig-down:
	gsmigrate --dir=${PG_DIR} --drv="postgres" --dsn=${PG_DSN} down;
.PHONY: mig-down

# Миграция вверх до конца
mig-up:
	gsmigrate --dir=${PG_DIR} --drv="postgres" --dsn=${PG_DSN} up;
.PHONY: mig-up

# Инженеринг типов proto
# @cd $(DIR) && protoc --proto_path=pb -I=thirdparty --go_out=plugins=grpc:pb --grpc-gateway_out=logtostderr=true:pb --swagger_out=logtostderr=true,allow_merge=true:pb pb/*.proto
pb:
	@cd $(DIR) && protoc --proto_path=pb -I=thirdparty --go_out=plugins=grpc:pb --grpc-gateway_out=logtostderr=true:pb pb/*.proto
.PHONY: pb

## Clearing project temporary files
clean:
	@GOPATH="$(DIR)" go clean -cache
	@rm -rf ${DIR}/bin/*; true
	@rm -rf ${DIR}/run/*.pid; true
	@rm -rf ${DIR}/log/*.log; true
	@rm -rf ${DIR}/rpmbuild; true
	@rm -rf ${DIR}/*.log; true
	@export DIR=
.PHONY: clean

## Help for main targets
help:
	@echo "Usage: make [target]"
	@echo "  target is:"
	@echo "    dep                  - Загрузка и обновление зависимостей проекта"
	@echo "    dep-dev              - Загрузка и обновление зависимостей проекта для среды разработки"
	@echo "    gen                  - Кодогенерация с использованием go generate"
	@echo "    build                - Компиляция приложения"
	@echo "    run                  - Запуск приложения в продакшн режиме"
	@echo "    dev                  - Запуск приложения в режиме разработки"
	@echo "    kill                 - Отправка приложению сигнала kill -HUP, используется в случае зависания"
	@echo "    m-[driver]-[command] - Работа с миграциями базы данных"
	@echo "                           Используемые базы данных (driver) описываются в файле .env"
	@echo "                           Доступные драйвера баз данных: mysql clickhouse sqlite3 postgres redshift tidb"
	@echo "                           Доступные команды: up, down, create, status, redo, version"
	@echo "                           Пример команд при включённой базе данных mysql:"
	@echo "                             make m-mysql-up      - примернение миграций до самой последней версии"
	@echo "                             make m-mysql-down    - отмена последней миграции"
	@echo "                             make m-mysql-create  - создание нового файла миграции"
	@echo "                             make m-mysql-status  - статус всех миграций базы данных"
	@echo "                             make m-mysql-redo    - отмена и повторное применение последней миграции"
	@echo "                             make m-mysql-version - отображение версии базы данных (применённой миграции)"
	@echo "                           Подробная информаци по командам доступна в документации утилиты gsmigrate"
	@echo "    version              - Вывод на экран версии приложения"
	@echo "    rpm                  - Создание RPM пакета"
	@#echo "    bench                - Запуск тестов производительности проекта"
	@#echo "    test                 - Запуск тестов проекта"
	@#echo "    cover                - Запуск тестов проекта с отображением процента покрытия кода тестами"
	@#echo "    lint                 - Запуск проверки кода с помощью gometalinter"
	@echo "    clean                - Очистка папки проекта от временных файлов"
.PHONY: help

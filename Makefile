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

# Сваггер
swag:
	# swag i --parseVendor --parseDependency -o template/swagger;
	swag i -o template/swagger;
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

# database restore
dbinit:
	psql -h "$(PG_HOST)" -p "$(PG_PORT)" -U $(PG_USER) -w -f bin/dump.sql -d $(PG_NAME)
.PHONY: dbinit

# database dump
dbdump:
	pg_dump -F p -h $(PG_HOST) -p $(PG_PORT) -U $(PG_USER) -w -f bin/dump.sql -d $(PG_NAME)
.PHONY: dbdump

# генерация типов по БД
db:
	cd $(DIR) && go run cmd/gendb/main.go;
	cd $(DIR)/src && go fmt ./typ && goimports -w typ;
.PHONY: db

# Инженеринг типов proto
# @cd $(DIR) && protoc --proto_path=pb -I=thirdparty --go_out=plugins=grpc:pb --grpc-gateway_out=logtostderr=true:pb --swagger_out=logtostderr=true,allow_merge=true:pb pb/*.proto
pb:
	@cd $(DIR) && protoc --proto_path=pb -I=thirdparty --go_out=plugins=grpc:pb --grpc-gateway_out=logtostderr=true:pb pb/*.proto
.PHONY: pb

# Инженеринг моделей по существующей структуре БД
mdcar:
	@go run models/generate/main0.go -md mdcar -pb pbcar
	@sqlboiler -c models/sqlboiler_car.yaml -p mdcar -o models/mdcar --no-auto-timestamps --no-tests --wipe psql
	@go run models/generate/main1.go -md mdcar -pb pbcar
	@go run models/generate/main2.go -md mdcar -pb pbcar
	@goimports -w .
	@go run models/generate/main3.go -md mdcar -pb pbcar
	@protoc -I ./ --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative models/pbcar/*.proto;
	@go fmt ./... && goimports -w .
.PHONY: mdcar

# Help
h:
	@echo "Usage: make [target]"
	@echo "  target is:"
	@echo "    swag		- Генерация документации swagger"
	@echo "    fmt			- Форматирование кодовой базы"
	@echo "    lint		- Линтеры"
	@echo "    test		- Тесты"
	@echo "    run			- Запуск в режиме разработки"
	@echo "    dev			- Запуск в режиме отладки"
	@echo "    mig			- Создание шаблона миграции"
	@echo "    mig-st		- Статус миграции"
	@echo "    mig-down		- Миграция на одну позицию вниз"
	@echo "    mig-up		- Миграция вверх до конца"
	@echo "    dbinit		- Восстановление БД из дампа bin/dump.sql (БД должна существовать)"
	@echo "    dbdump		- Создание дампа БД bin/dump.sql"
	@echo "    db:			- Инженеринг типов по БД"
	@echo "    pb			- Инженеринг GRPC"

.PHONY: h
help: h
.PHONY: help

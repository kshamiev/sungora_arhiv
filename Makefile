# Инициализация
ifeq ("$(wildcard .env)","")
  RSP1      := $(shell cp -v .example_env .env)
endif

include .env

default: help

# Зависимости
dep:
	go mod tidy
	go mod vendor
.PHONY: dep

# Сваггер
swag:
	swag i --parseVendor -o internal/config;
	rm -f internal/config/swagger.json
	rm -f internal/config/swagger.yaml
.PHONY: swag

# FMT & GOIMPORT
fmt:
	go fmt ./... && goimports -w .
.PHONY: fmt

# Linters
lint:
	golangci-lint run -c .golangci.yml
.PHONY: lint

# Test
test:
	go test ./...
.PHONY: test

# Сборка
com:
	go build -o bin/app .;
.PHONY: com

# Запуск в режиме разработки
run: com
	bin/app -c config.yml;
.PHONY: run

# Запуск в режиме отладки
dev: dep swag fmt lint test com
	rm -rd vendor
	bin/app -c config.yml;
.PHONY: dev

# Создание шаблона миграции
mig:
	@gsmigrate --dir=${PG_DIR} --drv="postgres" --dsn=${PG_DSN} create tpl;
.PHONY: mig

# Статус миграции
mig-st:
	@gsmigrate --dir=${PG_DIR} --drv="postgres" --dsn=${PG_DSN} status;
.PHONY: mig-st

# Миграция на одну позицию вниз
mig-down:
	@gsmigrate --dir=${PG_DIR} --drv="postgres" --dsn=${PG_DSN} down;
.PHONY: mig-down

# Миграция вверх до конца
mig-up:
	@gsmigrate --dir=${PG_DIR} --drv="postgres" --dsn=${PG_DSN} up;
.PHONY: mig-up

# database restore
dbinit:
	@psql -h "$(PG_HOST)" -p "$(PG_PORT)" -U $(PG_USER) -w -d $(PG_NAME) -f bin/dump.sql
.PHONY: dbinit

# database full dump
dbdump:
	@pg_dump -F p -h $(PG_HOST) -p $(PG_PORT) -U $(PG_USER) -w -d $(PG_NAME) -f bin/dump.sql
.PHONY: dbdump

# database schema dump
dbdump-s:
	@pg_dump -F p -h $(PG_HOST) -p $(PG_PORT) -U $(PG_USER) -s -w -d $(PG_NAME) -f bin/dump.sql
.PHONY: dbdump-s

# database data dump
dbdump-a:
	@pg_dump -F p -h $(PG_HOST) -p $(PG_PORT) -U $(PG_USER) -a -w -d $(PG_NAME) -f bin/dump.sql
.PHONY: dbdump-a

# Инженеринг моделей по существующей структуре БД

SERVICE1 := sample
ser-sample:
	@go run services/generate/main.go -step $(SERVICE1)-1
	@sqlboiler -c etc/config.yml -p md$(SERVICE1) -o services/md$(SERVICE1) --no-auto-timestamps --no-tests --wipe psql
	@go run services/generate/main.go -step $(SERVICE1)-2
	@go run services/generate/main.go -step $(SERVICE1)-3
	@cd $(DIR)/services && goimports -w .
	@go run services/generate/main.go -step $(SERVICE1)-4
	@cd $(DIR)/services && protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb$(SERVICE1)/*.proto;
	@cd $(DIR)/services && go fmt ./... && goimports -w .
.PHONY: ser-sample

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
	@echo "    mig-dn		- Миграция на одну позицию вниз"
	@echo "    mig-up		- Миграция вверх до конца"
	@echo "    dbinit		- Восстановление БД из дампа bin/dump.sql (БД должна существовать)"
	@echo "    dbdump		- Создание дампа БД bin/dump.sql"
	@echo "    ser-sample:		- Инженеринг типов по БД и работа с GRPC в парадигме масштабируемого сервиса"

.PHONY: h
help: h
.PHONY: help

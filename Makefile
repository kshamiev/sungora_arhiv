# Инициализация
ifeq ("$(wildcard .env)","")
  RSP1      := $(shell cp -v .example_env .env)
endif

include .env

default: help

# Сваггер
swag:
	# swag i --parseVendor --parseDependency -o template/swagger;
	swag i --parseVendor -o template/swagger;
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
	go build -o $(DIR)/bin/app $(DIR);
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
	psql -h "$(PG_HOST)" -p "$(PG_PORT)" -U $(PG_USER) -w -d $(PG_NAME) -f bin/dump.sql
.PHONY: dbinit

# database full dump
dbdump:
	pg_dump -F p -h $(PG_HOST) -p $(PG_PORT) -U $(PG_USER) -w -d $(PG_NAME) -f bin/dump.sql
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
# @protoc -I ./ --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative models/pbcar/*.proto;
mdsun:
	@go run types/generate/main0.go -md mdsun -pb pbsun
	@sqlboiler -c conf/sqlboiler_sun.yaml -p mdsun -o types/mdsun --no-auto-timestamps --no-tests --wipe psql
	@go run types/generate/main1.go -md mdsun -pb pbsun
	@go run types/generate/main2.go -md mdsun -pb pbsun
	@goimports -w .
	@go run types/generate/main3.go -md mdsun -pb pbsun
	@protoc -I=thirdparty --proto_path=./ --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative types/pbsun/*.proto;
	@cd $(DIR)/types/mdsun && go fmt ./... && goimports -w .
.PHONY: mdsun

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

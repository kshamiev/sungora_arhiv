# Sample go application
---

- **/api**
  исходный код конкретного приложения (бизнес логика, бизнес функционал и api)
- **/app**
  исходный код конкретного приложения (инициализация нужного инструментария, взаимодействие с внешним миром, ...)
- **/bin**
  бинарник приложения
- **/etc**
  сертификаты, конфиги и щаблоны приложения или проекта
- **/doc**
  исключительно документация, туда же должна генерируя вся документация создаваемая автоматически.
- **/lib**
  общие библиотеки и инструментарий проекта (выносится в общую репу проекта).
- **/migrations**
  миграции БД приложения
- **/services**
  описание сервисов grpc и типы доступные между сервисами - приложениями проекта (выносится в общую репу проекта).
- **/vendor**
  сторонние библиотеки используемые в приложении
- **/www**
  клиентская часть папка может приезжать из другого репозитория (html, css, js, img)

### profile

    // test
    go build -gcflags=-m
    go test -bench=. -benchmem
    go test -bench=. -cpuprofile cpu.out -memprofile mem.out
    go tool pprof ... cpu.out || mem.out

    // service
    http://localhost:8080/api/sun/debug/pprof/index

    ab -n 100000 -c 10 http://localhost:8080/api/sun/user-test/d9f982ee-0bf3-4a2e-9b7f-0c571ac7d253
    ab -k -n 100000 -c 8 http://localhost:8080/api/sun/general/ping

    curl http://localhost:8080/api/sun/debug/pprof/trace?seconds=10 -o trace.out
    go tool trace -http "0.0.0.0:8080" ./tracetest trace.out

    go tool pprof http://localhost:8080/api/sun/debug/pprof/profile?seconds=10
    go tool pprof http://localhost:8080/api/sun/debug/pprof/allocs
    go tool pprof -alloc_objects http://localhost:8080/api/sun/debug/pprof/allocs
    go tool pprof http://localhost:8080/api/sun/debug/pprof/heap
    go tool pprof -inuse_objects http://localhost:8080/api/sun/debug/pprof/heap
    go tool pprof http://localhost:8080/api/sun/debug/pprof/goroutine

    -alloc_objects просмотр количества выделенных объектов на протяжении всего жизненного цикла приложения.
    -alloc_space размер выделенной памяти на протяжении всего жизненного цикла приложения.
    -inuse_objects просмотр количества объектов, используемых во время профилирования
    -inuse_space просмотр объема памяти, используемой во время профилирования

### DOCKER

#### Graylog

```dockerfile
docker run --rm --name mongo -d mongo:4.2

docker run --rm --name elasticsearch \
    -e "http.host=0.0.0.0" \
    -e "discovery.type=single-node" \
    -e "ES_JAVA_OPTS=-Xms512m -Xmx512m" \
    -d elasticsearch:7.9.2
    -d elasticsearch:7.16.2
    -d docker.elastic.co/elasticsearch/elasticsearch-oss:7.10.2

docker run --rm --name graylog --link mongo --link elasticsearch \
    -p 9000:9000 -p 12201:12201 -p 1514:1514 \
    -e GRAYLOG_HTTP_EXTERNAL_URI="http://127.0.0.1:9000/" \
    -e GRAYLOG_PASSWORD_SECRET="yESMn44eQlUXTqN2RyXqJzbxatiHZwGZwWJu9pUHNOQAQQm1NmKiQwtk7l5u6pC0m7ub6ilyFh0YqepA9" \
    -e GRAYLOG_ROOT_PASSWORD_SHA2="65e84be33532fb784c48129675f9eff3a682b27168c0ea744b2cf58ee02337c5" \
    -d graylog/graylog:4.2
```

http://localhost:9000/

    admin
    qwerty

#### Jaeger

https://www.jaegertracing.io/docs/1.20/getting-started/

```dockerfile
docker run -d --rm --name sample-jaeger --net sun \
    -p 16686:16686 -p 14268:14268 \
    jaegertracing/all-in-one:1.20
```

http://localhost:16686

#### Minio

```dockerfile
docker run -d --rm --name sample-minio --net sun \
    -p 9020:9000 -p 9010:9001 \
    -e MINIO_ROOT_USER="admin" -e MINIO_ROOT_PASSWORD="xxx-xxx-xxx" \
    -v /mnt/data/sample:/data \
    minio/minio \
    server /data --address ":9020" --console-address ":9010"
```

http://localhost:9010

    guest
    guest

#### RabbitMQ

```dockerfile
docker run -d --rm --name sample-rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.9-management
```

http://localhost:15672/#/

#### Сборка и запуск приложения

```dockerfile
docker build --no-cache -t sample .

docker run --rm -d --name sample --net sun\
    -p 8080:8080 \
    --mount type=bind,src=/home/domains/sample.local/www,dst=/home/app/www \
    sample
```

### TODO or TASK

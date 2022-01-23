# Sungora
---

- **/api**
  исходный код конкретного приложения (бизнес логика, бизнес функционал и api)
- **/app**
  исходный код конкретного приложения (инициализация нужного инструментария, взаимодействие с внешним миром, ...)
- **/bin**
  бинарник приложения
- **/conf**
  шаблоны и конфиги приложения или проекта
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

    go tool pprof http://localhost:8080/debug/pprof/profile?seconds=30
    go tool pprof http://localhost:8080/debug/pprof/allocs
    go tool pprof http://localhost:8080/debug/pprof/heap
    go tool pprof http://localhost:8080/debug/pprof/goroutine

### DOCKER

#### Graylog

```dockerfile
docker run --name graylog --link mongo --link elasticsearch \
    -p 9000:9000 -p 12201:12201 -p 1514:1514 \
    -e GRAYLOG_HTTP_EXTERNAL_URI="http://127.0.0.1:9000/" \
    -e GRAYLOG_PASSWORD_SECRET="yESMn44eQlUXTqN2RyXqJzbxatiHZwGZwWJu9pUHNOQAQQm1NmKiQwtk7l5u6pC0m7ub6ilyFh0YqepA9" \
    -e GRAYLOG_ROOT_PASSWORD_SHA2="65e84be33532fb784c48129675f9eff3a682b27168c0ea744b2cf58ee02337c5" \
    -d graylog/graylog:4.2
```

http://127.0.0.1:9000/

    admin
    qwerty

#### Jaeger

https://www.jaegertracing.io/docs/1.20/getting-started/

```dockerfile
docker run -d --rm --name sungora-jaeger --net sun \
    -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
    -p 5775:5775/udp -p 6831:6831/udp -p 6832:6832/udp \
    -p 5778:5778 -p 16686:16686 -p 14268:14268 -p 14250:14250 -p 9411:9411 \
    jaegertracing/all-in-one:1.20

docker run -d --rm --name sungora-jaeger --net sun \
    -p 127.0.0.1:16686:16686 -p 127.0.0.1:14268:14268 \
    jaegertracing/all-in-one:1.20
```

http://localhost:16686

#### Minio

```dockerfile
docker run -d --rm --name sungora-minio --net sun \
    -p 127.0.0.1:9000:9000 -p 127.0.0.1:9001:9001 \
    -e MINIO_ROOT_USER="admin" -e MINIO_ROOT_PASSWORD="xxx-xxx-xxx" \
    -v /mnt/data/sungora:/data \
    minio/minio \
    server /data --address ":9000" --console-address ":9001"
```

http://localhost:9001

#### Сборка и запуск приложения

```dockerfile
docker build --no-cache --rm -t kshamiev/sungora .

docker run --rm -d --name sungora --net sun\
    -p 127.0.0.1:8080:8080 \
    --mount type=bind,source=/home/domains/sungora.local/www,target=/home/app/www \
    kshamiev/sungora
```

### TODO or TASK

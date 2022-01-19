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

### Jaeger

https://www.jaegertracing.io/docs/1.20/getting-started/

```dockerfile
docker run -d --rm --name jaeger \
    -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
    -p 5775:5775/udp -p 6831:6831/udp -p 6832:6832/udp \
    -p 5778:5778 -p 16686:16686 -p 14268:14268 -p 14250:14250 -p 9411:9411 \
    jaegertracing/all-in-one:1.20
```

http://localhost:16686

### Minio

```dockerfile
docker run -d --rm --name minio-sungora \
    -p 9000:9000 -p 9001:9001 \
    -e MINIO_ROOT_USER="admin" -e MINIO_ROOT_PASSWORD="Cf5IttjOxXnl" \
    -v /mnt/data/sungora:/data \
    minio/minio \
    server /data --address ":9000" --console-address ":9001"
```

Сборка

```dockerfile
docker build --no-cache -t kshamiev/sun:v1 .
```

Запуск

```dockerfile
docker run --rm -d --name sun1 \
    -p 127.0.0.1:8080:8080 -p 127.0.0.1:7071:7071 -p 127.0.0.1:9000:9000 -p 127.0.0.1:14268:14268 \
    --mount type=bind,source=/home/domains/www.funtik.ru/www,target=/home/app/www \
    kshamiev/sun:v1
```

### DOCKER

```dockerfile
docker build --rm -t kshamiev/sungora:v1.10.100 .
docker run -d --rm --net host --name sungora kshamiev/sungora:v1.10.100
```

### TODO or TASK

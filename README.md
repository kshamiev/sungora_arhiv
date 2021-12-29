# Sungora
---

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
- **/src**
  исходный код конкретного приложения
- **/template**
  шаблоны на языке шаблонизатора golang для использования на стороне сервера
- **/thirdparty**
  для генерации grpc кода
- **/vendor**
  сторонние библиотеки используемые в приложении
- **/www**
  клиентская часть папка может приезжать из другого репозитория

### profile

    go tool pprof http://localhost:8080/debug/pprof/profile?seconds=30
    go tool pprof http://localhost:8080/debug/pprof/allocs
    go tool pprof http://localhost:8080/debug/pprof/heap
    go tool pprof http://localhost:8080/debug/pprof/goroutine

### Jaeger

https://www.jaegertracing.io/docs/1.20/getting-started/

Запуск в докере:

docker run -d --rm --name jaeger -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
-p 5775:5775/udp -p 6831:6831/udp -p 6832:6832/udp \
-p 5778:5778 -p 16686:16686 -p 14268:14268 -p 14250:14250 -p 9411:9411 \
jaegertracing/all-in-one:1.20

docker run -d --rm --name jaeger -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 -p 14268:14268 -p 16686:16686 -p 9411:9411 \
jaegertracing/all-in-one:1.20

Просмотр

http://localhost:16686

### DOCKER

    docker build --rm -t kshamiev/sungora:v1.10.100 .
    docker run -d --rm --net host --name sungora kshamiev/sungora:v1.10.100

### TODO or TASK

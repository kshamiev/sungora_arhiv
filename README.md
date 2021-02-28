# Sungora
---

- **/cmd**
вспомогательные утилиты   
- **/conf**
шаблоны и конфиги, включая локальные для запуска проекта локально. Деплоится в /etc/ИМЯ_ПРОЕКТА
- **/doc**
исключительно документация, туда же должна генерируя вся документация создаваемая автоматически.
- **/lib**
либы которые должны быть вынесены в отдельную репу.
здесь пока не будут отработаны.
- **/migrations**
миграции БД
- **/src**
исходный код проекта
- **/template**
папка полностью относится к golang - там шаблоны на языке шаблонизатора golang, она на линуксе деплоится как правило в /usr/share/ИМЯ_ПРОЕКТА/template
- **/www**
клиентская часть папка может приезжать из другого репозитория, например из репозитория в котором всё сделано на ангуляре и на выходе готовое web приложение ангуляра

### profile

    go tool pprof http://localhost:8080/debug/pprof/profile?seconds=30
    go tool pprof http://localhost:8080/debug/pprof/allocs
    go tool pprof http://localhost:8080/debug/pprof/heap
    go tool pprof http://localhost:8080/debug/pprof/goroutine

### Jaeger
https://www.jaegertracing.io/docs/1.20/getting-started/

Запуск в докере:

docker run -d --name jaeger1 \
  -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14268:14268 \
  -p 14250:14250 \
  -p 9411:9411 \
  jaegertracing/all-in-one:1.20

Просмотр

http://localhost:16686

### TODO or TASK

реализация кастомного контекста

FROM kshamiev/assembly:v1 as build1

WORKDIR /home/app

COPY . .

RUN rm -rf /home/app/bin

RUN go build -o bin/app cmd/main.go

FROM kshamiev/service:v1

WORKDIR /home/app

COPY --from=build1 /home/app/bin bin
COPY --from=build1 /home/app/etc etc
COPY --from=build1 /home/app/www_test www_test

RUN mkdir www
VOLUME www

EXPOSE 8080:8080

CMD bin/app -c etc/config_docker.yaml

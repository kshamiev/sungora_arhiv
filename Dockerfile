FROM kshamiev/assembly:v1 as build1

WORKDIR /home/app

COPY . .

RUN rm -rf /home/app/bin

RUN go build -o bin/app .

FROM kshamiev/service:v1

WORKDIR /home/app

COPY --from=build1 /home/app/bin bin
COPY --from=build1 /home/app/etc etc

RUN mkdir www
VOLUME www

EXPOSE 8080:8080

CMD bin/app -c config_docker.yml

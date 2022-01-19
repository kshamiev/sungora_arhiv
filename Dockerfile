FROM kshamiev/assembly:v1 as build1

WORKDIR /home/app

COPY . .

RUN rm -rf /home/app/bin

RUN go build -o bin/app .

FROM kshamiev/service:v1

WORKDIR /home/app

COPY --from=build1 /home/app/bin bin
COPY --from=build1 /home/app/conf conf

EXPOSE 8080:8080
EXPOSE 7071:7071
EXPOSE 9000:9000
EXPOSE 14268:14268

CMD bin/app -c conf/config_docker.yaml

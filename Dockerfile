FROM kshamiev/assembly:v1 as build1

WORKDIR /home/app

COPY . .

RUN rm -rf /home/app/bin

RUN go build -o bin/app .

FROM kshamiev/service:v1

WORKDIR /home/app

COPY --from=build1 /home/app/bin bin
COPY --from=build1 /home/app/conf conf
COPY --from=build1 /home/app/www www

CMD bin/app -c conf/config.yaml

## docker build --no-cache -t kshamiev/sun:v1 .
## docker run --rm -it kshamiev/sun:v1
## docker run --rm -d --net host --name sun1 kshamiev/sun:v1
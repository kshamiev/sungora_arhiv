FROM ubuntu:20.04

WORKDIR /home/app

COPY bin bin
COPY conf conf
COPY template template
COPY www www

EXPOSE 8080
EXPOSE 5432

CMD bin/app -c conf/config.yaml

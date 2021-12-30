FROM ubuntu:20.04

LABEL maintainer="Konstantin Shamiev aka ilosa <konstantin@shamiev.ru>"

ENV ADMIN="Konstantin Shamiev"

WORKDIR /home/app

COPY bin bin
COPY conf conf
COPY template template
COPY www www

RUN mkdir data

EXPOSE 8080
EXPOSE 5432

CMD bin/app -c conf/config.yaml
FROM ubuntu:20.04

LABEL maintainer="Konstantin Shamiev aka ilosa <konstantin@shamiev.ru>"

WORKDIR /home/app

COPY bin bin
COPY conf conf
COPY template template
COPY www www

CMD bin/app -c conf/config.yaml

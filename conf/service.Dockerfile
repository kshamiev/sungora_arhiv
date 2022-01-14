FROM ubuntu:20.04

LABEL maintainer="Konstantin Shamiev aka ilosa <konstantin@shamiev.ru>"

RUN apt-get update

# docker build --no-cache -t kshamiev/service:v1 -f service.Dockerfile .
# docker push kshamiev/service:v1
# docker run --rm -it kshamiev/service:v1
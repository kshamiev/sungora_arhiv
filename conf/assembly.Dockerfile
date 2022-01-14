FROM ubuntu:20.04

LABEL maintainer="Konstantin Shamiev aka ilosa <konstantin@shamiev.ru>"

RUN apt-get update && apt-get install make && apt-get install -y git

RUN rm -rf /usr/local/go
COPY go /usr/local

ENV PATH=$PATH:/usr/local/go/bin

# docker build --no-cache -t kshamiev/assembly:v1 -f assembly.Dockerfile .
# docker push kshamiev/assembly:v1
# docker run --rm -it kshamiev/assembly:v1
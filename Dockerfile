FROM ubuntu:latest
RUN mkdir -p /src
WORKDIR /src/
RUN apt-get update && \
    apt-get install -y golang
COPY . /src/
RUN go build

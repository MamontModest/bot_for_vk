FROM golang:latest

WORKDIR /usr/src/app



COPY src .
RUN mkdir -p /usr/local/bin/

COPY go.mod .
COPY go.sum .
RUN go mod tidy

RUN go build -v -o /usr/local/bin/app

CMD [ "app" ]
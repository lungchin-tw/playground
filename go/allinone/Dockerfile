FROM golang:1.19.13-alpine3.17


COPY src /app/
WORKDIR /app

RUN go mod tidy && go mod verify

RUN go build -v -o /usr/local/bin/app

CMD ["app"]
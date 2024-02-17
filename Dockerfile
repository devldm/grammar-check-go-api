FROM golang:1.22.0-alpine3.19

ENV PORT 8080
ENV HOST 0.0.0.0

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go mod download && go mod verify

RUN go build -o main .

EXPOSE 8080

CMD ["/app/main"]
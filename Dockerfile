FROM golang:1.21

EXPOSE 9000

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main src/main.go

CMD [ "./main" ]




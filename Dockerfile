FROM golang:1.21

WORKDIR /

EXPOSE 9000

RUN go build src/main.go

CMD [ "./main" ]
FROM golang:1.16-alpine

RUN go build main.go

EXPOSE 9000

CMD [ "./main" ]

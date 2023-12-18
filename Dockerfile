FROM golang:1.21

RUN go build

EXPOSE 9000

CMD [ "./main" ]




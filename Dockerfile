FROM golang:1.21

RUN go build /gohttp/src/main.go main

EXPOSE 9000

CMD [ "./gohttp/main" ]




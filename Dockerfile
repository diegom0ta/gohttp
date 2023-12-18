FROM golang:1.21

RUN mkdir /gohttp

ADD . /gohttp

EXPOSE 9000

CMD [ "./gohttp/main" ]




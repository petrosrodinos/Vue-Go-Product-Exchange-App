FROM golang:1.19

LABEL petros rodinos <petros.rodinos@yahoo.com>

WORKDIR /go/src/app
COPY ./ .

RUN go get -d -v
RUN go build -v

EXPOSE 9090

CMD ["./api"]



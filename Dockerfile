FROM golang:latest
WORKDIR /go/src
COPY ./src .
RUN go get -u github.com/cosmtrek/air
CMD ["air", "-c", ".air.toml"]
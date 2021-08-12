FROM golang:latest
WORKDIR /go/src
COPY . .
RUN go get -u github.com/cosmtrek/air
CMD ["air", "-c", ".air.toml"]
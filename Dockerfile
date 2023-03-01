FROM golang:1.18-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o authorizationGolang ./cmd/main.go

CMD ["./authorizationGolang"]
FROM golang:latest

RUN go version
ENV GOPATH=/cmd

EXPOSE 5000

COPY ./ ./

RUN go mod download
RUN go build -o authorizationGolang ./cmd/main.go

CMD ["./authorizationGolang"]
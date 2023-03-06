FROM golang:latest

RUN go version
ENV GOPATH=/cmd

COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql

RUN chmod +x wait-for-postgres.sh

ENTRYPOINT ["/usr/local/bin/wait-for-postgres.sh"]


RUN go mod download
RUN go build -o authorizationGolang ./cmd/main.go

CMD ["./authorizationGolang"]
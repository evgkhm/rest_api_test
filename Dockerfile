FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o rest_api_test ./cmd/main.go

CMD ["./rest_api_test"]
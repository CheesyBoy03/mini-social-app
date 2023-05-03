FROM golang:1.18

RUN go version
ENV GOPATH=/

COPY . .

RUN go mod download
RUN go build -o mini-social-app ./cmd/main.go

RUN chmod +x mini-social-app

CMD ["mini-social-app"]

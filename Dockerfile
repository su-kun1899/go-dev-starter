FROM golang:1.19

WORKDIR /app
RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest

CMD ["air"]

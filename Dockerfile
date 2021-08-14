FROM golang:alpine as builder

WORKDIR /go/src/fpl-find-a-manager
COPY . .

RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o app ./cmd/server/


FROM alpine:latest

WORKDIR /root/
COPY --from=builder /go/src/fpl-find-a-manager/app .

CMD ./app
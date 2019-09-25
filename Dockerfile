FROM golang as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GODS=linux GOARCH=amd64 go build

FROM ubuntu

COPY --from=builder /app/letsgo /app/
COPY --from=builder /app/.env /app/

RUN mkdir /app/log

EXPOSE 8080
ENTRYPOINT ["/app/letsgo"]
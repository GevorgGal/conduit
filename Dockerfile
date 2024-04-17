FROM golang:1.21 as builder

EXPOSE 8080

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/conduit-connector-influxdb ./cmd/conduit-connector-influxdb

FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/bin/conduit-connector-influxdb .

CMD ["./conduit-connector-influxdb"]

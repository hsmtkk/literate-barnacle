FROM golang:1.20 AS builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
WORKDIR /app/back/server/room
RUN CGO_ENABLED=0 go build -o app

FROM gcr.io/distroless/static-debian11 AS runtime
WORKDIR /app
COPY --from=builder /app/back/server/room/app /app/app
ENTRYPOINT ["/app/app"]

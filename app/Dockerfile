# builder
FROM golang:1.26.0 as builder

WORKDIR /app

COPY main.go .

RUN CGO_ENABLED=0 go build -o server main.go

# runner
FROM scratch

COPY --from=builder /app/server /bin/server

EXPOSE 8080

ENTRYPOINT ["/bin/server"]
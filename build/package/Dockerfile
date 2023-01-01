FROM golang:1.17-alpine as builder

WORKDIR /app
COPY ../.. .
RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/main main.go

FROM gcr.io/distroless/static:nonroot

WORKDIR /opt/
COPY --from=builder /app/bin/main .
USER 65532:65532

CMD ["./main"]
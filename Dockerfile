FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY  . .
RUN go build -o todo-app cmd/main.go

FROM alpine as run_stage
WORKDIR /out
COPY --from=builder /app/todo-app ./binary
EXPOSE 8000
CMD ["./binary"]
FROM golang:1.22 AS builder
ENV PROJECT_PATH=/app/tournament_management
ENV CGO_ENABLED=0
ENV GOOS=linux
COPY . ${PROJECT_PATH}
WORKDIR ${PROJECT_PATH}
RUN go build cmd/tournament_management/main.go

FROM golang:alpine
WORKDIR /app/cmd/tournament_management
COPY --from=builder /app/tournament_management/main .
EXPOSE 30001
CMD ["./main"]

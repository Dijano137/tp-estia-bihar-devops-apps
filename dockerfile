FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY . .
RUN go build

FROM scratch
COPY --from=builder /app/backend /backend
EXPOSE 8081
CMD ["/backend"]
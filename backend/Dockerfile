FROM golang:1.18-alpine3.14 as builder
WORKDIR /backend
COPY . .
RUN go build -o main main.go

FROM alpine:3.14
WORKDIR /backend
COPY --from=builder /backend/main .

CMD [ "/backend/main" ]
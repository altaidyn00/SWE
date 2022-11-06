FROM golang:1.18-alpine:3.14 as builder
WORKDIR /hospital
COPY . .
RUN go build -o main main.go

FROM alpine:3.14
WORKDIR /hospital
COPY --from=builder /receiver/main .

CMD [ "/receiver/main" ]
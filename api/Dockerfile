FROM golang:apline as builder

RUN mkdir /build

ADD . /build/

WORKDIR /build

RUN go build -o main .

# stage 2

FROM alpine

RUN addUser -S -D -H -h /app appUser

USER appUser

COPY . /app

COPY --from=builder /build/main /app/

WORKDIR /app

EXPOSE 3000

CMD ["./main"]
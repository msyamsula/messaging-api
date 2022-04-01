FROM ubuntu

WORKDIR /app

COPY messaging-api /app

CMD ["./messaging-api"]



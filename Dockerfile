FROM ubuntu

WORKDIR /app

COPY messaging-api .

CMD ["./messaging-api"]

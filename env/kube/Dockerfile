FROM ubuntu

WORKDIR /app

COPY messaging-api /app
COPY env/kube/.env /app/env/prod/

CMD ["./messaging-api"]



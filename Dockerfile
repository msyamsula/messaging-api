FROM ubuntu

WORKDIR /app

COPY messaging-api .
COPY wait-for-it.sh wait-for-it.sh
RUN chmod +x wait-for-it.sh

CMD ["./messaging-api"]



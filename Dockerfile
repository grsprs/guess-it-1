FROM golang:1.21-alpine

WORKDIR /app

COPY student/ ./student/
COPY script.sh ./

RUN chmod +x script.sh

CMD ["sh", "-c", "tail -f /dev/null"]

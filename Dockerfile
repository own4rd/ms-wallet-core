FROM golang:1.20

WORKDIR /app

RUN apt-get update && apt-get install -y librdkafka-dev

# PRENDER O CONTAINER
CMD ["tail", "-f", "/dev/null"]


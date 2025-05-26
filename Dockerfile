FROM golang:1.23.2

WORKDIR /app

RUN apt-get update && apt-get install -y librdkafka-dev

# PRENDER O CONTAINER
CMD ["tail", "-f", "/dev/null"]


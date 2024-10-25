FROM golang:1.19

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

RUN apt-get update && apt-get install -y librdkafka-dev

CMD ["tail", "-f", "/dev/null"]
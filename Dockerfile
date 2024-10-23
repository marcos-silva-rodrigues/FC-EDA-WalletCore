FROM golang:1.19

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

CMD ["tail", "-f", "/dev/null"]
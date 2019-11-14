FROM golang:latest

WORKDIR /app

COPY . /app

RUN ls .

ENV GOPATH /app
ENV GOBIN $GOPATH/bin
ENV PATH $GOPATH/bin:$PATH

RUN go install app

ENTRYPOINT ./bin/app
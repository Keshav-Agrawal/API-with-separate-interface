# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

RUN apk add --no-cache git

WORKDIR /app

#COPY go.mod ./
#COPY go.sum ./
#RUN go mod download


COPY . ./
RUN go mod download
#RUN go get github.com/Keshav-Agrawal/mongoapi/datasource/mongo

#RUN go get github.com/Keshav-Agrawal/mongoapi/router
RUN go build -o /docker-app

EXPOSE 8080

CMD [ "/docker-app" ]
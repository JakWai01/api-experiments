FROM golang:alpine AS build

RUN apk add git

RUN mkdir /src
ADD . /src
WORKDIR /src

RUN go build -o /tmp/api-experiments ./main.go

FROM alpine:edge

COPY --from=build /tmp/api-experiments /sbin/api-experiments

CMD /sbin/api-experiments
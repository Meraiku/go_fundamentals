FROM golang:1.23-alpine AS build-env
RUN apk add --no-cache make
ADD . /project/src
WORKDIR /project/src
RUN make build




FROM debian:stable-slim

COPY --from=build-env  /project/src/.bin/out ./bin/out
COPY ./certs/* /etc/ssl/certs/

CMD [ "/bin/out" ]

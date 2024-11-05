FROM debian:stable-slim

COPY ./.bin/out ./bin/out
COPY ./sql/migrations ./sql/migrations

CMD [ "/bin/out" ]

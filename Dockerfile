FROM alpine:latest

ENV app=tfc-cap-updater

RUN mkdir /app
WORKDIR /app
ADD $app /app/$app

CMD ./$app
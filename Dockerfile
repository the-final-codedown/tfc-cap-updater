FROM alpine:latest

ENV app=tfc-cap-updater

WORKDIR /app
ADD $app ./$app

CMD ./$app
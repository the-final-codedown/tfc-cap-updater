FROM scratch

ENV app=tfc-cap-cli

RUN mkdir /app
WORKDIR /app
ADD $app /app/$app
ADD cap.json /app/cap.json

CMD ./$app
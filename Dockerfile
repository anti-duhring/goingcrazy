FROM ubuntu:latest

WORKDIR /home/api

COPY goingcrazy /home/api/

CMD [ "./goingcrazy" ]
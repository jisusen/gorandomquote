FROM golang:1.17

WORKDIR /usr/src/app/backend

COPY ./backend/ /usr/src/app/backend/

RUN go build

EXPOSE 8080

CMD ["./backend"]
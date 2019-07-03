FROM golang:1.11

WORKDIR /go/src/github.com/felixgravila/psqlrestapi
COPY . .

RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/lib/pq
RUN go install

EXPOSE 8000

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.5.0/wait /wait
RUN chmod +x /wait

CMD /wait && psqlrestapi

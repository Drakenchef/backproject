FROM golang:1.20.2-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

#install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

#make wait-for-ppostres.sh executable
RUN chmod +x wait-for-postgres.sh


#build go app
RUN go mod download
RUN go build -o backproject-app ./cmd/main.go

CMD ["./backproject-app"]
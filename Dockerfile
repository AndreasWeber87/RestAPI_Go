FROM golang:1.19

WORKDIR /home/ic20b050/app
ADD . /home/ic20b050/app

EXPOSE 9000

CMD go run main.go
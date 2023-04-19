FROM golang:1.20.3-alpine

# set Workdir inside the image
WORKDIR /home/ic20b050/app 
# copy the current dir from the host in the image dir
ADD . /home/ic20b050/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
RUN go mod download && go mod verify
RUN go build -v -o /home/ic20b050/app/GoAPI ./...

# release the port to the host
EXPOSE 7000

CMD ["/home/ic20b050/app/GoAPI"]

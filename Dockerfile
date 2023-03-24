FROM golang:1.20.2-alpine

# Workdir innerhalb des Containers festlegen
WORKDIR /home/ic20b050/app 
# Kopiert das aktuelle Verzeichnis vom Host in das Image Verzeichnis
ADD . /home/ic20b050/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
RUN go mod download && go mod verify
RUN go build -v -o /home/ic20b050/app/restapi ./...

# gibt den Port 7000 frei
EXPOSE 7000

CMD ["/home/ic20b050/app/restapi"]

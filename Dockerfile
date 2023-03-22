FROM golang:1.20.2-alpine

# Workdir innerhalb des Containers festlegen
WORKDIR /go/src
# Kopiert das aktuelle Verzeichnis vom Host in das Image Verzeichnis
ADD . /go/src

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
RUN go mod download && go mod verify
RUN go build -v -o /go/src/app ./...

# gibt den Port 7000 frei
EXPOSE 7000

CMD ["/go/src/app"]

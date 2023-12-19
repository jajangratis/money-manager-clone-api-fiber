FROM golang:1.21.5

EXPOSE 3000/udp
EXPOSE 3000/tcp

WORKDIR /usr/src/app
COPY . .
RUN go mod download && go mod verify
RUN go build -v -o -ldflags="-s -w" -o money-manager-clone-api-fiber


CMD ["./money-manager-clone-api-fiber"]
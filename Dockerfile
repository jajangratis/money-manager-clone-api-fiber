FROM golang:1.21.4

WORKDIR /app
COPY . .

RUN go build -ldflags="-s -w" -o money-manager-clone-api-fiber

EXPOSE 3000
CMD ["./money-manager-clone-api-fiber"]
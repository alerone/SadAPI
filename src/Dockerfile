#Etapa 1: Construcción

FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . . 

RUN go build -o SadAPI.exe .

FROM alpine:latest 

WORKDIR /app

COPY --from=builder /app/SadAPI.exe /app

EXPOSE 8080

CMD ["./SadAPI.exe"]
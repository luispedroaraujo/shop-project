FROM golang:1.21-alpine AS build

# Install required packages for CGO (gcc and libc-dev)
RUN apk add --no-cache \
    build-base \
    gcc \
    libc-dev \
    libgcc

WORKDIR /app

COPY . .

RUN go build -o shop-api .

CMD ["./shop-api"]

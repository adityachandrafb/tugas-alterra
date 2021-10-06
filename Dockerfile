FROM golang:1.17-alpine AS build

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o app

FROM alpine:latest
COPY --from=build /app/app /app

EXPOSE 8080
CMD ["/app"]



# FROM golang:1.17-alpine

# WORKDIR /app

# COPY . .
# RUN go mod tidy
# RUN go build -o app


# EXPOSE 8080
# CMD ["./app"]

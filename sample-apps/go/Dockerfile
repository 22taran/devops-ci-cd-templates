FROM golang:1.21-alpine AS build
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /sample-app

FROM alpine:3.18
WORKDIR /app
COPY --from=build /sample-app .
EXPOSE 8080
USER nobody
ENTRYPOINT ["./sample-app"]

FROM rust:1.73-alpine AS build
WORKDIR /app
RUN apk add --no-cache musl-dev
COPY Cargo.toml Cargo.lock* ./
COPY src/ ./src/
RUN cargo build --release

FROM alpine:3.18
WORKDIR /app
COPY --from=build /app/target/release/sample-rust-app .
EXPOSE 8080
USER nobody
ENTRYPOINT ["./sample-rust-app"]

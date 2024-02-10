# Build the application
FROM golang:1.22-bookworm AS build
WORKDIR /app
COPY /app .
RUN go mod download && CGO_ENABLED=0 GOOS=linux go build -o /resume-server

# Copy the compiled binary into a clean image
FROM alpine:3.19 AS release
WORKDIR /
COPY --from=build /resume-server /resume-server
EXPOSE 8080
ENTRYPOINT ["/resume-server", "--port", "8080"]

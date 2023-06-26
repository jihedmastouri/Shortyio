# First stage: build the Go binary
FROM golang:1.19-alpine AS builder
# Specify DIR as an ARG
ARG DIR=.
# Set the working directory
WORKDIR /app
# Copy the source code into the container
COPY . .
# Copy the vendor directory into the container
COPY ${DIR}/vendor/. vendor
# Build the binary
RUN go build -o server ${DIR}/main.go

# Second stage: create the final image
FROM alpine
# Copy the binary from the first stage into the final image
COPY --from=builder /app/server /server
# Set the command to run the binary
CMD ["/server"]

# First stage: build the Go binary
FROM golang:1.19-alpine AS builder
# Specify DIR as an ARG
ARG DIR=.
# Set the working directory
WORKDIR /app
# Copy the source code into the container
COPY . .
# Change the current directory
WORKDIR ${DIR}
# Download dependencies
RUN go mod download
# Build the binary
RUN go build -o /app/server

# Second stage: create the final image
FROM alpine
# Copy the binary from the first stage into the final image
COPY --from=builder /app/server /server
# Set the command to run the binary
CMD ["/server"]

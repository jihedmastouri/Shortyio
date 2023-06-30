# First stage: build the Go binary
FROM golang:1.19-alpine AS builder
# Specify DIR as an ARG
ARG DIR=.
# Copy the source code
COPY . /app/
# Set the working directory
WORKDIR /app/${DIR}/
RUN go mod tidy -e
# Build the binary
RUN go build -o /app/server

# Second stage: create the final image
FROM alpine
# Copy the binary from the first stage into the final image
COPY --from=builder /app/server /server
# Set the command to run the binary
CMD ["/server"]

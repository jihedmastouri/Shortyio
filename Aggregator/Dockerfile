# First stage: build the Go binary
FROM golang:1.19-alpine AS builder
# Set the working directory
WORKDIR /app
# Copy the go.mod and go.sum files
COPY ./go.mod ./go.sum ./
# Download dependencies
RUN go mod download
# Copy the source files
COPY . .
# Build the binary
RUN go build -o server ./*.go

# Second stage: create the final image
FROM alpine
# Copy the binary from the first stage into the final image
COPY --from=builder /app/server /server
# Copy the temp.sql from the first stage into the final image
COPY --from=builder /app/temp.sql /temp.sql
# Set the command to run the binary
CMD ["/server"]

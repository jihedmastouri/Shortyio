# First stage: build the Go binary
FROM golang:1.19-alpine AS builder
# Copy the source code
COPY . /app/
# Set the working directory
WORKDIR /app/Aggregator
# Build the binary
RUN go build -o /app/server

# Second stage: create the final image
FROM alpine
# Copy the binary from the first stage into the final image
COPY --from=builder /app/server /server
# Copy the temp.sql from the first stage into the final image
COPY --from=builder /app/Aggregator/temp.sql /temp.sql
# Set the command to run the binary
CMD ["/server"]

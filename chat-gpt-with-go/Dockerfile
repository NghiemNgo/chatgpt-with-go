FROM golang:1.19.1

ENV TZ="Asia/Ho_Chi_Minh"
RUN date

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

WORKDIR /app/report-kt

# Build the Go app
RUN go build -o main .

# Expose port 2210 to the outside world
EXPOSE 2210

# Run the executable
CMD ["./main"]
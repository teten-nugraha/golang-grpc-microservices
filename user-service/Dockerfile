# Start dari image Go
FROM golang:1.23.2

# Set workspace
WORKDIR /app

# Copy semua file
COPY . .

# Download dependensi
RUN go mod tidy

# Build aplikasi
RUN go build -o service .

# Ekspos port sesuai dengan service
EXPOSE 50051

# Jalankan aplikasi
CMD ["./service"]

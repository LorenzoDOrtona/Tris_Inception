# ==========================================
# STAGE 1: The Builder (Il "Cantiere")
# ==========================================
# 1. Start from the official Go image
FROM golang:1.22-alpine AS builder

# 2. Set the working directory inside the container
WORKDIR /app

# 3. Copy go.mod and go.sum to download dependencies first
COPY go.mod go.sum ./
RUN go mod download

# 4. Copy the rest of your source code
COPY . .

# 5. Build the application!
# We tell Go to compile 'cmd/server/ServerMain.go' and output a file named 'game-server'
RUN CGO_ENABLED=0 GOOS=linux go build -o game-server ./cmd/server/ServerMain.go


# ==========================================
# STAGE 2: The Production Image (La "Casa Finita")
# ==========================================
# 6. Start from a completely empty, ultra-lightweight Linux image
FROM alpine:latest

# 7. Add root certificates (mandatory if your game makes external HTTPS calls)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# 8. Copy ONLY the compiled 'game-server' file from STAGE 1
COPY --from=builder /app/game-server .

# 9. THIS IS YOURS: Expose the port your Go server uses (e.g., 8080)
EXPOSE 8080

# 10. THIS IS YOURS: The command to run the compiled binary
CMD ["./game-server"]
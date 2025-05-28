# --------- Stage 1: Build frontend ---------
FROM node:20-alpine AS frontend-builder
WORKDIR /app
COPY readit-fe/package*.json ./
RUN npm install
COPY readit-fe .
RUN npm run build

# --------- Stage 2: Build backend with CGO support ---------
FROM golang:1.21 AS backend-builder
WORKDIR /app
COPY readit-be .
COPY --from=frontend-builder /app/dist ./ui

# Enable CGO and build
RUN CGO_ENABLED=1 go build -o server .

# --------- Final image with required C libraries ---------
FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
WORKDIR /root/
COPY --from=backend-builder /app/server .
COPY --from=backend-builder /app/ui ./ui
EXPOSE 8080
CMD ["./server"]
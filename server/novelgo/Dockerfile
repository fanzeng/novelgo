FROM golang:1.22-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN ls -l ./cmd/novelgo-server/
RUN go build -o main ./cmd/novelgo-server/main.go
EXPOSE 443
ENV PORT=443
ENV HOST=0.0.0.0
CMD ["sh", "-c", "export PORT=443 && ./main"]

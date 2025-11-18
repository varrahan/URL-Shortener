# URL Shortener

A lightweight, high-performance URL shortening service built with Go, Gin, and Redis.

## Features

- **Fast URL Shortening**: Generate short URLs using SHA-256 hashing and Base58 encoding
- **Redis Caching**: Efficient URL storage and retrieval with 1-hour cache duration
- **User-Specific Links**: Generate unique short URLs per user for the same long URL
- **RESTful API**: Clean and simple HTTP endpoints
- **Docker Support**: Easy deployment with Docker Compose

## Tech Stack

- **Go** - Backend language
- **Gin** - HTTP web framework
- **Redis** - In-memory data store
- **Docker** - Containerization

## Project Structure

```
.
├── cmd/main/              # Application entry point
├── internal/api/          # Internal API logic
│   ├── handler/          # HTTP request handlers
│   ├── shortener/        # URL shortening algorithm
│   ├── store/            # Redis storage layer
│   └── utils/            # Utility functions
├── pkg/api/              # Public API package
├── docker-compose.yaml   # Docker Compose configuration
├── dockerfile            # Docker image definition
└── README.md
```

## Prerequisites

- Go 1.16 or higher
- Redis server
- Docker and Docker Compose (for containerized deployment)

## Installation

### Local Development

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd url-shortener
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   
   Create a `.env` file in the project root:
   ```env
   REDIS_ADDR=localhost:6379
   INTERNAL_PORT=9000
   INTERNAL_ADDR=http://localhost:9000/
   EXPOSED_PORT=9000
   ```

4. **Run Redis locally**
   ```bash
   redis-server
   ```

5. **Start the application**
   ```bash
   go run cmd/main/main.go
   ```

### Docker Deployment

1. **Set up environment variables**
   
   Create a `.env` file with your configuration:
   ```env
   REDIS_ADDR=redis://redis:6379
   INTERNAL_PORT=9000
   INTERNAL_ADDR=http://localhost:9000/
   EXPOSED_PORT=9000
   ```

2. **Build and run with Docker Compose**
   ```bash
   docker-compose up --build
   ```

The service will be available at `http://localhost:9000`

## API Endpoints

### 1. Health Check

```http
GET /
```

**Response:**
```json
{
  "message": "Welcome to the URL Shortener API"
}
```

### 2. Create Short URL

```http
POST /create-short-url
Content-Type: application/json
```

**Request Body:**
```json
{
  "long_url": "https://example.com/very/long/url",
  "user_id": "user123"
}
```

**Response:**
```json
{
  "message": "short url created successfully",
  "short_url": "http://localhost:9000/aBcD1234"
}
```

### 3. Redirect to Original URL

```http
GET /:shortUrl
```

Redirects to the original long URL (HTTP 302).

## Usage Example

### Using cURL

```bash
# Create a short URL
curl -X POST http://localhost:9000/create-short-url \
  -H "Content-Type: application/json" \
  -d '{
    "long_url": "https://github.com/varrahan/url-shortener",
    "user_id": "john_doe"
  }'

# Access the short URL (redirects to original)
curl -L http://localhost:9000/aBcD1234
```

### Using JavaScript

```javascript
// Create short URL
const response = await fetch('http://localhost:9000/create-short-url', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
  },
  body: JSON.stringify({
    long_url: 'https://example.com',
    user_id: 'user123'
  })
});

const data = await response.json();
console.log(data.short_url);
```

## How It Works

1. **URL Hashing**: The service combines the long URL with a user ID and generates a SHA-256 hash
2. **Base58 Encoding**: The hash is converted to a Base58 string for URL-safe encoding
3. **Truncation**: The first 8 characters are used as the short URL identifier
4. **Storage**: The mapping is stored in Redis with a 1-hour expiration
5. **Retrieval**: When accessed, the short URL is looked up in Redis and redirects to the original URL

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `REDIS_ADDR` | Redis server address | `localhost:6379` |
| `INTERNAL_PORT` | Port the service runs on | `9000` |
| `INTERNAL_ADDR` | Base URL for short links | `http://0.0.0.0:9000/` |
| `EXPOSED_PORT` | External port (Docker) | `9000` |

## Configuration

- **Cache Duration**: URLs expire after 1 hour (configurable in `store.go`)
- **Short URL Length**: 8 characters (configurable in `shortener.go`)

## Development

### Running Tests

```bash
go test ./...
```

### Building the Binary

```bash
go build -o url-shortener cmd/main/main.go
```

## Docker Configuration

The project includes:
- **Dockerfile**: Multi-stage build for optimized image size
- **docker-compose.yaml**: Orchestrates the application and Redis
- **.dockerignore**: Excludes unnecessary files from the image

## License

This project is open source and available under the MIT License.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Acknowledgments

- Built with [Gin Web Framework](https://github.com/gin-gonic/gin)
- Uses [go-redis](https://github.com/go-redis/redis) for Redis integration
- Base58 encoding via [base58-go](https://github.com/itchyny/base58-go)
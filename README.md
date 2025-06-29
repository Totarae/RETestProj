# Pack Optimizer API

A web service that calculates optimal packaging based on available pack sizes. It selects the smallest possible number of packs with minimal overhead.

---

## Features

- REST API for calculating optimal packs
- Hot-reloadable configuration without server restart
- Web UI (HTML + JS)
- Structured HTTP request logging (using zap)
- Unit tests included

---

## Project Structure

```
awesomeProject10/
├── internal/
│   ├── config/         # Loads and watches packs.json
│   ├── handlers/       # HTTP handlers
│   ├── middleware/     # Logging middleware
│   ├── router/         # Chi router setup
│   └── service/        # Business logic
├── web/                # Static files for the UI
│   ├── index.html
│   └── static/
├── packs.json          # Configuration file with available pack sizes
├── main.go             # Entry point
└── go.mod / go.sum
```

---

## API Usage

### POST `/calculate`

**Request:**

```json
{
  "items": 12001
}
```

**Response:**

```json
{
  "packs": {
    "250": 1,
    "2000": 1,
    "5000": 2
  },
  "total": 12250
}
```

via Postman

## Configuration

Pack sizes are defined in `packs.json`:

```json
[250, 500, 1000, 2000, 5000]
```

Modifying this file will automatically update the config at runtime using `fsnotify`.


## Run Locally

```bash
go run main.go
```

Server runs at:  
`http://localhost:8080`

## Dependencies

- chi 
- zap 
- fsnotify 
- testify 
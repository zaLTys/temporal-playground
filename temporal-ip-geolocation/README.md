# Temporal IP Geolocation

A simple Temporal workflow application for IP geolocation.

## Prerequisites

- Go 1.25+
- Temporal CLI installed

## Getting Started

### 1. Start Temporal Server

```bash
temporal server start-dev
```

### 2. Start Worker

```bash
go run worker/main.go
```

### 3. Run Workflow

```bash
go run client/main.go <username>
```

Example:
```bash
go run client/main.go SomeUserName
```

The workflow will automatically fetch your IP address and location information.

## URLs

- **Temporal UI**: http://localhost:8088
- **Metrics**: http://localhost:9090

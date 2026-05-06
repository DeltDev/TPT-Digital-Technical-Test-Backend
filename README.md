# How to run backend

## Prerequisites
1. Docker Engine (Linux) or Docker for Desktop (Windows and Mac)
2. Go v1.26.2 or later

## Steps
1. Build PostgreSQL image
```bash
docker build -t tpt-postgres -f docker/Dockerfile .
```
2. Run PostgreSQL image
```bash
docker run -d --name tpt-db -p 5432:5432 tpt-postgres
```

3. Run Backend Server

```bash
go run cmd/server/main.go
```

## Other Commands
- Access PostgreSQL inside docker
```bash
docker exec -it tpt-db psql -U tpttechnicaltest -d tpttechnicaltest

When prompted password, the password is tpttechnicaltest
```

- Run unit test for all endpoints
```bash
go test -v ./internal/handler/...
```
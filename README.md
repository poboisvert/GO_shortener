# GO_shortener


## Requirements

- Docker
- Go
- Vs Code


## Running the App

```bash
git clone https://github.com/poboisvert/GO_shortener.git
```

```bash
cd GO_shortener
```

### Database SQL

```bash
cd server && docker-compose up
```

- Make sure the SQL server is running on PORT 3000

### API


```bash
go mod tidy
```


```bash
go run main.go
```

## How it works

- model
    - goly.go
    - model.go
- server
    - crud.go
    - server.go
- utils
    - random.go

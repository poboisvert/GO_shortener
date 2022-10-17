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
        - contain the structure of the query to link go and postgres SQL under Gorm
    - model.go
        - store the schema to communicate with the DB
- server
    - crud.go
        - contain all functions to link to an unique route
    - server.go
        - contain all routes to perform an event
- utils
    - random.go
        - generate a random value


## Postman

url: localhost:5000/url

Method: POST

Body: RAW with JSON format

```bash
{
    "url_text": "",
    "redirect": "https://mega.io/",
    "random": true
}
```

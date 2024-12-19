# HTTP Calculator Service

[Русский](README_ru.md)

## Overview

School project made for Yandex Lyceum. Simple calculator; evaluates math expression and gives an http response.

## Launching

Specify port in `cmd/main.go` (by default, port is `8080`)

Run in terminal:

```
go run cmd/main.go
```

## Using

Service accepts http requests on `/api/v1/calculate` endpoint in following format:

```
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression": "2+2"
}'
```

Expression can consist of numbers and arithmetic operations (+, -, *, /, unary minus and brackets).

### Examples

| Request Body               | Path              | Method       | Status Code | Message                      |
| -------------------------- | ----------------- | ------------ | ----------- | ---------------------------- |
| `{"expression":"2+2"}`     | /unsupported/path | POST         | 404         | 404 page not found           |
| `{"expression":"2+2"}`     | /api/v1/calculate | Any but POST | 405         | Method not allowed; use POST |
| `{"some":"thing"}`         | /api/v1/calculate | POST         | 400         | Bad request body             |
| `{"expression":"0/0"}`     | /api/v1/calculate | POST         | 422         | Expression is not valid      |
| `{"expression":"2.5+2.5"}` | /api/v1/calculate | POST         | 200         | 5                            |

### Notes

- Spaces are ignored, e.g. "2 + 2 3" is equivalent to "2+23"
- Omitting operation signs near brackets is not permitted, e.g. "2(1+5)" will return an error, not "12"
- Unary minus is permitted only after opening bracket or as the first character, e.g. "2*-2" will return an error, not "-4"; on the other hand, "-3+2*(-1)" will successfully return "-5"
- A period, not a comma, is used as a separator for decimal fractions, e.g. "2.5" and not "2,5"
# Calculator HTTP Service

Made for Yandex Lyceum

## Launching

Specify port in `cmd/main.go` (by default, port is `8080`)

Run in terminal:

```
go run cmd/main.go
```

## Using

Service accepts http requests in following format:

```
{
    "expression": "your expression here"
}
```

Expression consist of integers and arithmetic operations (+, -, *, /, unary minus and brackets). Support for float values as input is WIP.

If expression is valid, response code is `200` and response is following:

```
{
    "result": "calculation result"
}
```

Otherwise, response code is `422` and response is following:

```
{
    "error": "Expression is not valid"
}
```

## Examples:

### 1. Valid expression

Request:

```
{
    "expression": " (26/(-(81))-46)*62*30-(3)*(85)-0/17"
}
```

Response code `200`, response:

```
{
    "result": "-86412.037037"
}
```

### 2. Invalid expression

Request:

```
{
    "expression: "(-4+4*(7+8))-4)"
}
```

Response code `422`, response:

```
{
    "error": "Expression is not valid"
}
```

### 3. Invalid expression

Request:

```
{
    "expression: "Hello Go"
}
```

Response code `422`, response:

```
{
    "error": "Expression is not valid"
}
```

### 4. Invalid request

Request:

```
{
    "randomKey": "randomValue"
}
```

Response code `422`, response:

```
{
    "error": "Expression is not valid"
}
```

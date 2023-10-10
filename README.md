# ftgo-tastify

## Install

```bash
go get github.com/osvaldosilitonga/ftgo-tastify
```

## After Install

```bash
go mod tidy
```

## How to use

### Retrive Order Detail

Take order details from database

```go
cli.OrderDetails(ulrStr)
```

### Create Table

Will create "Employee, Items, Orders, OrderItems, Payments" table to database

```go
cli.CreateTable(urlStr)
```

### Isert Data

Will insert sample data to "Employees, MenuItems. Orders, OrderItems, Payments" table database

```go
cli.InsertData(urlStr)
```

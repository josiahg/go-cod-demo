# go-cod-demo

A simple example showing how to connect a golang app to Cloudera Operational Database

This app demonstrates how to create a connection, create a table, insert a row, and read that row.

Ddepends on the excellent [Calcite Avatica Go](https://github.com/apache/calcite-avatica-go) driver.

# Usage

Clone this repository

```
$ git clone https://github.com/josiahg/go-cod-demo
$ cd go-cod-demo
```

Install the `calcite-avatica-go` module, which is required

```
$ go get github.com/apache/calcite-avatica-go
```

Open `main.go` and update the dsn to reflect your COD URL, username, and password. Instructions are in the file.

Run it!

```
go run main.go
```

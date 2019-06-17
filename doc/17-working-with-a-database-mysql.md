# Working with a database, connecting with mysql

Package [sql](https://golang.org/pkg/database/sql/) provides a generic interface around SQL (or SQL-like) databases.
The sql package must be used in conjunction with a database driver. See https://golang.org/s/sqldrivers for a list of drivers.

*In order to connect to mysql we’ll be using https://github.com/go-sql-driver/mysql as our MySQL driver. Go-SQL-Driver is a lightweight and fast MySQL driver that supports connections over TCP/IPv4, TCP/IPv6, Unix domain sockets or custom protocols and features automatic handling of broken connections.*

**1- Open opens a database specified by its database driver name and a driver-specific data source name**
```go
db, err := sql.Open(
    "mysql",
    "go_user:1234@tcp(127.0.0.1:3396)/go_db"
)
```

**2-Query executes a query that returns rows, typically a SELECT**
```go
results, err := db.Query("SELECT id, a_varchar FROM test")
result, err := db.Query("INSERT INTO test (a_varchar) VALUES ( 'TEST' )")
```

**3-Next prepares the next result row for reading with the Scan method**
```go
results, err := db.Query("SELECT id, a_varchar FROM test")
if err != nil {
    panic(err.Error()) // proper error handling instead of panic in your app
}

for results.Next() {
    var row Row
    // for each row, scan the result into our row struct
    err = results.Scan(&row.ID, &row.Name)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    // and then print out the row's Name attribute
    log.Printf("%d - %s", row.ID, row.Name)
}
```
**4-Scan copies the columns in the current row into the values pointed at by dest** 

### Full example:
- [Create a table, insert and select the data](../src/18-working-with-database-mysql/connect-mysql.go)
- [Create a docker container with a mysql database](../src/18-working-with-database-mysql/docker-compose.yml)

### Links:
- **https://tutorialedge.net/golang/golang-mysql-tutorial/**
- https://www.alexedwards.net/blog/organising-database-access
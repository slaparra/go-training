# Working with files

Package to operate the file: [os](https://golang.org/pkg/os/)

```go
csvFile, err := os.Open("file-to-read.csv")

if err != nil {
		log.Fatal(err)
}
```

### Parsing a csv file

*Package [csv](https://golang.org/pkg/encoding/csv) reads and writes comma-separated values (CSV) files*
*NewReader returns a new Reader that reads from r.*
```go
import "encoding/csv"

reader := csv.NewReader(csvFile)

```

[Read](https://golang.org/pkg/encoding/csv/#Reader.Read) reads one record (a slice of fields) from reader.
```go
for {
    values, err := reader.Read()
    
    if err == io.EOF {
        break
    }
    
    if err != nil {
        log.Fatal(err)
    }
    ...
```
## Links
- **[Code example](../src/17-files/read.go)**
- https://gobyexample.com/reading-files
- https://gobyexample.com/writing-files
- https://forum.golangbridge.org/t/simple-way-to-write-string-and-or-varible-to-file/12186/9
- https://golang.org/pkg/os/  
- https://golang.org/pkg/encoding/csv
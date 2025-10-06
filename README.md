# size

A Go package for parsing and formatting human-readable file sizes.

## Features
- Parse strings like `"1.5GB"`, `"1024B"`, `"2 MB"` into bytes
- Format byte sizes into human-readable strings (e.g., `1.50 KB`, `2.00 MB`)
- Convert between bytes, kilobytes, megabytes, gigabytes, terabytes, and petabytes

## Usage

### Import
```go
import "github.com/harryrose/size"
```

### Parse a size string
```go
bytes, err := size.ParseSize("1.5GB")
if err != nil {
    // handle error
}
fmt.Println(bytes) // 1610612736
```

### Format a size
```go
s := size.Size(1536)
fmt.Println(s.String()) // "1.50 KB"
```

### Convert between units
```go
s := size.Size(1048576)
fmt.Println(s.Bytes())      // 1048576
fmt.Println(s.Kilobytes()) // 1024
fmt.Println(s.Megabytes()) // 1
```

## Testing
Run all tests:
```sh
go test -v
```

## License
MIT


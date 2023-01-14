HyperDB Go Client
=================

Author: **[Afaan Bilal](https://afaan.dev)**

## Introduction
**HyperDB Go** is a Go client package for the [HyperDB server](https://github.com/AfaanBilal/hyperdb).

## Installation
````
go get github.com/AfaanBilal/hyperdb-go
````

## Example usage
````go
import (
    "fmt"
    "github.com/AfaanBilal/hyperdb-go"
)

hyper := hyperdb.Create("http://localhost:8765", "", "")

// Ping the server
rb := hyper.ping()
fmt.Println(rb) // true

// Get the version number
rs := hyper.version()
fmt.Println(rs) // "[HyperDB v0.1.0 (https://afaan.dev)]"

// Set a value
rs = hyper.set("test", "value")
fmt.Println(rs) // "value"

// Check if a key is present
rb = hyper.has("test")
fmt.Println(rb) // true

// Get a value
rs = hyper.get("test")
fmt.Println(rs) // "value"

// Get all stored data
rs = hyper.all()
fmt.Println(rs) // {test: "value"}

// Remove a key
rb = hyper.delete("test")
fmt.Println(rb) // true

// Delete all stored data
rb = hyper.clear()
fmt.Println(rb) // true

// Check if the store is empty
rb = hyper.empty()
fmt.Println(rb) // true

// Persist the store to disk
rb = hyper.save()
fmt.Println(rb) // true

// Reload the store from disk
rb = hyper.reload()
fmt.Println(rb) // true

// Delete all store data from memory and disk
rb = hyper.reset()
fmt.Println(rb) // true
````

## Test
`$ go test .`

````
=== RUN   TestHyperDB
true
[HyperDB v0.2.0 (https://afaan.dev)]
value
true
value
{"test":"value"}
true
true
true
true
true
true
--- PASS: TestHyperDB (0.01s)
PASS
ok      github.com/AfaanBilal/hyperdb-go        0.201s
````

## Contributing
All contributions are welcome. Please create an issue first for any feature request
or bug. Then fork the repository, create a branch and make any changes to fix the bug
or add the feature and create a pull request. That's it!
Thanks!

## License
**HyperDB Go** is released under the MIT License.
Check out the full license [here](LICENSE).

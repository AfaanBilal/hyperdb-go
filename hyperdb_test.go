package hyperdb

/**
 * HyperDB Go Client
 *
 * @author Afaan Bilal
 * @link   https://afaan.dev
 */

import (
	"fmt"
	"testing"
)

func TestHyperDB(t *testing.T) {
	hyperdb := Create("http://localhost:8765", "", "")

	// Ping the server
	rb := hyperdb.Ping()
	fmt.Println(rb) // true

	// Get the version number
	rs := hyperdb.Version()
	fmt.Println(rs) // "[HyperDB v0.1.0 (https://afaan.dev)]"

	// Set a value
	rs = hyperdb.Set("test", "value")
	fmt.Println(rs) // "value"

	// Check if a key is present
	rb = hyperdb.Has("test")
	fmt.Println(rb) // true

	// Get a value
	rs = hyperdb.Get("test")
	fmt.Println(rs) // "value"

	// Get all stored data
	rs = hyperdb.All()
	fmt.Println(rs) // {test: "value"}

	// Remove a key
	rb = hyperdb.Delete("test")
	fmt.Println(rb) // true

	// Delete all stored data
	rb = hyperdb.Clear()
	fmt.Println(rb) // true

	// Check if the store is empty
	rb = hyperdb.Empty()
	fmt.Println(rb) // true

	// Persist the store to disk
	rb = hyperdb.Save()
	fmt.Println(rb) // true

	// Reload the store from disk
	rb = hyperdb.Reload()
	fmt.Println(rb) // true

	// Delete all store data from memory and disk
	rb = hyperdb.Reset()
	fmt.Println(rb) // true
}

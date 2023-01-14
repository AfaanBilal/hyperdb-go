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
	rb := hyperdb.ping()
	fmt.Println(rb) // true

	// Get the version number
	rs := hyperdb.version()
	fmt.Println(rs) // "[HyperDB v0.1.0 (https://afaan.dev)]"

	// Set a value
	rs = hyperdb.set("test", "value")
	fmt.Println(rs) // "value"

	// Check if a key is present
	rb = hyperdb.has("test")
	fmt.Println(rb) // true

	// Get a value
	rs = hyperdb.get("test")
	fmt.Println(rs) // "value"

	// Get all stored data
	rs = hyperdb.all()
	fmt.Println(rs) // {test: "value"}

	// Remove a key
	rb = hyperdb.delete("test")
	fmt.Println(rb) // true

	// Delete all stored data
	rb = hyperdb.clear()
	fmt.Println(rb) // true

	// Check if the store is empty
	rb = hyperdb.empty()
	fmt.Println(rb) // true

	// Persist the store to disk
	rb = hyperdb.save()
	fmt.Println(rb) // true

	// Reload the store from disk
	rb = hyperdb.reload()
	fmt.Println(rb) // true

	// Delete all store data from memory and disk
	rb = hyperdb.reset()
	fmt.Println(rb) // true
}

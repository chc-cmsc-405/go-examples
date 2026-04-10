// json-decode.go — Reading JSON into Go data with json.NewDecoder.
//
// The mirror of json-response.go. That file encodes Go data into JSON for
// responses. This file does the reverse: decoding JSON from request bodies
// into Go data. This is how POST endpoints read incoming data.
//
// Compare to Python: data = json.loads(body) — one line, returns a dict.
// Go is more explicit: you create a decoder, point it at a reader, and
// decode into a typed variable. The & (address-of) is required because
// the decoder fills in the variable rather than returning a new one.
package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	// === Decoding JSON into a map ===

	// Simulating what a POST request body looks like — a JSON string.
	// In your health monitor, r.Body is the reader instead of strings.NewReader.
	jsonBody := `{"name": "GitHub", "url": "https://github.com"}`
	reader := strings.NewReader(jsonBody)

	// Decode into a map[string]string. The & passes a pointer — the decoder
	// fills in 'service' directly rather than returning a new value.
	var service map[string]string
	err := json.NewDecoder(reader).Decode(&service)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}

	// Now 'service' is a regular Go map — access it with brackets.
	fmt.Println("Decoded service:")
	fmt.Println("  Name:", service["name"])
	fmt.Println("  URL:", service["url"])

	// === Encoding (the other direction) ===

	// json.NewEncoder writes Go data → JSON (see json-response.go).
	// json.MarshalIndent does the same but returns a byte slice instead of
	// writing to a writer — useful for debugging and logging.
	data := map[string]interface{}{
		"service": "Google",
		"status":  "UP",
		"checks":  42,
	}

	// MarshalIndent produces pretty-printed JSON with 2-space indentation.
	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error encoding:", err)
		return
	}
	fmt.Println("\nEncoded JSON:")
	fmt.Println(string(jsonBytes))
}

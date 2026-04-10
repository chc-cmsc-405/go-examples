// iteration.go — Looping over maps and building new data from existing data.
//
// Go uses the range keyword for all iteration — maps, slices, strings.
// On maps, range gives you (key, value) pairs. Important: map iteration
// order is random. Run this file twice and the output may change order.
// This is by design — Go prevents you from depending on insertion order.
package main

import "fmt"

func main() {
	services := map[string]string{
		"Google":  "https://google.com",
		"GitHub":  "https://github.com",
		"Netflix": "https://netflix.com",
	}

	// === Iterate over key-value pairs ===

	// range on a map gives you (key, value) each iteration.
	// Compare to Python: for name, url in services.items()
	fmt.Println("All services:")
	for name, url := range services {
		fmt.Printf("  %s → %s\n", name, url)
	}

	// === Keys only ===

	// Omit the second variable to get just keys.
	// Compare to Python: for name in services
	fmt.Println("\nService names:")
	for name := range services {
		fmt.Println(" ", name)
	}

	// === Counting pattern ===

	// A common pattern: iterate over data and count occurrences.
	// map[string]int starts at 0 for any key, so count++ works
	// on the first access — no need to check if the key exists first.
	urls := []string{
		"https://google.com/search",
		"https://google.com/maps",
		"https://github.com/go",
		"https://google.com/mail",
	}

	domainCount := make(map[string]int)
	for _, url := range urls {
		// Simple domain extraction (just for demo)
		if len(url) > 8 {
			domainCount[url]++
		}
	}

	fmt.Println("\nURL counts:")
	for url, count := range domainCount {
		fmt.Printf("  %s: %d\n", url, count)
	}
}

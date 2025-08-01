package main

import (
	"fmt"
	"github.com/riclib/icon"
)

func main() {
	// Create a simple example using the icon package
	homeIcon := icon.IconHouse
	userIcon := icon.IconUser
	settingsIcon := icon.IconSettings

	// Print some information about the icons
	fmt.Printf("House icon constant: %s\n", string(homeIcon))
	fmt.Printf("User icon constant: %s\n", string(userIcon))
	fmt.Printf("Settings icon constant: %s\n", string(settingsIcon))
	fmt.Printf("Total icons available: %d\n", icon.IconCount())
	
	// Test search functionality
	searcher := icon.NewIconSearcher()
	results := searcher.Search("user")
	fmt.Printf("Found %d icons matching 'user'\n", len(results))
	
	// Show first 3 results
	for i, result := range results {
		if i >= 3 {
			break
		}
		fmt.Printf("  - %s (relevance: %d)\n", string(result.IconName), result.Relevance)
	}
}

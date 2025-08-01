# riclib/icon

Type-safe Lucide icons for Go with the templ templating engine.

This standalone library provides all 1600+ Lucide icons as Go constants with full type safety. It includes powerful search and categorization functionality, making it easy to find and use the perfect icon for your application.

The icons are generated from the [Lucide Icons](https://lucide.dev) project and are designed to work seamlessly with the [templ](https://templ.guide) templating engine for Go.

## Features

- 🎯 **Type-safe** - All 1600+ icon names as constants
- 🔍 **Searchable** - Built-in search functionality with relevance scoring
- 📁 **Categorized** - Icons organized by categories
- 🎨 **Customizable** - Full control via templ.Attributes
- 🚀 **Tree-shakable** - Only used icons are included in your build
- 🔄 **Up-to-date** - Generated from latest Lucide icons

## Installation

```bash
go get github.com/riclib/icon
```

## Usage

### Basic Usage

```go
import "github.com/riclib/icon"

// Using the Icon function with type-safe constants
component := icon.Icon(icon.IconHouse, templ.Attributes{"class": "w-6 h-6 text-blue-500"})
component := icon.Icon(icon.IconUser, templ.Attributes{"class": "w-6 h-6 text-green-500"})

// Direct component usage (no attributes)
component := icon.House()
component := icon.User()

// Direct component usage with attributes
component := icon.HouseWithAttrs(templ.Attributes{"class": "w-6 h-6"})
component := icon.UserWithAttrs(templ.Attributes{"class": "w-6 h-6"})
```

### Categories

Access icons by category:

```go
// Get all navigation icons
navIcons := icon.NavigationIcons()
for _, iconName := range navIcons {
    component := icon.Icon(iconName, templ.Attributes{})
}

// Get all available categories
categories := icon.AllCategories()

// Check an icon's category
category := icon.GetIconCategory(icon.IconHouse) // "buildings"
```

### Search Functionality

```go
// Create a search instance
search := icon.NewIconSearcher()

// Search for icons
results := search.Search("arrow")
for _, result := range results {
    // result.IconName - the icon constant
    // result.Relevance - relevance score (0-100)
    // result.MatchType - "exact", "tag", "category", or "partial"
    component := icon.Icon(result.IconName, templ.Attributes{})
}

// Search by tag
icons := search.SearchByTag("navigation")

// Search by category
icons := search.SearchByCategory("actions")

// Advanced search with options
results := search.SearchWithOptions("edit", icon.SearchOptions{
    MaxResults:   10,
    MinRelevance: 50,
    Categories:   []string{"actions"},
})
```

### Utility Functions

```go
// Check if an icon exists
if icon.IconExists("house") {
    // icon exists
}

// Get total icon count
count := icon.IconCount() // 1600+

// Get all available icons
allIcons := icon.AllIcons()

// Convert string to IconName
if iconName, ok := icon.IconByName("house"); ok {
    component := icon.Icon(iconName, templ.Attributes{})
}
```

## Examples

See the [examples directory](examples/) for complete working examples.

### Simple Go Usage

```go
package main

import (
    "fmt"
    "github.com/riclib/icon"
)

func main() {
    // Access icon constants
    homeIcon := icon.IconHouse
    userIcon := icon.IconUser
    
    fmt.Printf("House icon: %s\n", string(homeIcon))
    fmt.Printf("User icon: %s\n", string(userIcon))
    fmt.Printf("Total icons: %d\n", icon.IconCount())
    
    // Use search functionality
    searcher := icon.NewIconSearcher()
    results := searcher.Search("user")
    fmt.Printf("Found %d icons matching 'user'\n", len(results))
}
```

## Requirements

- Go 1.24+
- [templ](https://github.com/a-h/templ) v0.3.920+

## License

Icons are from [Lucide](https://lucide.dev) (ISC License).
This package is MIT licensed.

## Regenerating Icons

To update to the latest Lucide icons:

```bash
./update-icons.sh
```

This script will:
1. Clone the latest Lucide repository
2. Generate all icon components
3. Update the registry and search indexes
4. Create category groupings

## Contributing

Contributions are welcome! Please feel free to submit issues or pull requests.

For issues related to the icons themselves, please check the [Lucide Icons](https://github.com/lucide-icons/lucide) repository first.

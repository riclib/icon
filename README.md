# riclib/icon

Type-safe Lucide icons for Go with the templ templating engine.

This standalone library provides all 1600+ Lucide icons as Go constants with full type safety. It includes powerful search and categorization functionality, making it easy to find and use the perfect icon for your application.

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
@icon.Icon(icon.IconHome, templ.Attributes{"class": "w-6 h-6"})
@icon.Icon(icon.IconArrowRight, templ.Attributes{"class": "text-blue-500"})

// Direct component usage
@icon.Home(templ.Attributes{"class": "w-6 h-6"})
@icon.ArrowRight(templ.Attributes{"class": "text-blue-500"})
```

### Styling Icons

Icons use `currentColor` for stroke, making them easy to style:

```go
// Size with CSS classes
@icon.Icon(icon.IconHeart, templ.Attributes{"class": "w-4 h-4"})

// Inline styles
@icon.Icon(icon.IconStar, templ.Attributes{
    "style": "width: 24px; height: 24px; color: gold;",
})
```

### Categories

Access icons by category:

```go
// Get all navigation icons
navIcons := icon.NavigationIcons()
for _, iconName := range navIcons {
    @icon.Icon(iconName, attrs)
}

// Available categories:
// Navigation, Actions, Media, Communication, Files, UI, Data, 
// Devices, Social, Weather, Transportation, Business, and more...

// Check an icon's category
category := icon.GetIconCategory(icon.IconHome) // "navigation"
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
    @icon.Icon(result.IconName, attrs)
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

## Requirements

- Go 1.24+
- [templ](https://github.com/a-h/templ) v0.3.920+

## License

Icons are from [Lucide](https://lucide.dev) (ISC License).
This package is MIT licensed.

## Contributing

To regenerate the icons from the latest Lucide release:

```bash
go run cmd/generate-icons/main.go
```

This will:
1. Clone the latest Lucide repository
2. Generate all icon components
3. Update the registry and search indexes
4. Create category groupings

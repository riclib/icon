# Lucide Icons for templ

Type-safe Lucide icons for use with [templ](https://templ.guide) templating engine.

This directory contains the generated icon files. For full documentation, see the main README at the root of this repository.

## Features

- 🎯 **Type-safe** - All 1600+ icon names as constants
- 🔍 **Searchable** - Built-in search functionality with relevance scoring
- 📁 **Categorized** - Icons organized by categories
- 🎨 **Customizable** - Full control via templ.Attributes
- 🚀 **Tree-shakable** - Only used icons are included in your build
- 🔄 **Up-to-date** - Generated from latest Lucide icons

## Usage

### Basic Usage

```go
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

### Utility Functions

```go
// Check if an icon exists
if icon.IconExists("home") {
    // icon exists
}

// Get total icon count
count := icon.IconCount() // 1600+

// Get all available icons
allIcons := icon.AllIcons()

// Convert string to IconName
if iconName, ok := icon.IconByName("home"); ok {
    @icon.Icon(iconName, attrs)
}
```

## Examples

### Dynamic Icon Selection

```go
templ StatusIcon(status string) {
    switch status {
    case "success":
        @icon.Icon(icon.IconCheckCircle, templ.Attributes{"class": "text-green-500"})
    case "error":
        @icon.Icon(icon.IconXCircle, templ.Attributes{"class": "text-red-500"})
    case "warning":
        @icon.Icon(icon.IconAlertTriangle, templ.Attributes{"class": "text-yellow-500"})
    default:
        @icon.Icon(icon.IconInfo, templ.Attributes{"class": "text-blue-500"})
    }
}
```

### Icon Button

```go
templ IconButton(iconName icon.IconName, label string) {
    <button class="flex items-center gap-2 px-4 py-2 bg-blue-500 text-white rounded">
        @icon.Icon(iconName, templ.Attributes{"class": "w-4 h-4"})
        <span>{ label }</span>
    </button>
}

// Usage
@IconButton(icon.IconDownload, "Download")
@IconButton(icon.IconShare2, "Share")
```

### Search UI Component

```go
templ IconSearch(query string) {
    <div class="grid grid-cols-6 gap-2">
        { search := icon.NewIconSearcher() }
        { results := search.Search(query) }
        for _, result := range results[:12] { // Show first 12 results
            <div class="text-center p-2 hover:bg-gray-100 rounded">
                @icon.Icon(result.IconName, templ.Attributes{"class": "w-8 h-8 mx-auto"})
                <p class="text-xs mt-1">{ string(result.IconName) }</p>
            </div>
        }
    </div>
}
```

## Icon Attributes

All icons accept templ.Attributes for customization:

- **class** - CSS classes
- **style** - Inline styles
- **id** - Element ID
- **data-*** - Data attributes
- **aria-*** - Accessibility attributes
- Any other valid SVG attributes

## Performance

- Icons are generated at build time - no runtime overhead
- Tree-shaking ensures only used icons are included
- Search indexes are built once and reused
- Minimal dependencies (only templ required)

## Regenerating Icons

To update to the latest Lucide icons:

```bash
go run cmd/generate-icons/main.go
```

This will:
1. Clone the latest Lucide repository
2. Generate all icon components
3. Update the registry and search indexes
4. Create category groupings

## License

Icons are from [Lucide](https://lucide.dev) (ISC License).
This package is MIT licensed.

---
name: icon-search
description: Search and discover Lucide icons from the riclib/icon Go library. Use when the user needs to find icons by name, tag, or category for use in Go/templ code.
---

# Icon Search

Search 1600+ Lucide icons from the `github.com/riclib/icon` library.

## Commands

The binary is at `~/.claude/skills/icon-search/bin/icon-search`.

### Search for icons by keyword

```bash
~/.claude/skills/icon-search/bin/icon-search search "home"
~/.claude/skills/icon-search/bin/icon-search search "arrow right" --limit 5
```

### Get full info for a specific icon

```bash
~/.claude/skills/icon-search/bin/icon-search info house
~/.claude/skills/icon-search/bin/icon-search info arrow-right
```

Returns tags, categories, and ready-to-use Go/templ code snippets.

### List all categories

```bash
~/.claude/skills/icon-search/bin/icon-search categories
```

### List/filter icons

```bash
~/.claude/skills/icon-search/bin/icon-search list --category navigation
~/.claude/skills/icon-search/bin/icon-search list --tag arrow --limit 10
~/.claude/skills/icon-search/bin/icon-search list --category navigation --tag arrow
```

## Using icons in Go/templ code

### Import

```go
import "github.com/riclib/icon"
```

### In templ templates

```templ
// Direct call (adds "icon" CSS class automatically)
@icon.House()

// With extra CSS classes
@icon.House("w-4", "h-4", "text-blue-500")

// With arbitrary HTML attributes
@icon.HouseWithAttrs(templ.Attributes{"class": "w-4 h-4", "aria-hidden": "true"})

// Dynamic icon by name constant
@icon.Icon(icon.IconHouse, templ.Attributes{"class": "w-4 h-4"})
```

### In Go code (raw SVG string)

```go
svg := icon.HouseSVG()
```

## Naming convention

Icon names are kebab-case (e.g., `arrow-big-right`). The Go constant and function names use PascalCase:

- Kebab name: `arrow-big-right`
- Constant: `icon.IconArrowBigRight`
- Function: `icon.ArrowBigRight()`
- With attrs: `icon.ArrowBigRightWithAttrs(attrs)`
- Raw SVG: `icon.ArrowBigRightSVG()`

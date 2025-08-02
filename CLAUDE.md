# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is the riclib/icon library - a type-safe Lucide icon library for Go applications providing 1600+ SVG icons. It uses the templ templating engine for generating Go components.

## Key Architecture

The library follows a code generation pattern:
1. Icons are sourced from the official Lucide repository
2. The `cmd/generate-icons` tool processes SVG files and generates Go code
3. Generated files are committed to the repository for easy consumption

### Core Components

- **Icon Registry**: `registry.templ` defines the `IconName` type and all icon constants (e.g., `IconHouse`)
- **Icon Components**: `icons.templ` contains templ components for each icon with multiple access patterns:
  - Direct functions: `icon.House()`
  - With attributes: `icon.HouseWithAttrs(attrs)`
  - Generic function: `icon.Icon(iconName, attrs)`
  - Raw SVG: `icon.HouseSVG()`
- **Search Engine**: `search.go` implements full-text search with relevance scoring and tag/category filtering
- **Categories**: `categories.go` maps icons to 40+ categories for organization

## Development Commands

```bash
# Regenerate all icons from latest Lucide
./update-icons.sh

# Generate templ files to Go code
templ generate

# Run tests
go test ./...

# Generate icons manually with custom options
go run cmd/generate-icons/main.go -out . -package icon

# Build examples
cd examples/basic && go build
```

## Testing Approach

Tests focus on:
- Icon attribute merging functionality
- CSS class handling (ensuring "icon" class is always present)
- HTML output validation
- Search functionality correctness

Run tests with `go test ./...` - they use table-driven patterns typical for Go.

## Working with Generated Code

Most icon code is auto-generated. When modifying the library:
1. **Never edit generated files directly** - they will be overwritten
2. Generated files are clearly marked with "Code generated" comments
3. To add new functionality, modify the generator in `internal/lucidegen/` or the templates
4. After modifying generation logic, run `./update-icons.sh` to regenerate

## Adding New Features

When adding features:
1. Check if it requires modifying the code generator
2. Ensure backward compatibility - this is a public library
3. Add appropriate tests
4. Update README.md with usage examples if adding user-facing features
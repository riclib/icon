#!/bin/bash

# Script to update the Lucide icons in the library
# This assumes you have the open-props-css project set up with the generate-icons command

echo "Updating Lucide icons..."

# Generate new icons
go run cmd/generate-icons/main.go -out . -package icon

# Go back to the riclib-icon project

# Generate templ files
templ generate

echo "Icon update complete!"
echo "Don't forget to commit the changes:"
echo "  git add ."
echo "  git commit -m \"Update Lucide icons to latest version\""

#!/bin/bash

# Script to update the Lucide icons in the library
# This assumes you have the open-props-css project set up with the generate-icons command

echo "Updating Lucide icons..."

# Go to the open-props-css project directory (adjust path as needed)
cd ../open-props-css

# Generate new icons
go run cmd/generate-icons/main.go -out ../riclib-icon -package icon

# Go back to the riclib-icon project
cd ../riclib-icon

# Generate templ files
templ generate

echo "Icon update complete!"
echo "Don't forget to commit the changes:"
echo "  git add ."
echo "  git commit -m \"Update Lucide icons to latest version\""

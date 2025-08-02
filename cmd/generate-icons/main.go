package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/riclib/icon/internal/lucidegen"
)

func main() {
	var config lucidegen.Config

	// Define command-line flags
	flag.StringVar(&config.OutputDir, "out", "./icon", "Output directory for generated files")
	flag.StringVar(&config.PackageName, "package", "icon", "Go package name")
	flag.StringVar(&config.Prefix, "prefix", "", "Prefix for constant names")
	flag.BoolVar(&config.DryRun, "dry-run", false, "Preview what would be generated without creating files")
	flag.BoolVar(&config.Verbose, "verbose", true, "Enable verbose logging")
	flag.BoolVar(&config.IncludeSearch, "search", true, "Include search functionality")

	// Custom usage function
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Lucide Icon Generator for templ\n\n")
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  # Generate all icons with default settings\n")
		fmt.Fprintf(os.Stderr, "  %s\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  # Preview what would be generated\n")
		fmt.Fprintf(os.Stderr, "  %s -dry-run\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  # Generate to custom directory\n")
		fmt.Fprintf(os.Stderr, "  %s -out ./myicons -package myicons\n", os.Args[0])
	}

	flag.Parse()

	// Make output directory absolute
	absOut, err := filepath.Abs(config.OutputDir)
	if err != nil {
		log.Fatalf("Failed to resolve output directory: %v", err)
	}
	config.OutputDir = absOut

	// Run generation
	result, err := lucidegen.Generate(config)
	if err != nil {
		log.Fatalf("Generation failed: %v", err)
	}

	// Print results
	if !config.DryRun {
		fmt.Printf("\n✓ Successfully generated %d icons in %v\n", result.IconsGenerated, result.Duration)
		fmt.Printf("\nFiles created:\n")
		for _, file := range result.FilesCreated {
			relPath, _ := filepath.Rel(".", file)
			fmt.Printf("  - %s\n", relPath)
		}
		fmt.Printf("\nCategories: %v\n", result.Categories)
		fmt.Printf("\nNext steps:\n")
		fmt.Printf("  1. Run 'templ generate' in the %s directory\n", config.OutputDir)
		fmt.Printf("  2. Import the package: import \"%s\"\n", config.PackageName)
		fmt.Printf("  3. Use icons: @%s.Icon(%s.IconHome, templ.Attributes{})\n", config.PackageName, config.PackageName)
	}
}

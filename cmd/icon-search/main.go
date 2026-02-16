package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/riclib/icon"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	subcmd := os.Args[1]
	switch subcmd {
	case "search":
		cmdSearch(os.Args[2:])
	case "info":
		cmdInfo(os.Args[2:])
	case "categories":
		cmdCategories()
	case "list":
		cmdList(os.Args[2:])
	case "help", "-h", "--help":
		printUsage()
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", subcmd)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Fprintf(os.Stderr, `icon-search — query the riclib/icon Lucide icon library

Usage:
  icon-search search <query> [--limit N]
  icon-search info <icon-name>
  icon-search categories
  icon-search list [--category X] [--tag Y] [--limit N]
`)
}

// --- search ---

type searchResultJSON struct {
	Name       string   `json:"name"`
	Constant   string   `json:"constant"`
	Function   string   `json:"function"`
	Tags       []string `json:"tags"`
	Categories []string `json:"categories"`
	Relevance  int      `json:"relevance"`
	MatchType  string   `json:"match_type"`
}

func cmdSearch(args []string) {
	flags, positional := splitArgs(args)
	fs := flag.NewFlagSet("search", flag.ExitOnError)
	limit := fs.Int("limit", 20, "max results")
	fs.Parse(flags)

	if len(positional) < 1 {
		fmt.Fprintln(os.Stderr, "usage: icon-search search <query> [--limit N]")
		os.Exit(1)
	}
	query := strings.Join(positional, " ")

	searcher := icon.NewIconSearcher()
	results := searcher.Search(query)

	if *limit > 0 && len(results) > *limit {
		results = results[:*limit]
	}

	out := make([]searchResultJSON, len(results))
	for i, r := range results {
		name := string(r.IconName)
		funcName := toFuncName(name)
		out[i] = searchResultJSON{
			Name:       name,
			Constant:   "Icon" + funcName,
			Function:   funcName,
			Tags:       searcher.GetTagsForIcon(r.IconName),
			Categories: searcher.GetCategoriesForIcon(r.IconName),
			Relevance:  r.Relevance,
			MatchType:  r.MatchType,
		}
	}

	writeJSON(out)
}

// --- info ---

type iconInfoJSON struct {
	Name       string            `json:"name"`
	Constant   string            `json:"constant"`
	Function   string            `json:"function"`
	Tags       []string          `json:"tags"`
	Categories []string          `json:"categories"`
	Usage      map[string]string `json:"usage"`
}

func cmdInfo(args []string) {
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "usage: icon-search info <icon-name>")
		os.Exit(1)
	}
	name := strings.ToLower(args[0])

	searcher := icon.NewIconSearcher()

	// Try exact match first
	iconName := icon.IconName(name)
	tags := searcher.GetTagsForIcon(iconName)
	categories := searcher.GetCategoriesForIcon(iconName)

	// If no tags/categories found, icon doesn't exist
	if len(tags) == 0 && len(categories) == 0 {
		// Try searching for it
		results := searcher.Search(name)
		if len(results) > 0 && strings.ToLower(string(results[0].IconName)) == name {
			iconName = results[0].IconName
			tags = searcher.GetTagsForIcon(iconName)
			categories = searcher.GetCategoriesForIcon(iconName)
		} else {
			fmt.Fprintf(os.Stderr, "icon not found: %s\n", name)
			os.Exit(1)
		}
	}

	funcName := toFuncName(string(iconName))
	constant := "Icon" + funcName

	info := iconInfoJSON{
		Name:       string(iconName),
		Constant:   constant,
		Function:   funcName,
		Tags:       tags,
		Categories: categories,
		Usage: map[string]string{
			"import":     `import "github.com/riclib/icon"`,
			"direct":     fmt.Sprintf("@icon.%s()", funcName),
			"with_attrs": fmt.Sprintf(`@icon.%sWithAttrs(templ.Attributes{"class": "w-4 h-4"})`, funcName),
			"generic":    fmt.Sprintf(`@icon.Icon(icon.%s, templ.Attributes{"class": "w-4 h-4"})`, constant),
			"raw_svg":    fmt.Sprintf("icon.%sSVG()", funcName),
		},
	}

	writeJSON(info)
}

// --- categories ---

func cmdCategories() {
	writeJSON(icon.AllCategories())
}

// --- list ---

type listItemJSON struct {
	Name     string `json:"name"`
	Constant string `json:"constant"`
	Function string `json:"function"`
}

func cmdList(args []string) {
	fs := flag.NewFlagSet("list", flag.ExitOnError)
	category := fs.String("category", "", "filter by category")
	tag := fs.String("tag", "", "filter by tag")
	limit := fs.Int("limit", 0, "max results (0 = all)")
	fs.Parse(args)

	searcher := icon.NewIconSearcher()
	var icons []icon.IconName

	switch {
	case *category != "" && *tag != "":
		// Intersect category and tag results
		catIcons := searcher.SearchByCategory(*category)
		tagSet := make(map[icon.IconName]bool)
		for _, n := range searcher.SearchByTag(*tag) {
			tagSet[n] = true
		}
		for _, n := range catIcons {
			if tagSet[n] {
				icons = append(icons, n)
			}
		}
	case *category != "":
		icons = searcher.SearchByCategory(*category)
	case *tag != "":
		icons = searcher.SearchByTag(*tag)
	default:
		// Return all icons via empty search
		results := searcher.Search("")
		icons = make([]icon.IconName, len(results))
		for i, r := range results {
			icons[i] = r.IconName
		}
	}

	if *limit > 0 && len(icons) > *limit {
		icons = icons[:*limit]
	}

	out := make([]listItemJSON, len(icons))
	for i, n := range icons {
		funcName := toFuncName(string(n))
		out[i] = listItemJSON{
			Name:     string(n),
			Constant: "Icon" + funcName,
			Function: funcName,
		}
	}

	writeJSON(out)
}

// --- helpers ---

func writeJSON(v any) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(v); err != nil {
		fmt.Fprintf(os.Stderr, "json encode error: %v\n", err)
		os.Exit(1)
	}
}

// splitArgs separates flag arguments (--key val) from positional arguments,
// allowing flags and positional args to be intermixed.
func splitArgs(args []string) (flags, positional []string) {
	for i := 0; i < len(args); i++ {
		if strings.HasPrefix(args[i], "-") {
			flags = append(flags, args[i])
			// consume the next arg as the flag value if it exists and isn't a flag
			if i+1 < len(args) && !strings.HasPrefix(args[i+1], "-") {
				i++
				flags = append(flags, args[i])
			}
		} else {
			positional = append(positional, args[i])
		}
	}
	return
}

// toFuncName converts a kebab-case icon name to PascalCase Go identifier.
// Duplicated from internal/lucidegen to avoid coupling to internal package.
func toFuncName(name string) string {
	parts := strings.Split(name, "-")
	var result strings.Builder
	for _, part := range parts {
		if len(part) > 0 {
			result.WriteString(strings.ToUpper(part[:1]))
			if len(part) > 1 {
				result.WriteString(part[1:])
			}
		}
	}
	funcName := result.String()
	if len(funcName) > 0 && !(funcName[0] >= 'a' && funcName[0] <= 'z') && !(funcName[0] >= 'A' && funcName[0] <= 'Z') {
		funcName = "Icon" + funcName
	}
	return funcName
}

package icon

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestIconAttributeMerging(t *testing.T) {
	tests := []struct {
		name           string
		attrs          templ.Attributes
		expectedClass  string
		containsIcon   bool
	}{
		{
			name:          "No attributes",
			attrs:         nil,
			expectedClass: `class="icon"`,
			containsIcon:  true,
		},
		{
			name:          "Empty attributes",
			attrs:         templ.Attributes{},
			expectedClass: `class="icon"`,
			containsIcon:  true,
		},
		{
			name: "Single class attribute",
			attrs: templ.Attributes{
				"class": "size-4",
			},
			expectedClass: `class="icon size-4"`,
			containsIcon:  true,
		},
		{
			name: "Multiple classes",
			attrs: templ.Attributes{
				"class": "size-4 text-primary hover:text-secondary",
			},
			expectedClass: `class="icon size-4 text-primary hover:text-secondary"`,
			containsIcon:  true,
		},
		{
			name: "Mixed attributes with class",
			attrs: templ.Attributes{
				"class": "size-6",
				"id":    "my-icon",
				"data-test": "icon-test",
			},
			expectedClass: `class="icon size-6"`,
			containsIcon:  true,
		},
		{
			name: "Style attribute only",
			attrs: templ.Attributes{
				"style": "color: red;",
			},
			expectedClass: `class="icon"`,
			containsIcon:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Render the icon with attributes
			buf := &bytes.Buffer{}
			component := AArrowDownWithAttrs(tt.attrs)
			err := component.Render(context.Background(), buf)
			if err != nil {
				t.Fatalf("Failed to render component: %v", err)
			}

			html := buf.String()

			// Check if the expected class is present
			if !strings.Contains(html, tt.expectedClass) {
				t.Errorf("Expected HTML to contain %q, but got:\n%s", tt.expectedClass, html)
			}

			// Check if "icon" class is always present
			if tt.containsIcon && !strings.Contains(html, "icon") {
				t.Errorf("Expected HTML to always contain 'icon' class, but got:\n%s", html)
			}

			// Ensure the span tag exists
			if !strings.Contains(html, "<span") {
				t.Errorf("Expected HTML to contain a span tag, but got:\n%s", html)
			}

			// Check for other attributes when provided
			if id, hasID := tt.attrs["id"]; hasID {
				expectedID := `id="` + id.(string) + `"`
				if !strings.Contains(html, expectedID) {
					t.Errorf("Expected HTML to contain %q, but got:\n%s", expectedID, html)
				}
			}

			if dataTest, hasDataTest := tt.attrs["data-test"]; hasDataTest {
				expectedDataTest := `data-test="` + dataTest.(string) + `"`
				if !strings.Contains(html, expectedDataTest) {
					t.Errorf("Expected HTML to contain %q, but got:\n%s", expectedDataTest, html)
				}
			}
		})
	}
}

// TestMergeClasses tests the mergeClasses function directly
func TestMergeClasses(t *testing.T) {
	tests := []struct {
		name          string
		attrs         templ.Attributes
		expectedClass string
	}{
		{
			name:          "Nil attributes",
			attrs:         nil,
			expectedClass: "icon",
		},
		{
			name:          "Empty attributes",
			attrs:         templ.Attributes{},
			expectedClass: "icon",
		},
		{
			name: "Single class",
			attrs: templ.Attributes{
				"class": "test-class",
			},
			expectedClass: "icon test-class",
		},
		{
			name: "Multiple classes",
			attrs: templ.Attributes{
				"class": "class1 class2 class3",
			},
			expectedClass: "icon class1 class2 class3",
		},
		{
			name: "Empty class string",
			attrs: templ.Attributes{
				"class": "",
			},
			expectedClass: "icon",
		},
		{
			name: "Non-string class value",
			attrs: templ.Attributes{
				"class": 123, // Not a string
			},
			expectedClass: "icon",
		},
		{
			name: "Other attributes preserved",
			attrs: templ.Attributes{
				"class": "custom",
				"id":    "test-id",
				"style": "color: blue;",
			},
			expectedClass: "icon custom",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mergeClasses(tt.attrs)
			
			// Check class attribute
			if class, ok := result["class"].(string); !ok || class != tt.expectedClass {
				t.Errorf("Expected class %q, but got %q", tt.expectedClass, class)
			}

			// Check other attributes are preserved
			if tt.attrs != nil {
				for k, v := range tt.attrs {
					if k != "class" {
						if result[k] != v {
							t.Errorf("Expected attribute %q to be %v, but got %v", k, v, result[k])
						}
					}
				}
			}
		})
	}
}
package icon

import "github.com/a-h/templ"

// IconFunc is the signature shared by all icon rendering functions (e.g. icon.House).
// Consumers can use this type to accept any icon as a parameter.
type IconFunc func(class ...string) templ.Component

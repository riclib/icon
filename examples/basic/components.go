package main

import (
	"fmt"
	"github.com/a-h/templ"
	"github.com/riclib/icon"
)

func main() {
	// Create templ components for icons using the IconWithAttrs function
	homeComponent := icon.IconWithAttrs(icon.IconHouse, templ.Attributes{"class": "w-6 h-6 text-blue-500"})
	userComponent := icon.IconWithAttrs(icon.IconUser, templ.Attributes{"class": "w-6 h-6 text-green-500"})
	settingsComponent := icon.IconWithAttrs(icon.IconSettings, templ.Attributes{"class": "w-6 h-6 text-purple-500"})
	
	// Print information about the components
	fmt.Printf("Home component: %T\n", homeComponent)
	fmt.Printf("User component: %T\n", userComponent)
	fmt.Printf("Settings component: %T\n", settingsComponent)
	
	// You can also use direct component functions (no attributes)
	homeComponent2 := icon.House()
	userComponent2 := icon.User()
	settingsComponent2 := icon.Settings()
	
	fmt.Printf("Home component (no attrs): %T\n", homeComponent2)
	fmt.Printf("User component (no attrs): %T\n", userComponent2)
	fmt.Printf("Settings component (no attrs): %T\n", settingsComponent2)
	
	fmt.Println("Icon components are ready to use in templ templates!")
}

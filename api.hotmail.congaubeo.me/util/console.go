package util

import (
	"fmt"
	"sort"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/logrusorgru/aurora"
	"github.com/thoas/go-funk"
)

// ConsolePrintServiceSuccess ...
func ConsolePrintServiceSuccess(serviceName string, value string) {
	p1 := "⇨ " + serviceName
	// p2 := aurora.BgGreen(aurora.Bold(aurora.White(value)))
	p2 := aurora.Bold(aurora.Green(value))
	fmt.Println(p1, "- READY -", p2)
}

// ConsolePrintServerRoutes ...
func ConsolePrintServerRoutes(routes []*echo.Route) {
	// Remove all middleware routes
	routes = funk.Filter(routes, func(item *echo.Route) bool {
		return strings.Contains(item.Name, "controllers") // Only allow routes within project
	}).([]*echo.Route)

	// Sort by path alphabets
	sort.Slice(routes, func(i, j int) bool {
		return routes[i].Path < routes[j].Path
	})

	fmt.Println("")
	fmt.Println(aurora.Bold(aurora.Green("- ROUTES:")))
	for _, r := range routes {
		method := fmt.Sprintf("%-10v", "["+r.Method+"]")
		path := fmt.Sprintf("%-50v", r.Path)
		controller := fmt.Sprintf("--> %s", r.Name)
		fmt.Println(method, path, controller)
	}
}

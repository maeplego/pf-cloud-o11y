package route

import "strings"

// Pattern returns a low-cardinality route template for metrics and traces.
func Pattern(method, path string) string {
	if path == "" || path == "/" {
		return method + " /"
	}
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) >= 2 && parts[0] == "work" {
		parts[1] = ":id"
	}
	return method + " /" + strings.Join(parts, "/")
}

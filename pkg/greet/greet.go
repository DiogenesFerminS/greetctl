package greet

import "strings"

func Hello(name string) string {
	cleanName := normalizeName(name)
	return "Hello " + cleanName + " :)"
}

func normalizeName(name string) string {
	normalizedName := strings.TrimSpace(name)

	if name == "" {
		return "mundo"
	}

	return normalizedName
}

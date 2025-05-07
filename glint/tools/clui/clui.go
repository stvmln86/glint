// Package clui implements command-line user interface functions.
package clui

import (
	"fmt"
	"os"
	"strings"
)

// Env returns an existing environment variable.
func Env(name string) (string, error) {
	data, ok := os.LookupEnv(name)
	data = strings.TrimSpace(data)

	switch {
	case !ok:
		return "", fmt.Errorf("cannot find envvar %q - does not exist", name)
	case data == "":
		return "", fmt.Errorf("cannot find envvar %q - is blank", name)
	default:
		return data, nil
	}
}

// Parse returns a parsed argument map from a parameter slice and argument slice.
func Parse(paras, elems []string) (map[string]string, error) {
	var pairs = make(map[string]string)

	for i, para := range paras {
		name, dflt, ok := strings.Cut(para, ":")

		switch {
		case i < len(elems):
			pairs[name] = elems[i]
		case ok:
			pairs[name] = dflt
		default:
			return nil, fmt.Errorf("cannot parse arguments - %q not provided", name)
		}
	}

	return pairs, nil
}

// Split returns a command name and argument slice from an argument slice.
func Split(elems []string) (string, []string) {
	switch len(elems) {
	case 0:
		return "", nil
	case 1:
		return strings.ToLower(elems[0]), nil
	default:
		return strings.ToLower(elems[0]), elems[1:]
	}
}

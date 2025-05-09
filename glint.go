///////////////////////////////////////////////////////////////////////////////////////
//           glint · a minimal note-taking engine in Go · by Stephen Malone          //
///////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

///////////////////////////////////////////////////////////////////////////////////////
//                          part one · constants and globals                         //
///////////////////////////////////////////////////////////////////////////////////////

// GlintDir is the default notes directory $GLINT_DIR.
var GlintDir string

// GlintExt is the default file extension $GLINT_EXT.
var GlintExt string

// Stdout is the default standard output stream.
var Stdout = os.Stdout

// Usage is the command-line usage help string.
const Usage = `
Glint is a minimal note-taking engine, written in Go 1.24 by Stephen Malone.

Variables:
• $GLINT_DIR :: The directory containing your note files.
• $GLINT_EXT :: The extension your note files use.

Commands:
• edit NOTE :: Open NOTE in $EDITOR or $VISUAL.
• find TEXT :: List all notes with bodies containing TEXT.
• list NAME :: List all notes with names starting with NAME. 
• show NOTE :: Print NOTE's body to STDOUT. 
• hide NOTE :: Hide NOTE by changing extension to ".hidden".

See github.com/stvmln86/glint for more information.
`

///////////////////////////////////////////////////////////////////////////////////////
//                       part one · path manipulation functions                      //
///////////////////////////////////////////////////////////////////////////////////////

// PathName returns a path's lowercase name string.
func PathName(path string) string {
	base := filepath.Base(path)
	name := strings.SplitN(base, ".", 2)[0]
	return strings.ToLower(name)
}

///////////////////////////////////////////////////////////////////////////////////////
//                          part two · file system functions                         //
///////////////////////////////////////////////////////////////////////////////////////

// FileGlob returns a sorted slice of all paths in a directory matching an extension.
func FileGlob(dire, extn string) []string {
	glob := filepath.Join(dire, "*."+extn)
	paths, _ := filepath.Glob(glob)
	slices.Sort(paths)
	return paths
}

///////////////////////////////////////////////////////////////////////////////////////
//                           part three · command functions                          //
///////////////////////////////////////////////////////////////////////////////////////

// Commands is a map of all defined command functions.
var Commands = map[string]func(io.Writer, string) error{
	"list": CommandList,
}

// CommandList lists all existing notes with names starting with arg1.
func CommandList(w io.Writer, arg1 string) error {
	for _, path := range FileGlob(GlintDir, GlintExt) {
		name := PathName(path)
		if strings.HasPrefix(name, arg1) {
			fmt.Fprintf(w, "%s\n", name)
		}
	}

	return nil
}

///////////////////////////////////////////////////////////////////////////////////////
//                             part four · main functions                            //
///////////////////////////////////////////////////////////////////////////////////////

// die prints a formatted error message and exits.
func die(w io.Writer, text string, elems ...any) {
	text = fmt.Sprintf("Error: %s.\n", text)
	fmt.Fprintf(w, text, elems...)
	os.Exit(1)
}

// main executes the main Glint program.
func main() {
	var name, arg1 string

	// Set global variables.
	GlintDir = os.Getenv("GLINT_DIR")
	GlintExt = os.Getenv("GLINT_EXT")

	// Collect and confirm global variables.
	switch {
	case GlintDir == "":
		die(Stdout, "$GLINT_DIR is blank or not set")

	case GlintExt == "":
		die(Stdout, "$GLINT_EXT is blank or not set")

	case strings.HasPrefix(GlintExt, "."):
		GlintExt = strings.TrimPrefix(GlintExt, ".")
	}

	// Collect and confirm command-line arguments.
	switch len(os.Args) {
	case 0, 1:
		fmt.Fprintf(Stdout, strings.TrimSpace(Usage)+"\n")
		os.Exit(0)

	case 2:
		name = strings.ToLower(os.Args[1])

	default:
		name = strings.ToLower(os.Args[1])
		arg1 = os.Args[2]
	}

	// Collect and confirm command function.
	cfun, ok := Commands[name]
	if !ok {
		die(Stdout, "command %s does not exist", name)
	}

	// Execute chosen command.
	if err := cfun(Stdout, arg1); err != nil {
		die(Stdout, err.Error())
	}
}

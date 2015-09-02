package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/matryer/version"
)

/*

  version
  by Mat Ryer

  Command line tool to help manage version numbers
  for projects and directories.

  Copyright (c) 2014 Stretchr, Inc.

  Usage:

		Getting versions:

	    version /path       	- Reads the current version, or outputs
	                          	v0.0.0 if none is set.

			version -short /path	-	Gets smallest possible version number.


		Changing versions:

	    version /path 1.0    	- Sets the new version number to 1.0

	    version /path +      	- Increases the build version number
	                         	  so v1.2.3 becomes v1.2.4

	    version /path ++     	- Increases the minor build number
	                         	  so v1.2.3 becomes v1.3.0

	    version /path +++    	- Increases the major build number
	                         	  so v1.2.3 becomes v2.0.0

*/

const (
	// ExitCodeBadArgs is the exit code when arguments are bad.
	ExitCodeBadArgs = 1
)

// flags
var (
	shortFormat      = flag.Bool("short", false, "Whether the shortest possible version number should be output.  E.g. instead of v1.0.0, you'd just get v1.")
	includeV         = flag.Bool("v", true, "Whether to include the v (for version) prefix in the output or not.")
	suppressLinefeed = flag.Bool("n", false, "Whether to suppress the linefeed at the end of the output or not.")
)

func main() {

	// parse the flags
	flag.Parse()

	var err error
	var dir string
	var option string

	// get non-flag arguments
	var args []string
	for _, arg := range os.Args {
		if !strings.HasPrefix(arg, "-") {
			args = append(args, arg)
		}
	}

	switch len(args) {
	case 3:
		// dir, option
		option = args[2]
		fallthrough
	case 2:
		// dir
		dir = args[1]
	default:
		// unknown args
		writeError("Expected 1 or 2 arguments.")
		os.Exit(ExitCodeBadArgs)
		return
	}

	// check dir
	dir, err = filepath.Abs(dir)

	if err != nil {
		writeError("Bad path")
		os.Exit(ExitCodeBadArgs)
		return
	}

	// make sure the directory exists
	_, err = os.Stat(dir)
	if os.IsNotExist(err) {
		writeError(fmt.Sprintf("Directory does not exist: %s", dir))
		os.Exit(ExitCodeBadArgs)
		return
	}

	// check option
	var newV *version.Version
	switch option {
	case "":
		newV, err = version.LoadVersion(dir)
	case "+":
		_, newV, err = version.IncreaseVersion(dir, 0, 0, 1)
	case "++":
		_, newV, err = version.IncreaseVersion(dir, 0, 1, 0)
	case "+++":
		_, newV, err = version.IncreaseVersion(dir, 1, 0, 0)
	default:

		// try and parse the version
		newV, err = version.Parse(option)

		if err != nil {
			writeError(fmt.Sprintf("Invalid version or unacceptable option: %s", option))
			os.Exit(ExitCodeBadArgs)
			return
		}

		// set the version
		err = version.SaveVersion(dir, newV)

	}

	if err != nil {
		writeError(fmt.Sprintf("Failed to update version: %s", err))
		os.Exit(ExitCodeBadArgs)
		return
	}

	// return the new version
	fmt.Print(versionString(newV))

	// and a line feed?
	if !*suppressLinefeed {
		fmt.Print("\n")
	}

}

func versionString(v *version.Version) string {

	if *shortFormat {
		if *includeV {
			return v.StringShort()
		}
		return v.StringShortNumber()
	}

	if *includeV {
		return v.String()
	}
	return v.StringNumber()

}

func writeError(message string) {
	fmt.Println(message)
	writeHelp()
}
func writeHelp() {
	fmt.Println("\n---")
	fmt.Println("version - by Mat Ryer")
	fmt.Println("Copyright (c)2014 Stretchr, Inc.  http://www.stretchr.com/")
	fmt.Println("\nUSAGE")
	fmt.Println("  version [flags] path [option]")
	fmt.Println("  path   - Path to set the version for")
	fmt.Println("  option - none  Read the version")
	fmt.Println("         - +     Increase the build number (1.0.0 -> 1.0.1)")
	fmt.Println("         - ++    Increase the minor number (1.0.0 -> 1.1.0)")
	fmt.Println("         - +++   Increase the major number (1.0.0 -> 2.0.0)")
	fmt.Println("\nOptional flags:")
	flag.PrintDefaults()
	fmt.Println("")
}

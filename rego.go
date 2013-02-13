/**
 * The purpose of this package is
 * to take in two parameters.  The
 * first is a regular expression.
 *
 * The second is a pattern which should
 * be used to rename all the files to.
 *
 * This is useful for instance to renaming
 * a large number of files that were downloaded
 * from the internet so that they're nice.
 */
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

var (
	simulate *bool
)

func main() {
	var (
		dirname   *string
		recursive *bool
	)

	// Delcare our command line flags and then parse them
	//source := *flag.String("source", "", "The regex used to determine which files will be renamed (required)")
	//rename_pattern := *flag.String("rename-pattern", "", "The pattern to rename files by.  This will use regex group matches such as $1, $2, etc. (required)")

	dirname = flag.String("dir", ".", "The directory to scan for files")
	recursive = flag.Bool("recursive", false, "If sub-directories should be scanned recusively")
	simulate = flag.Bool("simulate", false, "Simulate the running of rego, don't actually rename files")
	flag.Parse()

	// Retrieve command line arguments
	source := flag.Arg(0)
	rename_pattern := flag.Arg(1)

	// See if there is the minimum number of command line arguments provided.
	// If there is nothing then we print the help message
	if flag.NArg() < 2 {
		fmt.Print("Usage of rego, example: rego \"(.*)\\.tgz\" \"$1.tar.gz\"\n\n")
		fmt.Println("Arguments:")
		flag.PrintDefaults()
		return
	}

	if len(source) == 0 || len(rename_pattern) == 0 {
		fmt.Println("A value must be supplied for both command line arguments.")
		return
	}

	// Compile and ensure source regular expression is valid
	src_regex, src_err := regexp.Compile(source)
	if src_err != nil {
		fmt.Println("Could not compile source regular expression.")
		fmt.Println(src_err)
		return
	}

	// Rename the files in the given directory
	renameFilesInDir(*dirname, src_regex, rename_pattern, *recursive)
}

/**
 * Rename the files in a given directory which match a regular expression.
 */
func renameFilesInDir(dirname string, src_regex *regexp.Regexp, rename_pattern string, recursive bool) {
	// Retrieve a reader for the directory
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		fmt.Println("Could not read the provided directory.")
		fmt.Printf("Error: %s\n", err)
	}

	// Go through all the files in the directory
	for _, v := range files {
		filename := v.Name()

		// If the filename matches the given regular expression then rename it
		if !v.IsDir() && src_regex.MatchString(filename) {
			new_filename := src_regex.ReplaceAllString(filename, rename_pattern)

			if !(*simulate) {
				os.Rename(filename, new_filename)
			}

			fmt.Printf("Renamed file %s to %s\n", filename, new_filename)

			// If the entry is a directory and we're recursively renaming then
			// launch a go-routine for renaming sub-directories
		} else if v.IsDir() && recursive {
			//go renameFilesInDir(filename, src_regex, rename_pattern, recursive)
		}
	}
}

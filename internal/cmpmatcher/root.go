package cmpmatcher

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Run() error {
	path := "./"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	// suffix
	if strings.HasSuffix(path, ".go") {
		return wrapCmpMatcher(path)
	} else {
		//nolint:errcheck
		return filepath.WalkDir(path, func(path string, entry os.DirEntry, err error) error {
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failure accessing a path %q: %v\n", path, err)
				return err
			}
			if entry.IsDir() {
				return nil
			}
			if strings.HasSuffix(entry.Name(), ".go") {
				wrapCmpMatcher(path)
			}
			return nil
		})
	}
}

package dir

import (
	"log"
	"os"
	"path/filepath"

	"github.com/gravetii/diztl/conf"
)

// GetOutputPath : Returns the output file path for the file.
func GetOutputPath(filename string) string {
	return filepath.Join(conf.OutputDir(), filename)
}

// Ensure : Checks if the required directories are created, creating them if not.
func Ensure() {
	ensureDirs(conf.ShareDirs())
	ensureDir(conf.OutputDir())
}

func ensureDirs(dirs []string) {
	for _, dir := range dirs {
		ensureDir(dir)
	}
}

func ensureDir(dir string) {
	info, err := os.Stat(dir)
	if os.IsNotExist(err) {
		log.Printf("Creating directory: %s\n", dir)
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			log.Fatalf("Could not create directory - %s: %v\n", dir, err)
		}
	} else if !info.IsDir() {
		log.Fatalf("Seems like there's a resource already: %s\n", dir)
	} else if err != nil {
		log.Fatalf("Could not ensure that directory exists - %s: %v\n", dir, err)
	}
}

// Split : Splits the path into individual tokens and returns a slice of the tokens.
func Split(path string) []string {
	res := []string{}
	path = filepath.Clean(path)
	dir, name := filepath.Split(path)
	dir = filepath.Clean(dir)
	if dir == "/" {
		res = append(res, dir)
	} else if dir != "" {
		r := Split(dir)
		if len(r) > 0 {
			res = append(res, r...)
		}
	}

	res = append(res, name)
	return res
}

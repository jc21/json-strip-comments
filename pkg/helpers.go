package pkg

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/jc21/json-strip-comments/pkg/model"
)

// GetContents grabs content from either stdin or a file
func GetContents(cfg model.ArgConfig) (string, error) {
	// check if reading froms stdin
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		contents, err := ioutil.ReadAll(os.Stdin)
		return string(contents), err
	}

	// check if filename is given
	if cfg.Filename != "" {
		if _, err := os.Stat(cfg.Filename); errors.Is(err, os.ErrNotExist) {
			return "", fmt.Errorf("File does not exist: %s", cfg.Filename)
		}
		contents, err := ioutil.ReadFile(cfg.Filename)
		return string(contents), err
	}

	return "", nil
}

// WriteContents writes to a file or stdout
func WriteContents(cfg model.ArgConfig, contents string) error {
	res := RemoveComments(contents, cfg.Empty)

	// write to a file
	outputFilename := replaceOutputFileWildcards(cfg)
	if outputFilename != "" {
		return writeToFile(outputFilename, res)
	}

	// Write to stdout
	fmt.Printf("%s\n", res)
	return nil
}

func writeToFile(filename, contents string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	// nolint: errcheck, gosec
	defer f.Close()
	_, err2 := f.WriteString(contents)
	if err2 != nil {
		return err2
	}
	return nil
}

func replaceOutputFileWildcards(cfg model.ArgConfig) string {
	// Only do this when the filename is defined on the command line,
	// wildcards don't apply to pipes
	if cfg.Filename != "" && cfg.Output != "" {
		inFile := path.Clean(cfg.Filename)

		// get base name of incoming file.
		// ie: `/path/to/something.jsonc` => `something`
		baseName := fileNameWithoutExt(path.Base(inFile))

		// Get the base path for the incoming file.
		dirName, _ := filepath.Abs(inFile)
		dirName = path.Dir(dirName)

		filename := strings.ReplaceAll(path.Clean(cfg.Output), "[folder]", fmt.Sprintf("%s/", dirName))
		filename = strings.ReplaceAll(filename, "[file]", baseName)

		return filename
	}

	return cfg.Output
}

func fileNameWithoutExt(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}

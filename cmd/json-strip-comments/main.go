package main

import (
	"fmt"
	"os"

	"github.com/jc21/json-strip-comments/pkg"
	"github.com/jc21/json-strip-comments/pkg/config"
)

func main() {
	cfg := config.GetConfig()
	contents, err := pkg.GetContents(cfg)
	if err != nil {
		fmt.Printf("ERROR: %s", err.Error())
		os.Exit(1)
	}
	if contents == "" {
		fmt.Println("Usage:\n  cat file.json | json-strip-comments [-e]\n  json-strip-comments [-e] /path/to/file.jsonc")
		os.Exit(1)
	}

	if err := pkg.WriteContents(cfg, contents); err != nil {
		fmt.Printf("ERROR: %s", err.Error())
		os.Exit(1)
	}
}

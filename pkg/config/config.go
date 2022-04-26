package config

import (
	"github.com/JeremyLoy/config"
	"github.com/alexflint/go-arg"
	"github.com/jc21/json-strip-comments/pkg/model"
)

// Populated at build time using ldflags
var appArguments model.ArgConfig

// GetConfig returns the ArgConfig
func GetConfig() model.ArgConfig {
	config.FromEnv().To(&appArguments)
	arg.MustParse(&appArguments)

	return appArguments
}

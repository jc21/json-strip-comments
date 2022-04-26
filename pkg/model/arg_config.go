package model

// ArgConfig is the settings for passing arguments to the command
type ArgConfig struct {
	Filename string `arg:"positional"`
	Output   string `arg:"-o" help:"Output to given filename instead of stdout"`
	Empty    bool   `arg:"-e" help:"Remove empty lines after comment removals"`
}

// Description returns a simple description of the command
func (ArgConfig) Description() string {
	return "Removes c style comments from a json file"
}

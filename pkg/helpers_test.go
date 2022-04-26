package pkg

import (
	"testing"

	"github.com/jc21/json-strip-comments/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestReplaceOutputFileWildcards(t *testing.T) {
	tests := []struct {
		name string
		cfg  model.ArgConfig
		want string
	}{
		{
			name: "Nothing set",
			cfg: model.ArgConfig{
				Filename: "",
				Output:   "",
			},
			want: "",
		},
		{
			name: "No output file",
			cfg: model.ArgConfig{
				Filename: "something.jsonc",
				Output:   "",
			},
			want: "",
		},
		{
			name: "incoming file with ext",
			cfg: model.ArgConfig{
				Filename: "MyFullData.jsonc",
				Output:   "[file].json",
			},
			want: "MyFullData.json",
		},
		{
			name: "incoming file without ext",
			cfg: model.ArgConfig{
				Filename: "MyFullData",
				Output:   "[file].json",
			},
			want: "MyFullData.json",
		},
		{
			name: "invalid wildcard",
			cfg: model.ArgConfig{
				Filename: "MyFullData",
				Output:   "[filename].json",
			},
			want: "[filename].json",
		},
		{
			name: "folder and file",
			cfg: model.ArgConfig{
				Filename: "/path/to/MyFullData.jsonc",
				Output:   "[folder][file].json",
			},
			want: "/path/to/MyFullData.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := replaceOutputFileWildcards(tt.cfg)
			assert.Equal(t, tt.want, res)
		})
	}
}

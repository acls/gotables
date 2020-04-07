package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_parseConfig(t *testing.T) {
	tests := []struct {
		name    string
		txt     string
		want    []route
		wantErr bool
	}{
		{"empty", strings.Join([]string{}, "\n"), nil, false},
		{"only comments", strings.Join([]string{
			"#",
			" #",
		}, "\n"), nil, false},
		{"one line", strings.Join([]string{
			"0.0.0.0:1389 192.168.25.19:389",
		}, "\n"), []route{
			route{src: "0.0.0.0:1389", dst: "192.168.25.19:389"},
		}, false},
		{"multiple lines", strings.Join([]string{
			"0.0.0.0:1389 192.168.25.19:389",
			"0.0.0.0:1636 192.168.25.19:636",
		}, "\n"), []route{
			route{src: "0.0.0.0:1389", dst: "192.168.25.19:389"},
			route{src: "0.0.0.0:1636", dst: "192.168.25.19:636"},
		}, false},
		{"empty line", strings.Join([]string{
			"0.0.0.0:1389 192.168.25.19:389",
			"",
			"0.0.0.0:1636 192.168.25.19:636",
		}, "\n"), []route{
			route{src: "0.0.0.0:1389", dst: "192.168.25.19:389"},
			route{src: "0.0.0.0:1636", dst: "192.168.25.19:636"},
		}, false},

		{"invalid line", strings.Join([]string{
			"0.0.0.0:1636",
		}, "\n"), nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseConfig(strings.NewReader(tt.txt))
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.EqualValues(t, tt.want, got)
			}
		})
	}
}

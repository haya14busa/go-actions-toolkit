package core

import "testing"

func TestCommand(t *testing.T) {
	tests := []struct {
		cmd  string
		p    CommandProperties
		mes  string
		want string
	}{
		{
			cmd:  "set-env",
			p:    NewCommandProperties("name", "ENV_NAME"),
			mes:  "ENV_VALUE",
			want: "::set-env name=ENV_NAME::ENV_VALUE",
		},
		{
			cmd:  "error",
			p:    NewCommandProperties("line", "1", "col", "2"),
			mes:  "text",
			want: "::error line=1,col=2::text",
		},
		{
			cmd:  "add-path",
			p:    nil,
			mes:  "/path/to/dir",
			want: "::add-path::/path/to/dir",
		},
	}
	for _, tt := range tests {
		if got := NewCommand(tt.cmd, tt.p, tt.mes).String(); got != tt.want {
			t.Errorf("Command = %s, want %s", got, tt.want)
		}
	}
}

func TestCommandProperties_Add(t *testing.T) {
	cp := NewCommandProperties()
	cp.Add("line", "14")
	if got := cp["line"]; got != "14" {
		t.Errorf("cp['line'] = %s, want 14", got)
	}
}

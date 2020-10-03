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

		// test cases for escaping
		{
			cmd:  "some-command",
			p:    nil,
			mes:  "percent % percent % cr \r cr \r lf \n lf \n",
			want: "::some-command::percent %25 percent %25 cr %0D cr %0D lf %0A lf %0A",
		},
		{
			cmd:  "some-command",
			p:    nil,
			mes:  "%25 %25 %0D %0D %0A %0A",
			want: "::some-command::%2525 %2525 %250D %250D %250A %250A",
		},
		{
			cmd:  "some-command",
			p:    NewCommandProperties("name", "percent % percent % cr \r cr \r lf \n lf \n colon : colon : comma , comma ,"),
			mes:  "",
			want: "::some-command name=percent %25 percent %25 cr %0D cr %0D lf %0A lf %0A colon %3A colon %3A comma %2C comma %2C::",
		},
		{
			cmd:  "some-command",
			p:    NewCommandProperties("name", "%25 %25 %0D %0D %0A %0A %3A %3A %2C %2C"),
			mes:  "",
			want: "::some-command name=%2525 %2525 %250D %250D %250A %250A %253A %253A %252C %252C::",
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

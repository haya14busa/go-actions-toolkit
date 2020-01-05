package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/google/subcommands"
	"github.com/haya14busa/go-actions-toolkit/core"
)

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&exportVariable{}, "")
	subcommands.Register(&setSecret{}, "")
	subcommands.Register(&addPath{}, "")
	subcommands.Register(&getInput{}, "")
	subcommands.Register(&setOutput{}, "")
	subcommands.Register(&debug{}, "")
	subcommands.Register(&errorCmd{}, "")
	subcommands.Register(&warning{}, "")
	subcommands.Register(&info{}, "")
	subcommands.Register(&startGroup{}, "")
	subcommands.Register(&endGroup{}, "")
	subcommands.Register(&saveState{}, "")
	subcommands.Register(&getState{}, "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}

//-----------------------------------------------------------------------
// ExportVariable
//-----------------------------------------------------------------------

type exportVariable struct {
	name  string
	value string
}

func (*exportVariable) Name() string { return "export-variable" }
func (*exportVariable) Synopsis() string {
	return "Sets env variable for this action and future actions in the job."
}
func (*exportVariable) Usage() string {
	return `export-variable -name=<name> -value=<value>
`
}
func (sc *exportVariable) SetFlags(f *flag.FlagSet) {
	f.StringVar(&sc.name, "name", "", "env name")
	f.StringVar(&sc.value, "value", "", "env value")
}
func (sc *exportVariable) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	core.ExportVariable(sc.name, sc.value)
	return subcommands.ExitSuccess
}

//-----------------------------------------------------------------------
// SetSecret
//-----------------------------------------------------------------------

type setSecret struct{}

func (*setSecret) Name() string { return "set-secret" }
func (*setSecret) Synopsis() string {
	return "Registers a secret which will get masked from logs."
}
func (*setSecret) Usage() string {
	return `set-secret <value>
`
}
func (sc *setSecret) SetFlags(f *flag.FlagSet) {}
func (sc *setSecret) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	core.SetSecret(f.Arg(0))
	return subcommands.ExitSuccess
}

//-----------------------------------------------------------------------
// AddPath
//-----------------------------------------------------------------------

type addPath struct{}

func (*addPath) Name() string { return "add-path" }
func (*addPath) Synopsis() string {
	return "Prepends inputPath to the PATH (for this action and future actions)."
}
func (*addPath) Usage() string {
	return `add-path <value>
`
}
func (sc *addPath) SetFlags(f *flag.FlagSet) {}
func (sc *addPath) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	core.AddPath(f.Arg(0))
	return subcommands.ExitSuccess
}

//-----------------------------------------------------------------------
// GetInput
//-----------------------------------------------------------------------

type getInput struct{}

func (*getInput) Name() string { return "get-input" }
func (*getInput) Synopsis() string {
	return `Gets the value of an input.`
}
func (*getInput) Usage() string {
	return `get-input <name>
`
}
func (sc *getInput) SetFlags(f *flag.FlagSet) {}
func (sc *getInput) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	fmt.Println(core.GetInput(f.Arg(0)))
	return subcommands.ExitSuccess
}

//-----------------------------------------------------------------------
// SetOutput
//-----------------------------------------------------------------------

type setOutput struct {
	name  string
	value string
}

func (*setOutput) Name() string { return "set-output" }
func (*setOutput) Synopsis() string {
	return `Sets the value of an output.`
}
func (*setOutput) Usage() string {
	return `set-output -name=<name> -value=<value>
`
}
func (sc *setOutput) SetFlags(f *flag.FlagSet) {
	f.StringVar(&sc.name, "name", "", "name")
	f.StringVar(&sc.value, "value", "", "value")
}
func (sc *setOutput) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	core.SetOutput(sc.name, sc.value)
	return subcommands.ExitSuccess
}

//-----------------------------------------------------------------------
// Debug
//-----------------------------------------------------------------------

type debug struct {
	file string
	line int
	col  int
}

func (*debug) Name() string { return "debug" }
func (*debug) Synopsis() string {
	return `Writes debug message to user log.`
}
func (*debug) Usage() string {
	return `debug [-file=<file name> -line=<value> -col=<value>] <message>
`
}
func (sc *debug) SetFlags(f *flag.FlagSet) {
	f.StringVar(&sc.file, "file", "", "")
	f.IntVar(&sc.line, "line", 0, "")
	f.IntVar(&sc.col, "col", 0, "")
}
func (sc *debug) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	core.Debug(f.Arg(0), &core.LogOption{File: sc.file, Line: sc.line, Col: sc.col})
	return subcommands.ExitSuccess
}

//-----------------------------------------------------------------------
// Error
//-----------------------------------------------------------------------

type errorCmd struct {
	file string
	line int
	col  int
}

func (*errorCmd) Name() string { return "error" }
func (*errorCmd) Synopsis() string {
	return `Writes error message to user log.`
}
func (*errorCmd) Usage() string {
	return `error [-file=<file name> -line=<value> -col=<value>] <message>
`
}
func (sc *errorCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&sc.file, "file", "", "")
	f.IntVar(&sc.line, "line", 0, "")
	f.IntVar(&sc.col, "col", 0, "")
}
func (sc *errorCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	core.Error(f.Arg(0), &core.LogOption{File: sc.file, Line: sc.line, Col: sc.col})
	return subcommands.ExitSuccess
}

//-----------------------------------------------------------------------
// Warning
//-----------------------------------------------------------------------

type warning struct {
	file string
	line int
	col  int
}

func (*warning) Name() string { return "warning" }
func (*warning) Synopsis() string {
	return `Writes warning message to user log.`
}
func (*warning) Usage() string {
	return `warning [-file=<file name> -line=<value> -col=<value>] <message>
`
}
func (sc *warning) SetFlags(f *flag.FlagSet) {
	f.StringVar(&sc.file, "file", "", "")
	f.IntVar(&sc.line, "line", 0, "")
	f.IntVar(&sc.col, "col", 0, "")
}
func (sc *warning) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	core.Warning(f.Arg(0), &core.LogOption{File: sc.file, Line: sc.line, Col: sc.col})
	return subcommands.ExitSuccess
}

//-----------------------------------------------------------------------
// Info
//-----------------------------------------------------------------------

type info struct{}

func (*info) Name() string { return "info" }
func (*info) Synopsis() string {
	return "Writes info to log."
}
func (*info) Usage() string {
	return `info <value>
`
}
func (sc *info) SetFlags(f *flag.FlagSet) {}
func (sc *info) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	core.Info(f.Arg(0))
	return subcommands.ExitSuccess
}

//-----------------------------------------------------------------------
// StartGroup
//-----------------------------------------------------------------------

type startGroup struct{}

func (*startGroup) Name() string { return "start-group" }
func (*startGroup) Synopsis() string {
	return "Begin an output group."
}
func (*startGroup) Usage() string {
	return `start-group <name>
`
}
func (sc *startGroup) SetFlags(f *flag.FlagSet) {}
func (sc *startGroup) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	core.StartGroup(f.Arg(0))
	return subcommands.ExitSuccess
}

//-----------------------------------------------------------------------
// EndGroup
//-----------------------------------------------------------------------

type endGroup struct{}

func (*endGroup) Name() string { return "end-group" }
func (*endGroup) Synopsis() string {
	return "End an output group."
}
func (*endGroup) Usage() string {
	return `end-group
`
}
func (sc *endGroup) SetFlags(f *flag.FlagSet) {}
func (sc *endGroup) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	core.EndGroup()
	return subcommands.ExitSuccess
}

//-----------------------------------------------------------------------
// SaveState
//-----------------------------------------------------------------------

type saveState struct {
	name  string
	value string
}

func (*saveState) Name() string { return "save-state" }
func (*saveState) Synopsis() string {
	return `Saves state for current action, the state can only be retrieved by this action's post job execution.`
}
func (*saveState) Usage() string {
	return `save-state -name=<name> -value=<value>
`
}
func (sc *saveState) SetFlags(f *flag.FlagSet) {
	f.StringVar(&sc.name, "name", "", "name")
	f.StringVar(&sc.value, "value", "", "value")
}
func (sc *saveState) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	core.SaveState(sc.name, sc.value)
	return subcommands.ExitSuccess
}

//-----------------------------------------------------------------------
// GetState
//-----------------------------------------------------------------------

type getState struct{}

func (*getState) Name() string { return "get-state" }
func (*getState) Synopsis() string {
	return "Gets the value of an state set by this action's main execution."
}
func (*getState) Usage() string {
	return `get-state <name>
`
}
func (sc *getState) SetFlags(f *flag.FlagSet) {}
func (sc *getState) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	fmt.Println(core.GetState(f.Arg(0)))
	return subcommands.ExitSuccess
}

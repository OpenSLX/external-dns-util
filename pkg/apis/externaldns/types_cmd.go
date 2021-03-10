package externaldns

import "github.com/alecthomas/kingpin"

// CmdArgs provides additional arguments for external-dns-util
type CmdArgs struct {
	Command       string
	CommandName   string
	CommandType   string
	CommandTarget string
	CommandTTL    int
}

// GlobalCmdArgs contains the parsed arguments
var GlobalCmdArgs CmdArgs

func parseCmdArgs(args *[]string, app *kingpin.Application) {
	*args = append([]string{"--source", "fake"}, *args...)

	app.Command("show", "")
	addCmd := app.Command("add", "")
	addCmd.Arg("name", "").StringVar(&GlobalCmdArgs.CommandName)
	addCmd.Arg("type", "").StringVar(&GlobalCmdArgs.CommandType)
	addCmd.Arg("target", "").StringVar(&GlobalCmdArgs.CommandTarget)
	addCmd.Arg("ttl", "").IntVar(&GlobalCmdArgs.CommandTTL)
	delCmd := app.Command("del", "")
	delCmd.Arg("name", "").StringVar(&GlobalCmdArgs.CommandName)
	delCmd.Arg("type", "").StringVar(&GlobalCmdArgs.CommandType)

	GlobalCmdArgs.Command, _ = app.Parse(*args)
}

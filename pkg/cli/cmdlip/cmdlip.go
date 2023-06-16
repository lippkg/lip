package cmdlip

import (
	"flag"
	"fmt"
	"os"

	"github.com/lippkg/lip/pkg/cli/cmdlipautoremove"
	"github.com/lippkg/lip/pkg/cli/cmdlipcache"
	"github.com/lippkg/lip/pkg/cli/cmdlipexec"
	"github.com/lippkg/lip/pkg/cli/cmdlipinstall"
	"github.com/lippkg/lip/pkg/cli/cmdliplist"
	"github.com/lippkg/lip/pkg/contexts"
	"github.com/lippkg/lip/pkg/logging"
)

type FlagDict struct {
	helpFlag    bool
	versionFlag bool
	verboseFlag bool
	quietFlag   bool
}

const helpMessage = `
Usage:
  lip [options] [<command> [subcommand options]] ...

Commands:
  autoremove                  Uninstall tooths that are not depended by any other tooths.
  cache                       Inspect and manage Lip's cache.
  exec                        Execute a Lip tool.
  install                     Install a tooth.
  list                        List installed tooths.
  show                        Show information about installed tooths.
  tooth                       Maintain a tooth.
  uninstall                   Uninstall a tooth.

Options:
  -h, --help                  Show help.
  -V, --version               Show version and exit.
  -v, --verbose               Show verbose output.
  -q, --quiet                 Show only errors.
`

func Run(ctx contexts.Context, args []string) error {
	var err error

	flagSet := flag.NewFlagSet("lip", flag.ContinueOnError)

	// Rewrite the default messages.
	flagSet.Usage = func() {
		// Do nothing.
	}

	// Parse flags.
	var flagDict FlagDict
	flagSet.BoolVar(&flagDict.helpFlag, "help", false, "")
	flagSet.BoolVar(&flagDict.helpFlag, "h", false, "")
	flagSet.BoolVar(&flagDict.versionFlag, "version", false, "")
	flagSet.BoolVar(&flagDict.versionFlag, "V", false, "")
	flagSet.BoolVar(&flagDict.verboseFlag, "verbose", false, "")
	flagSet.BoolVar(&flagDict.verboseFlag, "v", false, "")
	flagSet.BoolVar(&flagDict.quietFlag, "quiet", false, "")
	flagSet.BoolVar(&flagDict.quietFlag, "q", false, "")
	err = flagSet.Parse(args[1:])
	if err != nil {
		return fmt.Errorf("cannot parse flags: %w", err)
	}

	// Set logging level.
	if flagDict.verboseFlag {
		logging.SetLoggingLevel(logging.DebugLevel)
	} else if flagDict.quietFlag {
		logging.SetLoggingLevel(logging.ErrorLevel)
	} else {
		logging.SetLoggingLevel(logging.InfoLevel)
	}

	// Help flag has the highest priority.
	if flagDict.helpFlag {
		logging.Info(helpMessage)
		return nil
	}

	// Version flag has the second highest priority.
	if flagDict.versionFlag {
		logging.Info("Lip %v from %v", ctx.LipVersion().String(), os.Args[0])
		return nil
	}

	// Verbose and quiet flags are mutually exclusive.
	if flagDict.verboseFlag && flagDict.quietFlag {
		return fmt.Errorf("verbose and quiet flags are mutually exclusive")
	}

	// If there is a subcommand, run it and exit.
	if flagSet.NArg() >= 1 {
		switch flagSet.Arg(0) {
		case "autoreremove":
			err = cmdlipautoremove.Run(ctx, flagSet.Args()[1:])
			if err != nil {
				return fmt.Errorf("failed to run the 'autoreremove' command: %w", err)
			}
			return nil

		case "cache":
			err = cmdlipcache.Run(ctx, flagSet.Args()[1:])
			if err != nil {
				return fmt.Errorf("failed to run the 'cache' command: %w", err)
			}
			return nil

		case "exec":
			err = cmdlipexec.Run(ctx, flagSet.Args()[1:])
			if err != nil {
				return fmt.Errorf("failed to run the 'exec' command: %w", err)
			}
			return nil

		case "install":
			err = cmdlipinstall.Run(ctx, flagSet.Args()[1:])
			if err != nil {
				return fmt.Errorf("failed to run the 'install' command: %w", err)
			}

		case "list":
			err = cmdliplist.Run(ctx, flagSet.Args()[1:])
			if err != nil {
				return fmt.Errorf("failed to run the 'list' command: %w", err)
			}

		default:
			return fmt.Errorf("unknown command: lip %v", flagSet.Arg(0))
		}
	}

	return fmt.Errorf("no command specified. See 'lip --help' for more information")
}
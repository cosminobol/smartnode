package obol

import (
	"github.com/urfave/cli/v2"

	cliutils "github.com/rocket-pool/smartnode/v2/rocketpool-cli/utils"
	"github.com/rocket-pool/smartnode/v2/shared/utils"
)

// Register commands
func RegisterCommands(app *cli.App, name string, aliases []string) {
	app.Commands = append(app.Commands, &cli.Command{
		Name:    name,
		Aliases: aliases,
		Usage:   "Manage Obol Distributed Validator",
		Subcommands: []*cli.Command{
			{
				Name:    "health",
				Aliases: []string{"h"},
				Usage:   "Get Charon service health",
				Action: func(c *cli.Context) error {
					// Validate args
					utils.ValidateArgCount(c, 0)

					// Run
					return getCharonHealth(c)
				},
			},
			{
				Name:    "manage-charon-service",
				Aliases: []string{"s"},
				Usage:   "Start, stop or restart the Charon service",
				Action: func(c *cli.Context) error {
					// Validate args
					utils.ValidateArgCount(c, 0)

					// Run
					return manageCharonService(c)
				},
			},
			{
				Name:    "create-enr",
				Aliases: []string{"n"},
				Usage:   "Create ENR for a Charon client",
				Action: func(c *cli.Context) error {
					// Validate args
					utils.ValidateArgCount(c, 0)

					// Run
					return createENR(c)
				},
			},
			{
				Name:    "charon-dkg",
				Aliases: []string{"d"},
				Usage:   "Run the Distributed Key Generation (DKG) ceremony",
				Action: func(c *cli.Context) error {
					// Validate args
					utils.ValidateArgCount(c, 0)

					// Run
					return runCharonDkg(c)
				},
			},
			{
				Name:    "get-validator-public-keys",
				Aliases: []string{"k"},
				Usage:   "Retrieves validator public keys",
				Action: func(c *cli.Context) error {
					// Validate args
					utils.ValidateArgCount(c, 0)

					// Run
					return getValidatorPublicKeys(c)
				},
			},
			{
				Name:    "dv-exit-sign",
				Aliases: []string{"e"},
				Usage:   "Sign a partial DV exit",
				Action: func(c *cli.Context) error {
					// Validate args
					utils.ValidateArgCount(c, 0)

					// Run
					return dvExitSign(c)
				},
			},
			{
				Name:    "dv-exit-broadcast",
				Aliases: []string{"b"},
				Usage:   "Publish a DV exit",
				Action: func(c *cli.Context) error {
					// Validate args
					utils.ValidateArgCount(c, 0)

					// Run
					return dvExitBroadcast(c)
				},
			}
		},
	})
}

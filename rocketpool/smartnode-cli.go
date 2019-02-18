package main

import (
    "fmt"
    "log"
    "os"

    "github.com/urfave/cli"

    "github.com/rocket-pool/smartnode-cli/rocketpool/api/deposit"
    "github.com/rocket-pool/smartnode-cli/rocketpool/api/fee"
    "github.com/rocket-pool/smartnode-cli/rocketpool/api/minipool"
    "github.com/rocket-pool/smartnode-cli/rocketpool/api/node"
    "github.com/rocket-pool/smartnode-cli/rocketpool/api/service"
)


// Run
func main() {

    // Add logo to application help template
    cli.AppHelpTemplate = fmt.Sprintf(`
______           _        _    ______           _ 
| ___ \         | |      | |   | ___ \         | |
| |_/ /___   ___| | _____| |_  | |_/ /__   ___ | |
|    // _ \ / __| |/ / _ \ __| |  __/ _ \ / _ \| |
| |\ \ (_) | (__|   <  __/ |_  | | | (_) | (_) | |
\_| \_\___/ \___|_|\_\___|\__| \_|  \___/ \___/|_|

%s`, cli.AppHelpTemplate)

    // Initialise application
    app := cli.NewApp()

    // Configure application
    app.Name     = "rocketpool"
    app.Usage    = "Rocket Pool node operator utilities"
    app.Version  = "0.0.1"
    app.Authors  = []cli.Author{
        cli.Author{
            Name:  "Darren Langley",
            Email: "darren@rocketpool.net",
        },
        cli.Author{
            Name:  "David Rugendyke",
            Email: "david@rocketpool.net",
        },
        cli.Author{
            Name:  "Jake Pospischil",
            Email: "jake@rocketpool.net",
        },
    }
    app.Copyright = "(c) 2019 Rocket Pool Pty Ltd"

    // Register global application options & defaults
    app.Flags = []cli.Flag{
        cli.StringFlag{
            Name:   "database, d",
            Usage:  "Rocket Pool CLI database absolute `path`",
            Value:  os.Getenv("HOME") + "/.rocketpool/rocketpool-cli.db",
        },
        cli.StringFlag{
            Name:   "keychain, k",
            Usage:  "PoW chain account keychain absolute `path`",
            Value:  os.Getenv("HOME") + "/.rocketpool/accounts",
        },
        cli.StringFlag{
            Name:   "provider, p",
            Usage:  "PoW chain provider `url`",
            Value:  "http://localhost:8545",
        },
        cli.StringFlag{
            Name:   "storageAddress, s",
            Usage:  "PoW chain Rocket Pool storage contract `address`",
            Value:  "0x70a5F2eB9e4C003B105399b471DAeDbC8d00B1c5", // Ganache
        },
    }

    // Register api commands
     deposit.RegisterCommands(app, "deposit",  []string{"d"})
         fee.RegisterCommands(app, "fee",      []string{"f"})
    minipool.RegisterCommands(app, "minipool", []string{"m"})
        node.RegisterCommands(app, "node",     []string{"n"})
     service.RegisterCommands(app, "service",  []string{"s"})

    // Run application
    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }

}


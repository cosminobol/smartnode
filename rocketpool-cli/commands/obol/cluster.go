package obol

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/rocket-pool/smartnode/v2/rocketpool-cli/client"
	"github.com/rocket-pool/smartnode/v2/rocketpool-cli/utils/tx"
)

func charonDkg(c *cli.Context) error {
	// Get RP client
	rp, err := client.NewClientFromCtx(c)
	if err != nil {
		return err
	}

	response, err := rp.Api.Obol.CharonDkg()
	if err != nil {
		return fmt.Errorf("Error creating a DV cluster: %w", err)
	}
	// Log & return
	fmt.Println("Successfully triggered cluster DKG creation.")
	return nil
}

func createENR(c *cli.Context) error {
	// Get RP client
	rp, err := client.NewClientFromCtx(c)
	if err != nil {
		return err
	}

	response, err := rp.Api.Obol.CreateENR()
	if err != nil {
		return fmt.Errorf("Error creating ENR: %w", err)
	}
	// Log & return
	fmt.Println("Successfully created ENR")
	return nil
}

func manageCharonService(c *cli.Context) error {
	// Get RP client
	rp, err := client.NewClientFromCtx(c)
	if err != nil {
		return err
	}

	response, err := rp.Api.Obol.ManageCharonService()
	if err != nil {
		return fmt.Errorf("Error managinig charon service: %w", err)
	}
	// Log & return
	fmt.Println("Successfully updated Charon service state")
	return nil
}

func getCharonHealth(c *cli.Context) error {
	// Get RP client
	rp, err := client.NewClientFromCtx(c)
	if err != nil {
		return err
	}

	response, err := rp.Api.Obol.GetCharonHealth()
	if err != nil {
		return fmt.Errorf("Error fetching charon health: %w", err)
	}
	// Log & return
	fmt.Println("Successfully fetched charon health")
	return nil
}


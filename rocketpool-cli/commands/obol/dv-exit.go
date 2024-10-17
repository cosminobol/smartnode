package obol

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/rocket-pool/smartnode/v2/rocketpool-cli/client"
	"github.com/rocket-pool/smartnode/v2/rocketpool-cli/utils/tx"
)

func dvExitBroadcast(c *cli.Context) error {
	// Get RP client
	rp, err := client.NewClientFromCtx(c)
	if err != nil {
		return err
	}

	// Check lot can be created
	response, err := rp.Api.Obol.DvExitBroadcast()
	if err != nil {
		return fmt.Errorf("Error triggering broadcast for a DV exit: %w", err)
	}
	// Log & return
	fmt.Println("Successfully triggered DV exit broadcast.")
	return nil
}

func dvExitSign(c *cli.Context) error {
	// Get RP client
	rp, err := client.NewClientFromCtx(c)
	if err != nil {
		return err
	}

	// Check lot can be created
	response, err := rp.Api.Obol.DvExitSign()
	if err != nil {
		return fmt.Errorf("Error signing for a DV exit: %w", err)
	}
	// Log & return
	fmt.Println("Successfully signed for a DV exit.")
	return nil
}


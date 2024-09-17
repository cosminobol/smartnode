package obol

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/rocket-pool/smartnode/v2/rocketpool-cli/client"
	"github.com/rocket-pool/smartnode/v2/rocketpool-cli/utils/tx"
)

func getValidatorPublicKeys(c *cli.Context) error {
	// Get RP client
	rp, err := client.NewClientFromCtx(c)
	if err != nil {
		return err
	}

	// Check lot can be created
	response, err := rp.Api.Obol.GetValidatorPublicKeys()
	if err != nil {
		return fmt.Errorf("Error fetching validator public keys: %w", err)
	}
	// Log & return
	fmt.Println("Successfully fetched validator public keys.")
	return nil
}

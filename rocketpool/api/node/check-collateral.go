package node

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	batch "github.com/rocket-pool/batch-query"
	"github.com/rocket-pool/rocketpool-go/node"
	"github.com/rocket-pool/rocketpool-go/rocketpool"
	"github.com/rocket-pool/smartnode/shared/services/config"
	"github.com/rocket-pool/smartnode/shared/types/api"
	rputils "github.com/rocket-pool/smartnode/shared/utils/rp"
)

type nodeCollateralHandler struct {
}

func (h *nodeCollateralHandler) CreateBindings(rp *rocketpool.RocketPool) error {
	return nil
}

func (h *nodeCollateralHandler) GetState(node *node.Node, mc *batch.MultiCaller) {
}

func (h *nodeCollateralHandler) PrepareResponse(rp *rocketpool.RocketPool, cfg *config.RocketPoolConfig, node *node.Node, opts *bind.TransactOpts, response *api.NodeCheckCollateralResponse) error {
	// Check collateral
	collateral, err := rputils.CheckCollateral(rp, node.Details.Address, nil)
	if err != nil {
		return err
	}
	response.EthMatched = collateral.EthMatched
	response.EthMatchedLimit = collateral.EthMatchedLimit
	response.PendingMatchAmount = collateral.PendingMatchAmount

	// Check if there's sufficient collateral including pending bond reductions
	remainingMatch := big.NewInt(0).Sub(response.EthMatchedLimit, response.EthMatched)
	remainingMatch.Sub(remainingMatch, response.PendingMatchAmount)
	response.InsufficientCollateral = (remainingMatch.Cmp(big.NewInt(0)) < 0)
	return nil
}

package obol

import (
	"fmt"
	"net/url"

	"log"

	"os/exec"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/gorilla/mux"
	batch "github.com/rocket-pool/batch-query"
	"github.com/rocket-pool/rocketpool-go/v2/rocketpool"

	"github.com/rocket-pool/node-manager-core/api/server"
	"github.com/rocket-pool/node-manager-core/api/types"
	"github.com/rocket-pool/smartnode/v2/shared/types/api"
)

// ===============
// === Factory ===
// ===============

type GetValidatorPublicKeysContextFactory struct {
	handler *ObolHandler
}

func (f *GetValidatorPublicKeysContextFactory) Create(args url.Values) (*GetValidatorPublicKeysContext, error) {
	c := &GetValidatorPublicKeysContext{
		handler: f.handler,
	}
	return c, nil
}

func (f *GetValidatorPublicKeysContextFactory) RegisterRoute(router *mux.Router) {
	server.RegisterSingleStageRoute[*GetValidatorPublicKeysContext, api.GetValidatorPublicKeysData](
		router, "obol/get-validator-public-keys", f, f.handler.logger.Logger, f.handler.serviceProvider.ServiceProvider,
	)
}

// ===============
// === Context ===
// ===============

type GetValidatorPublicKeysContext struct {
	handler *ObolHandler
	rp      *rocketpool.RocketPool

	endpoint string
}

func (c *GetValidatorPublicKeysContext) Initialize() (types.ResponseStatus, error) {
	sp := c.handler.serviceProvider
	c.rp = sp.GetRocketPool()

	// Requirements
	status, err := sp.RequireNodeRegistered(c.handler.ctx)
	if err != nil {
		return status, err
	}

	return types.ResponseStatus_Success, nil
}

func (c *GetValidatorPublicKeysContext) PrepareData(data *api.GetValidatorPublicKeysData, opts *bind.TransactOpts) (types.ResponseStatus, error) {
    cmd := exec.Command(
        "docker", "exec", "-it", "charon-distributed-validator-node-charon-1", 
        "/bin/sh", "-c", 
        fmt.Sprintf(`charon exit active-validator-list --beacon-node-endpoints="%s"`, c.endpoint),
    )
	output, err := cmd.CombinedOutput()

	if err != nil {
		return types.ResponseStatus_Error, fmt.Errorf("Error running docker command: %s", err)
	}
	log.Printf("Command output: %s", string(output))
	return types.ResponseStatus_Success, nil
}


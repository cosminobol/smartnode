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

type DvExitBroadcastContextFactory struct {
	handler *ObolHandler
}

func (f *DvExitBroadcastContextFactory) Create(args url.Values) (*DvExitBroadcastContext, error) {
	c := &DvExitBroadcastContext{
		handler: f.handler,
	}
	return c, nil
}

func (f *DvExitBroadcastContextFactory) RegisterRoute(router *mux.Router) {
	server.RegisterSingleStageRoute[*DvExitBroadcastContext, api.DvExitBroadcastData](
		router, "obol/dv-exit-broadcast", f, f.handler.logger.Logger, f.handler.serviceProvider.ServiceProvider,
	)
}

// ===============
// === Context ===
// ===============

type DvExitBroadcastContext struct {
	handler *ObolHandler
	rp      *rocketpool.RocketPool

	endpoint string
	validatorPublicKeys string
	publishTimeout string
}

func (c *DvExitBroadcastContext) Initialize() (types.ResponseStatus, error) {
	sp := c.handler.serviceProvider
	c.rp = sp.GetRocketPool()

	// Requirements
	status, err := sp.RequireNodeRegistered(c.handler.ctx)
	if err != nil {
		return status, err
	}

	return types.ResponseStatus_Success, nil
}

func (c *DvExitBroadcastContext) GetState(mc *batch.MultiCaller) {
}

func (c *DvExitBroadcastContext) PrepareData(data *api.DvExitBroadcastData, opts *bind.TransactOpts) (types.ResponseStatus, error) {
	cmd := exec.Command(
		"docker", "exec", "-it", "charon-distributed-validator-node-charon-1", 
		"/bin/sh", "-c", 
		fmt.Sprintf(`charon exit sign --beacon-node-endpoints="%s" --validator-public-key="%s" --publish-timeout="%s"`, 
			c.endpoint, 
			c.validatorPublicKeys, 
			c.publishTimeout,
		),
	)
	

	output, err := cmd.CombinedOutput()

	if err != nil {
		return types.ResponseStatus_Error, fmt.Errorf("Error running docker command: %s", err)
	}
	log.Printf("Command output: %s", string(output))
	return types.ResponseStatus_Success, nil
}


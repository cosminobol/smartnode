package obol

import (
	"fmt"
	// "strings"
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

type DvExitSignContextFactory struct {
	handler *ObolHandler
}

func (f *DvExitSignContextFactory) Create(args url.Values) (*DvExitSignContext, error) {
	c := &DvExitSignContext{
		handler: f.handler,
	}
	return c, nil
}

func (f *DvExitSignContextFactory) RegisterRoute(router *mux.Router) {
	server.RegisterSingleStageRoute[*DvExitSignContext, api.DvExitSignData](
		router, "obol/dv-exit-sign", f, f.handler.logger.Logger, f.handler.serviceProvider.ServiceProvider,
	)
}

// ===============
// === Context ===
// ===============

type DvExitSignContext struct {
	handler *ObolHandler
	rp      *rocketpool.RocketPool

	endpoint string
}

func (c *DvExitSignContext) Initialize() (types.ResponseStatus, error) {
	sp := c.handler.serviceProvider
	c.rp = sp.GetRocketPool()

	// Requirements
	status, err := sp.RequireNodeRegistered(c.handler.ctx)
	if err != nil {
		return status, err
	}

	// Bindings

	return types.ResponseStatus_Success, nil
}

func (c *DvExitSignContext) PrepareData(data *api.DvExitSignData, opts *bind.TransactOpts) (types.ResponseStatus, error) {
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


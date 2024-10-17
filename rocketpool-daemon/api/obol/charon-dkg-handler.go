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

type CharonDkgContextFactory struct {
	handler *ObolHandler
}

func (f *CharonDkgContextFactory) Create(args url.Values) (*CharonDkgContext, error) {
	c := &CharonDkgContext{
		handler: f.handler,
	}
	return c, nil
}

func (f *CharonDkgContextFactory) RegisterRoute(router *mux.Router) {
	server.RegisterSingleStageRoute[*CharonDkgContext, api.CharonDkgData](
		router, "obol/charon-dkg", f, f.handler.logger.Logger, f.handler.serviceProvider.ServiceProvider,
	)
}

// ===============
// === Context ===
// ===============

type CharonDkgContext struct {
	handler *ObolHandler
	rp      *rocketpool.RocketPool

	password string
}

func (c *CharonDkgContext) Initialize() (types.ResponseStatus, error) {
	sp := c.handler.serviceProvider
	c.rp = sp.GetRocketPool()

	// Requirements
	status, err := sp.RequireNodeRegistered(c.handler.ctx)
	if err != nil {
		return status, err
	}

	return types.ResponseStatus_Success, nil
}

func (c *CharonDkgContext) PrepareData(data *api.CharonDkgData, opts *bind.TransactOpts) (types.ResponseStatus, error) {
    cmd := exec.Command(
        "docker", "run", "--rm",
        "-v", fmt.Sprintf("%s:/opt/charon", c.password), 
        "obolnetwork/charon:v1.1.0",              
        "dkg", "--publish",                       
    )

	output, err := cmd.CombinedOutput()

	if err != nil {
		return types.ResponseStatus_Error, fmt.Errorf("Error running docker command: %s", err)
	} 
	log.Printf("Command output: %s", string(output))
	return types.ResponseStatus_Success, nil
}


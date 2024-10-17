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

type CreateENRContextFactory struct {
	handler *ObolHandler
}

func (f *CreateENRContextFactory) Create(args url.Values) (*CreateENRContext, error) {
	c := &CreateENRContext{
		handler: f.handler,
	}
	return c, nil
}

func (f *CreateENRContextFactory) RegisterRoute(router *mux.Router) {
	server.RegisterSingleStageRoute[*CreateENRContext, api.CreateENRData](
		router, "obol/create-enr", f, f.handler.logger.Logger, f.handler.serviceProvider.ServiceProvider,
	)
}

// ===============
// === Context ===
// ===============

type CreateENRContext struct {
	handler *ObolHandler
	rp      *rocketpool.RocketPool

	password string
}

func (c *CreateENRContext) Initialize() (types.ResponseStatus, error) {
	sp := c.handler.serviceProvider
	c.rp = sp.GetRocketPool()

	// Requirements
	status, err := sp.RequireNodeRegistered(c.handler.ctx)
	if err != nil {
		return status, err
	}

	return types.ResponseStatus_Success, nil
}

func (c *CreateENRContext) PrepareData(data *api.CreateENRData, opts *bind.TransactOpts) (types.ResponseStatus, error) {
    cmd := exec.Command(
        "docker", "run", "--rm",
        "-v", fmt.Sprintf("%s:/opt/charon", c.password), 
        "obolnetwork/charon:v1.1.0",              
        "create", "enr",                          
    )

	output, err := cmd.CombinedOutput()

	if err != nil {
		return types.ResponseStatus_Error, fmt.Errorf("Error running docker command: %s", err)
	}
	log.Printf("Command output: %s", string(output))
	return types.ResponseStatus_Success, nil
}


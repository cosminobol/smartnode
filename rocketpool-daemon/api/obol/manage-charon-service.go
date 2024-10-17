package obol

import (
	"fmt"
	"net/url"

	// "log"

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

type ManageCharonServiceContextFactory struct {
	handler *ObolHandler
}

func (f *ManageCharonServiceContextFactory) Create(args url.Values) (*ManageCharonServiceContext, error) {
	c := &ManageCharonServiceContext{
		handler: f.handler,
	}
	return c, nil
}

func (f *ManageCharonServiceContextFactory) RegisterRoute(router *mux.Router) {
	server.RegisterSingleStageRoute[*ManageCharonServiceContext, api.ManageCharonServiceData](
		router, "obol/manage-charon-service", f, f.handler.logger.Logger, f.handler.serviceProvider.ServiceProvider,
	)
}

// ===============
// === Context ===
// ===============

type ManageCharonServiceContext struct {
	handler *ObolHandler
	rp      *rocketpool.RocketPool

	serviceName string
	action string
}

func (c *ManageCharonServiceContext) Initialize() (types.ResponseStatus, error) {
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

func (c *ManageCharonServiceContext) PrepareData(data *api.ManageCharonServiceData, opts *bind.TransactOpts) (types.ResponseStatus, error) {
	var cmd *exec.Cmd

	// Choose the appropriate command based on the action
	switch c.action {
		case "start":
			cmd = exec.Command("docker", "start", c.serviceName)
		case "stop":
			cmd = exec.Command("docker", "stop", c.serviceName)
		case "restart":
			cmd = exec.Command("docker", "restart", c.serviceName)
		default:
			return types.ResponseStatus_Error, fmt.Errorf("invalid action: %s. Use start, stop, or restart", c.action)
	}

	// Execute the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		return types.ResponseStatus_Error, fmt.Errorf("failed to %s service %s: %v, output: %s", c.action, c.serviceName, err, output)
	}

	fmt.Printf("Service %s %sed successfully.\n", c.serviceName, c.action)
	return types.ResponseStatus_Success, nil
}


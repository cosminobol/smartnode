package obol

import (
	"fmt"
	"net/url"
	"net/http"

	"log"
	"io"

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

type GetCharonHealthContextFactory struct {
	handler *ObolHandler
}

func (f *GetCharonHealthContextFactory) Create(args url.Values) (*GetCharonHealthContext, error) {
	c := &GetCharonHealthContext{
		handler: f.handler,
	}
	return c, nil
}

func (f *GetCharonHealthContextFactory) RegisterRoute(router *mux.Router) {
	server.RegisterSingleStageRoute[*GetCharonHealthContext, api.GetCharonHealthData](
		router, "obol/get-charon-health", f, f.handler.logger.Logger, f.handler.serviceProvider.ServiceProvider,
	)
}

// ===============
// === Context ===
// ===============

type GetCharonHealthContext struct {
	handler *ObolHandler
	rp      *rocketpool.RocketPool

	endpoint string
}

func (c *GetCharonHealthContext) Initialize() (types.ResponseStatus, error) {
	sp := c.handler.serviceProvider
	c.rp = sp.GetRocketPool()

	// Requirements
	status, err := sp.RequireNodeRegistered(c.handler.ctx)
	if err != nil {
		return status, err
	}

	return types.ResponseStatus_Success, nil
}

func (c *GetCharonHealthContext) PrepareData(data *api.GetCharonHealthData, opts *bind.TransactOpts) (types.ResponseStatus, error) {
    // The URL for the health check
    url := fmt.Sprintf(`http://%s/health`, c.endpoint)

 	resp, err := http.Get(url)
	if err != nil {
		return types.ResponseStatus_Error, fmt.Errorf("Error running health check command: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return types.ResponseStatus_Error, fmt.Errorf("Request failed with status code: %s", resp.StatusCode)
    } 

	body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("Error reading response body: %s", err)
    }

    fmt.Printf("Response body:\n%s\n", body)
	return types.ResponseStatus_Success, nil
}


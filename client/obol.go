package client

import (
	"github.com/rocket-pool/node-manager-core/api/client"
	"github.com/rocket-pool/node-manager-core/api/types"
	"github.com/rocket-pool/smartnode/v2/shared/types/api"
)

type ObolRequester struct {
	context client.IRequesterContext
}

func NewObolRequester(context client.IRequesterContext) *ObolRequester {
	return &ObolRequester{
		context: context,
	}
}

func (r *ObolRequester) GetName() string {
	return "DvTest"
}
func (r *ObolRequester) GetRoute() string {
	return "dvtest"
}
func (r *ObolRequester) GetContext() client.IRequesterContext {
	return r.context
}

// Trigger a DV exit broadcast
func (r *ObolRequester) DvExitBroadcast() (*types.ApiResponse[api.DvExitBroadcastData], error) {
	return client.SendGetRequest[api.DvExitBroadcastData](r, "obol/dv-exit-broadcast", "DvExitBroadcast", nil)
}

// Trigger a DV exit sign
func (r *ObolRequester) DvExitSign() (*types.ApiResponse[api.DvExitSignData], error) {
	return client.SendGetRequest[api.DvExitSignData](r, "obol/dv-exit-sign", "DvExitSign", nil)
}

// Retrieve validator public keys
func (r *ObolRequester) GetValidatorPublicKeys() (*types.ApiResponse[api.GetValidatorPublicKeysData], error) {
	return client.SendGetRequest[api.GetValidatorPublicKeysData](r, "obol/get-validator-public-keys", "GetValidatorPublicKeys", nil)
}

// Creates cluster - DKG
func (r *ObolRequester) CharonDkg() (*types.ApiResponse[api.CharonDkgData], error) {
	return client.SendGetRequest[api.CharonDkgData](r, "obol/charon-dkg", "CharonDkg", nil)
}

// Creates ENR
func (r *ObolRequester) CreateENR() (*types.ApiResponse[api.CreateENRData], error) {
	return client.SendGetRequest[api.CreateENRData](r, "obol/create-enr", "CreateENR", nil)
}

// Manages Charon service
func (r *ObolRequester) ManageCharonService() (*types.ApiResponse[api.ManageCharonServiceData], error) {
	return client.SendGetRequest[api.ManageCharonServiceData](r, "obol/manage-charon-service", "ManageCharonService", nil)
}

// Retrieve Charon service health
func (r *ObolRequester) GetCharonHealth() (*types.ApiResponse[api.GetCharonHealthData], error) {
	return client.SendGetRequest[api.GetCharonHealthData](r, "obol/get-charon-health", "GetCharonHealth", nil)
}

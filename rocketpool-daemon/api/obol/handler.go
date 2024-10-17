package obol

import (
	"context"

	"github.com/gorilla/mux"

	"github.com/rocket-pool/node-manager-core/api/server"
	"github.com/rocket-pool/node-manager-core/log"
	"github.com/rocket-pool/smartnode/v2/rocketpool-daemon/common/services"
)

type ObolHandler struct {
	logger          *log.Logger
	ctx             context.Context
	serviceProvider *services.ServiceProvider
	factories       []server.IContextFactory
}

func NewObolHandler(logger *log.Logger, ctx context.Context, serviceProvider *services.ServiceProvider) *ObolHandler {
	h := &ObolHandler{
		logger:          logger,
		ctx:             ctx,
		serviceProvider: serviceProvider,
	}
	h.factories = []server.IContextFactory{
		&CharonDkgContextFactory{h},
		&CreateENRContextFactory{h},
		&DvExitBroadcastContextFactory{h},
		&DvExitSignContextFactory{h},
		&GetCharonHealthContextFactory{h},
		&GetValidatorPublicKeysContextFactory{h},
		&ManageCharonServiceContextFactory{h},
	}
	return h
}

func (h *ObolHandler) RegisterRoutes(router *mux.Router) {
	subrouter := router.PathPrefix("/obol").Subrouter()
	for _, factory := range h.factories {
		factory.RegisterRoute(subrouter)
	}
}

package handlers

import (
	"github.com/rts-core/versioner/api/generated/restapi/operations"
	"github.com/rts-core/versioner/db/access"
)

func Register(
	api *operations.VersionerAPI,
	applicationAccessor access.ApplicationsAccessor,
) {
	// Applications
	api.ApplicationsConsumeVersionHandler = NewConsumeVersionHandler(applicationAccessor)
	api.ApplicationsGetVersionHandler = NewGetVersionHandler(applicationAccessor)
}

package handlers

import (
	"versioner/api/generated/restapi/operations"
	"versioner/db/access"
)

func Register(
	api *operations.VersionerAPI,
	applicationAccessor access.ApplicationsAccessor,
) {
	// Applications
	api.ApplicationsConsumeVersionHandler = NewConsumeVersionHandler(applicationAccessor)
	api.ApplicationsGetVersionHandler = NewGetVersionHandler(applicationAccessor)
}
